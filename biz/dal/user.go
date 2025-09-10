package dal

import (
	"Ephyra-genesis-api/biz/bizerror"
	"Ephyra-genesis-api/biz/dal/mysql"
	"Ephyra-genesis-api/biz/dal/redis"
	"Ephyra-genesis-api/storage/mysql/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

// 这个是用来保存用户是否完成每天任务的结构，长度是300的全0字符串是初始值
// defaultTaskStatus[0] 表示第一天的任务状态，defaultTaskStatus[0]=0 表示第一天没有完成，defaultTaskStatus[0]=1表示第一天完成了
// 起始的时间是 2025-09-06
// 后面判断当天有没有完成的时候，是看当前的天数跟 起始时间的 offset，然后去 defaultTaskStatus数组查找
// 这样就可以计算出从当天算起，之前连续完成了多少天，然后还可以计算历史最长的连续天数
// api 侧不更新这个数据，只有在cronjob的时候才会更新这个数据
var defaultTaskStatus = strings.Repeat("0", 300)

// 错误定义
var (
	ErrUserNotFound = errors.New("user not found")
)

// 徽章常量定义
const (
	BadgeFirstParticipation = 1 // 初次参与徽章
	BadgeSevenDayOracle     = 2 // 七日先知徽章
	BadgeHundredDayWise     = 3 // 百问智者徽章
)

// 积分计算常量
const (
	BasePoints         = 50  // 基础积分
	ConsecutiveBonus   = 10  // 连续参与奖励（每天递增10分）
	MaxDailyPoints     = 110 // 每日最高积分
	MaxConsecutiveDays = 7   // 最大连续天数奖励
	MaxExtraPoints     = 60  // 最大额外积分（第7天及以后）
)

// 起始时间常量
var StartDate = time.Date(2025, 9, 9, 0, 0, 0, 0, time.UTC)

// 任务奖励结果
type TaskRewardResult struct {
	Points          int32   `json:"points"`           // 基础积分
	ExtraPoints     int32   `json:"extra_points"`     // 连续参与额外积分
	NewBadges       []int64 `json:"new_badges"`       // 新获得的徽章ID
	IsNewUser       bool    `json:"is_new_user"`      // 是否新用户
	ConsecutiveDays int64   `json:"consecutive_days"` // 当前连续天数
}

// ========== User 相关操作 ==========

func GetUserByAddress(ctx context.Context, address string) (*model.User, error) {
	var users []*model.User
	err := mysql.DB.Model(&model.User{}).Where(`address = ?`, address).Find(&users).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetUserByAddress address: %v error: %v", address, err)
		return nil, bizerror.DBError
	}

	if len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
}

func GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	var user model.User
	err := mysql.DB.Model(&model.User{}).Where(`id = ?`, userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "GetUserByID userID: %v error: %v", userID, err)
		return nil, bizerror.DBError
	}
	return &user, nil
}

func CreateNewUser(ctx context.Context, address string) (*model.User, error) {
	userModel := &model.User{
		Address:    address,
		UserName:   "",
		AvatarURL:  "",
		Points:     0,
		Badges:     "[]",
		TaskStatus: defaultTaskStatus,
		UpdateAt:   time.Now(),
		CreateAt:   time.Now(),
	}

	err := mysql.DB.Model(&model.User{}).Create(userModel).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "CreateNewUser error: %v", err)
		return nil, err
	}
	return userModel, nil
}

func UpdateUserProfile(ctx context.Context, userID int64, userName, avatarURL string) error {
	err := mysql.DB.Model(&model.User{}).Where(`id = ?`, userID).Updates(map[string]interface{}{
		"user_name":  userName,
		"avatar_url": avatarURL,
		"update_at":  time.Now(),
	}).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "UpdateUserProfile userID: %v error: %v", userID, err)
		return bizerror.DBError
	}
	return nil
}

func GetUserRanking(ctx context.Context, rankType string, limit int) ([]*model.User, error) {
	var users []*model.User
	query := mysql.DB.Model(&model.User{}).Order("points DESC")

	// 根据排名类型添加时间过滤条件
	switch rankType {
	//case "daily":
	//	// 今日排名，这里简化处理，实际可能需要根据积分日志计算
	//	query = query.Where("DATE(update_at) = CURDATE()")
	//case "weekly":
	//	// 本周排名
	//	query = query.Where("WEEK(update_at) = WEEK(NOW())")
	case "all":
		// 总排名，不需要额外条件
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&users).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetUserRanking error: %v", err)
		return nil, bizerror.DBError
	}
	return users, nil
}

func GetUserRankingPosition(ctx context.Context, userID int64, rankType string) (int64, error) {
	var rank int64
	var query string

	switch rankType {
	case "daily":
		query = `SELECT COUNT(*) + 1 FROM user WHERE points > (SELECT points FROM user WHERE id = ?) AND DATE(update_at) = CURDATE()`
	case "weekly":
		query = `SELECT COUNT(*) + 1 FROM user WHERE points > (SELECT points FROM user WHERE id = ?) AND WEEK(update_at) = WEEK(NOW())`
	case "all":
		query = `SELECT COUNT(*) + 1 FROM user WHERE points > (SELECT points FROM user WHERE id = ?)`
	}

	err := mysql.DB.Raw(query, userID).Scan(&rank).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetUserRankingPosition error: %v", err)
		return 0, bizerror.DBError
	}
	return rank, nil
}

// ========== Badge 相关操作 ==========

func GetAllBadges(ctx context.Context) ([]*model.Badge, error) {
	var badges []*model.Badge
	err := mysql.DB.Model(&model.Badge{}).Find(&badges).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetAllBadges error: %v", err)
		return nil, bizerror.DBError
	}
	return badges, nil
}

func GetBadgesByIDs(ctx context.Context, badgeIDs []int64) ([]*model.Badge, error) {
	var badges []*model.Badge
	if len(badgeIDs) == 0 {
		return badges, nil
	}

	err := mysql.DB.Model(&model.Badge{}).Where("id IN ?", badgeIDs).Find(&badges).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetBadgesByIDs error: %v", err)
		return nil, bizerror.DBError
	}
	return badges, nil
}

// ========== Questions 相关操作 ==========

func GetDailyQuestions(ctx context.Context, limit int) ([]*model.Question, error) {
	var questions []*model.Question

	// 默认获取10个问题
	if limit <= 0 {
		limit = 10
	}

	// 使用ORDER BY RAND()随机获取问题
	err := mysql.DB.Model(&model.Question{}).Order("RAND()").Limit(limit).Find(&questions).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetDailyQuestions error: %v", err)
		return nil, bizerror.DBError
	}
	return questions, nil
}

// ========== UserDailyTaskLog 相关操作 ==========

func GetUserDailyTaskLogCount(ctx context.Context) (int64, error) {
	var count int64
	err := mysql.DB.Model(&model.UserDailyTaskLog{}).Count(&count).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetUserDailyTaskLogCount error: %v", err)
		return 0, bizerror.DBError
	}
	return count, nil
}

func GetUserDailyTaskLogByUserAndDay(ctx context.Context, userID int64, currentDay int32) (*model.UserDailyTaskLog, error) {
	var log model.UserDailyTaskLog
	err := mysql.DB.Model(&model.UserDailyTaskLog{}).Where("user_id = ? AND current_day = ?", userID, currentDay).First(&log).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "GetUserDailyTaskLogByUserAndDay error: %v", err)
		return nil, bizerror.DBError
	}
	return &log, nil
}

func GetUserDailyTaskLogByTxHash(ctx context.Context, txHash string) (*model.UserDailyTaskLog, error) {
	var log model.UserDailyTaskLog
	err := mysql.DB.Model(&model.UserDailyTaskLog{}).Where("tx_hash = ?", txHash).First(&log).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "GetUserDailyTaskLogByTxHash error: %v", err)
		return nil, bizerror.DBError
	}
	return &log, nil
}

func GetUserDailyTaskLogsByUser(ctx context.Context, userID int64) ([]*model.UserDailyTaskLog, error) {
	var logs []*model.UserDailyTaskLog
	err := mysql.DB.Model(&model.UserDailyTaskLog{}).Where("user_id = ?", userID).Order("current_day DESC").Find(&logs).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetUserDailyTaskLogsByUser error: %v", err)
		return nil, bizerror.DBError
	}
	return logs, nil
}

// ========== UserPointsLog 相关操作 ==========

func GetUserPointsLogsByUser(ctx context.Context, userID int64, limit int) ([]*model.UserPointsLog, error) {
	var logs []*model.UserPointsLog
	query := mysql.DB.Model(&model.UserPointsLog{}).Where("user_id = ?", userID).Order("create_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&logs).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "GetUserPointsLogsByUser error: %v", err)
		return nil, bizerror.DBError
	}
	return logs, nil
}

func CreateUserPointsLog(ctx context.Context, userID int64, totalPoints, basePoints, extraPoints int32, pointType string) error {
	log := &model.UserPointsLog{
		UserID:      userID,
		TotalPoints: totalPoints,
		BasePoints:  basePoints,
		ExtraPoints: extraPoints,
		Type:        pointType,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	err := mysql.DB.Model(&model.UserPointsLog{}).Create(log).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "CreateUserPointsLog error: %v", err)
		return bizerror.DBError
	}
	return nil
}

// ========== 辅助函数 ==========

// 计算用户贡献统计
func CalculateUserContributions(ctx context.Context, userID int64, taskStatus string) (totalDays, consecutiveDays, maxConsecutiveDays int64, err error) {
	// 获取用户的任务日志
	logs, err := GetUserDailyTaskLogsByUser(ctx, userID)
	if err != nil {
		return 0, 0, 0, err
	}

	// 计算总参与天数
	totalDays = int64(len(logs))

	// 根据taskStatus计算连续天数
	consecutiveDays, maxConsecutiveDays = calculateConsecutiveDays(taskStatus)

	return totalDays, consecutiveDays, maxConsecutiveDays, nil
}

func calculateConsecutiveDays(taskStatus string) (current, max int64) {
	if len(taskStatus) == 0 {
		return 0, 0
	}

	current = 0
	max = 0
	tempMax := int64(0)

	// 从最后一天开始计算当前连续天数
	for i := len(taskStatus) - 1; i >= 0; i-- {
		if taskStatus[i] == '1' {
			current++
		} else {
			break
		}
	}

	// 计算历史最长连续天数
	for i := 0; i < len(taskStatus); i++ {
		if taskStatus[i] == '1' {
			tempMax++
			if tempMax > max {
				max = tempMax
			}
		} else {
			tempMax = 0
		}
	}

	return current, max
}

// 解析用户徽章JSON
func ParseUserBadges(badgesJSON string) ([]int64, error) {
	var badges []int64
	if badgesJSON == "" || badgesJSON == "[]" {
		return badges, nil
	}

	err := json.Unmarshal([]byte(badgesJSON), &badges)
	if err != nil {
		return nil, fmt.Errorf("parse user badges error: %v", err)
	}
	return badges, nil
}

// 序列化用户徽章为JSON
func SerializeUserBadges(badges []int64) string {
	if len(badges) == 0 {
		return "[]"
	}

	data, err := json.Marshal(badges)
	if err != nil {
		return "[]"
	}
	return string(data)
}

// ========== 任务奖励计算 ==========

// CalculateTaskReward 计算任务奖励（可被API调用预览积分）
func CalculateTaskReward(ctx context.Context, userModel *model.User, targetDay int32) (*TaskRewardResult, error) {
	result := &TaskRewardResult{
		Points:          0,
		ExtraPoints:     0,
		NewBadges:       []int64{},
		IsNewUser:       false,
		ConsecutiveDays: 0,
	}

	// 检查是否为新用户（第一次参与）
	isNewUser := userModel.TaskStatus == defaultTaskStatus || userModel.TaskStatus == "" || !strings.Contains(userModel.TaskStatus, "1")
	result.IsNewUser = isNewUser

	// 模拟更新后的task_status来计算连续天数
	simulatedTaskStatus := userModel.TaskStatus
	if len(simulatedTaskStatus) < int(targetDay)+1 {
		simulatedTaskStatus = simulatedTaskStatus + strings.Repeat("0", int(targetDay)+1-len(simulatedTaskStatus))
	}

	// 设置目标天数为完成状态
	taskStatusRunes := []rune(simulatedTaskStatus)
	if int(targetDay) < len(taskStatusRunes) {
		taskStatusRunes[targetDay] = '1'
		simulatedTaskStatus = string(taskStatusRunes)
	}

	// 计算连续天数（从targetDay往前计算）
	consecutiveDays := calculateConsecutiveDaysFromDay(simulatedTaskStatus, int(targetDay))
	result.ConsecutiveDays = consecutiveDays

	// 计算积分：基础50分 + 连续天数奖励
	// 第1天: 50分, 第2天: 60分, 第3天: 70分, ..., 第7天: 110分, 第8天及以后: 110分
	basePoints := int32(BasePoints) // 基础50分
	extraPoints := int32(0)         // 额外积分

	if consecutiveDays > 1 {
		// 额外积分 = (连续天数 - 1) * 10，最多60分
		bonus := int32((consecutiveDays - 1) * ConsecutiveBonus)
		if bonus > int32(MaxExtraPoints) {
			bonus = int32(MaxExtraPoints)
		}
		extraPoints = bonus
	}

	result.Points = basePoints
	result.ExtraPoints = extraPoints

	// 解析用户当前徽章
	userBadges, err := ParseUserBadges(userModel.Badges)
	if err != nil {
		hlog.CtxErrorf(ctx, "[CalculateTaskReward] Parse user badges error: %v", err)
		userBadges = []int64{}
	}

	// 检查新徽章
	newBadges := []int64{}

	// 1. 初次参与徽章
	if isNewUser && !containsBadge(userBadges, BadgeFirstParticipation) {
		newBadges = append(newBadges, BadgeFirstParticipation)
	}

	// 2. 七日先知徽章（连续7天参与，只获得一次）
	if consecutiveDays >= 7 && !containsBadge(userBadges, BadgeSevenDayOracle) {
		newBadges = append(newBadges, BadgeSevenDayOracle)
	}

	// 3. 百问智者徽章（累计参与100天）
	totalParticipationDays := int64(strings.Count(simulatedTaskStatus, "1"))
	if totalParticipationDays >= 100 && !containsBadge(userBadges, BadgeHundredDayWise) {
		newBadges = append(newBadges, BadgeHundredDayWise)
	}

	result.NewBadges = newBadges

	return result, nil
}

// calculateConsecutiveDaysFromDay 从指定天数往前计算连续天数
func calculateConsecutiveDaysFromDay(taskStatus string, targetDay int) int64 {
	if len(taskStatus) == 0 || targetDay < 0 || targetDay >= len(taskStatus) {
		return 0
	}

	consecutiveDays := int64(0)
	// 从targetDay开始往前计算连续天数
	for i := targetDay; i >= 0; i-- {
		if taskStatus[i] == '1' {
			consecutiveDays++
		} else {
			break
		}
	}

	return consecutiveDays
}

// containsBadge 检查徽章列表是否包含指定徽章
func containsBadge(badges []int64, badgeID int64) bool {
	for _, badge := range badges {
		if badge == badgeID {
			return true
		}
	}
	return false
}

// ProcessTaskCompletion 处理任务完成（包含数据库更新）
func ProcessTaskCompletion(ctx context.Context, userModel *model.User, targetDay int32, answer string, timestamp int64, txHash string) error {
	// 1. 计算奖励
	rewardResult, err := CalculateTaskReward(ctx, userModel, targetDay)
	if err != nil {
		return err
	}

	// 2. 创建用户每日任务日志
	newBadgesJSON := SerializeUserBadges(rewardResult.NewBadges)

	taskLog := &model.UserDailyTaskLog{
		UserID:      userModel.ID,
		Answer:      answer,
		CurrentDay:  targetDay,
		Badges:      newBadgesJSON,
		TxHash:      txHash,
		Points:      rewardResult.Points,      // 直接使用计算结果
		ExtraPoints: rewardResult.ExtraPoints, // 直接使用计算结果
		CreateAt:    time.Unix(timestamp, 0),
		UpdateAt:    time.Now(),
	}

	// 3. 更新用户task_status（使用传入的userModel，避免重复查询）
	updatedTaskStatus, err := updateTaskStatusString(userModel.TaskStatus, targetDay)
	if err != nil {
		return err
	}

	// 4. 更新用户徽章
	var updatedBadges []int64
	if len(rewardResult.NewBadges) > 0 {
		currentBadges, _ := ParseUserBadges(userModel.Badges)
		updatedBadges = append(currentBadges, rewardResult.NewBadges...)
	} else {
		updatedBadges, _ = ParseUserBadges(userModel.Badges)
	}
	updatedBadgesJSON := SerializeUserBadges(updatedBadges)
	totalPoints := rewardResult.Points + rewardResult.ExtraPoints

	// 5. 在事务中执行所有数据库操作
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户每日任务日志
		if err := tx.Model(&model.UserDailyTaskLog{}).Create(taskLog).Error; err != nil {
			hlog.CtxErrorf(ctx, "[ProcessTaskCompletion] Create task log error: %v", err)
			return err
		}

		// 更新用户信息
		if err := tx.Model(&model.User{}).Where("id = ?", userModel.ID).Updates(map[string]interface{}{
			"task_status": updatedTaskStatus,
			"badges":      updatedBadgesJSON,
			"points":      tx.Raw("points + ?", totalPoints),
			"update_at":   time.Now(),
		}).Error; err != nil {
			hlog.CtxErrorf(ctx, "[ProcessTaskCompletion] Update user error: %v", err)
			return err
		}

		// 创建积分日志
		pointsLog := &model.UserPointsLog{
			UserID:      userModel.ID,
			TotalPoints: totalPoints,
			BasePoints:  rewardResult.Points,
			ExtraPoints: rewardResult.ExtraPoints,
			Type:        "daily_task",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		}
		if err := tx.Model(&model.UserPointsLog{}).Create(pointsLog).Error; err != nil {
			hlog.CtxErrorf(ctx, "[ProcessTaskCompletion] Create points log error: %v", err)
			return err
		}

		return nil
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "[ProcessTaskCompletion] Transaction failed: %v", err)
		return bizerror.DBError
	}

	// 7. 记录用户当天获得的积分到Redis ZSet用于排名
	eventDate := time.Unix(timestamp, 0)
	err = RecordUserDailyPoints(ctx, userModel.ID, totalPoints, eventDate)
	if err != nil {
		// Redis记录失败不影响主流程，只记录日志
		hlog.CtxErrorf(ctx, "[ProcessTaskCompletion] Record daily points to Redis failed: %v", err)
	}

	hlog.CtxInfof(ctx, "[ProcessTaskCompletion] Successfully processed task for user %d, day %d, base points: %d, extra points: %d, total: %d, new badges: %v",
		userModel.ID, targetDay, rewardResult.Points, rewardResult.ExtraPoints, totalPoints, rewardResult.NewBadges)

	return nil
}

// updateTaskStatusString 更新task_status字符串
func updateTaskStatusString(currentStatus string, targetDay int32) (string, error) {
	taskStatus := currentStatus
	if len(taskStatus) < int(targetDay)+1 {
		taskStatus = taskStatus + strings.Repeat("0", int(targetDay)+1-len(taskStatus))
	}

	taskStatusRunes := []rune(taskStatus)
	if int(targetDay) < len(taskStatusRunes) {
		taskStatusRunes[targetDay] = '1'
		taskStatus = string(taskStatusRunes)
	}

	return taskStatus, nil
}

// CalculateDayOffset 计算基于起始时间的天数偏移
func CalculateDayOffset(targetTime time.Time) int32 {
	targetTime = targetTime.UTC()
	daysDiff := int32(targetTime.Sub(StartDate).Hours() / 24)
	if daysDiff < 0 {
		return 0
	}
	return daysDiff
}

// ========== 链上事件处理 ==========

// ProcessAnswerOnChainEvent 处理链上答案提交事件
func ProcessAnswerOnChainEvent(ctx context.Context, userAddress, answer string, questionID int64, timestamp int64, day int64, txHash string) error {
	// 1. 根据地址获取或创建用户
	userModel, err := GetUserByAddress(ctx, userAddress)
	if err != nil {
		return err
	}

	// 如果用户不存在，创建新用户
	if userModel == nil {
		userModel, err = CreateNewUser(ctx, userAddress)
		if err != nil {
			return err
		}
	}

	// 2. 检查是否已经存在该 txHash 的任务记录（避免重复处理）
	existingLog, err := GetUserDailyTaskLogByTxHash(ctx, txHash)
	if err != nil {
		return err
	}

	// 如果已经存在记录，跳过
	if existingLog != nil {
		hlog.CtxInfof(ctx, "[ProcessAnswerOnChainEvent] Task log already exists for txHash: %s", txHash)
		return nil
	}

	// 3. 使用新的任务完成处理逻辑（包含Redis积分记录）
	err = ProcessTaskCompletion(ctx, userModel, int32(day), answer, timestamp, txHash)
	if err != nil {
		return err
	}

	hlog.CtxInfof(ctx, "[ProcessAnswerOnChainEvent] Successfully processed answer event for user %s, day %d", userAddress, day)
	return nil
}

// UpdateUserPoints 更新用户总积分（保留用于其他地方可能的调用）
func UpdateUserPoints(ctx context.Context, userID int64, pointsToAdd int64) error {
	err := mysql.DB.Model(&model.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"points":    mysql.DB.Raw("points + ?", pointsToAdd),
		"update_at": time.Now(),
	}).Error

	if err != nil {
		hlog.CtxErrorf(ctx, "[UpdateUserPoints] Update user points error: %v", err)
		return bizerror.DBError
	}

	return nil
}

// ========== Redis 积分排名 ==========

// RecordUserDailyPoints 记录用户当天获得的积分到Redis ZSet
func RecordUserDailyPoints(ctx context.Context, userID int64, totalPoints int32, date time.Time) error {
	// 生成当天的ZSet key
	zsetKey := redis.UserDailyPointsZSetKey(date)

	// 使用ZIncrBy增加用户的积分
	err := redis.RedisClient.ZIncrBy(ctx, zsetKey, float64(totalPoints), fmt.Sprintf("%d", userID)).Err()
	if err != nil {
		hlog.CtxErrorf(ctx, "[RecordUserDailyPoints] ZIncrBy error: %v", err)
		return bizerror.DBError
	}

	// 设置key的过期时间为7天，避免Redis内存占用过多
	redis.RedisClient.Expire(ctx, zsetKey, 10*24*time.Hour)

	hlog.CtxInfof(ctx, "[RecordUserDailyPoints] Recorded %d points for user %d on %s", totalPoints, userID, date.Format(time.DateOnly))
	return nil
}

// GetUserDailyRanking 获取用户在指定日期的积分排名
func GetUserDailyRanking(ctx context.Context, userID int64, date time.Time) (int64, error) {
	zsetKey := redis.UserDailyPointsZSetKey(date)

	// 获取用户排名（从高到低排序，排名从0开始）
	rank, err := redis.RedisClient.ZRevRank(ctx, zsetKey, fmt.Sprintf("%d", userID)).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetUserDailyRanking] ZRevRank error: %v", err)
		return 0, bizerror.DBError
	}

	// 排名从1开始计数
	return rank + 1, nil
}

// GetTopUsersByDay 获取指定日期的积分排行榜
func GetTopUsersByDay(ctx context.Context, date time.Time, limit int64) ([]string, []float64, error) {
	zsetKey := redis.UserDailyPointsZSetKey(date)

	// 获取排行榜前N名（从高到低）
	result, err := redis.RedisClient.ZRevRangeWithScores(ctx, zsetKey, 0, limit-1).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetTopUsersByDay] ZRevRangeWithScores error: %v", err)
		return nil, nil, bizerror.DBError
	}

	userIDs := make([]string, len(result))
	scores := make([]float64, len(result))

	for i, item := range result {
		userIDs[i] = item.Member.(string)
		scores[i] = item.Score
	}

	return userIDs, scores, nil
}

// GetDailyRankingFromRedis 从Redis获取日榜数据
func GetDailyRankingFromRedis(ctx context.Context, userID int64, limit int) ([]*model.User, int64, error) {
	today := time.Now()

	// 获取当天排行榜数据
	userIDs, scores, err := GetTopUsersByDay(ctx, today, int64(limit))
	if err != nil {
		return nil, 0, err
	}

	// 获取用户排名
	var userRanking int64 = 0
	if userID > 0 {
		userRanking, err = GetUserDailyRanking(ctx, userID, today)
		if err != nil {
			userRanking = 0
		}
	}

	// 根据userIDs获取用户详细信息
	var users []*model.User
	if len(userIDs) > 0 {
		var userIDsInt64 []int64
		for _, idStr := range userIDs {
			userIDsInt64 = append(userIDsInt64, cast.ToInt64(idStr))
		}

		if len(userIDsInt64) > 0 {
			err = mysql.DB.Model(&model.User{}).Where("id IN ?", userIDsInt64).Find(&users).Error
			if err != nil {
				hlog.CtxErrorf(ctx, "[GetDailyRankingFromRedis] Get users error: %v", err)
				return nil, 0, bizerror.DBError
			}

			// 按照Redis中的顺序重新排列用户
			userMap := make(map[int64]*model.User)
			for _, u := range users {
				userMap[u.ID] = u
			}

			orderedUsers := make([]*model.User, 0, len(userIDs))
			for i, idStr := range userIDs {
				var id int64
				if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
					if u, exists := userMap[id]; exists {
						// 设置当天获得的积分（从Redis分数获取）
						u.Points = int64(scores[i])
						orderedUsers = append(orderedUsers, u)
					}
				}
			}
			users = orderedUsers
		}
	}

	return users, userRanking, nil
}

// GetWeeklyRankingFromRedis 从Redis获取周榜数据
func GetWeeklyRankingFromRedis(ctx context.Context, userID int64, limit int) ([]*model.User, int64, error) {
	// 获取本周开始时间（周一）
	now := time.Now()
	weekStart := now.AddDate(0, 0, -int(now.Weekday())+1)
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())

	zsetKey := redis.UserWeeklyPointsZSetKey(weekStart)

	// 获取排行榜前N名
	result, err := redis.RedisClient.ZRevRangeWithScores(ctx, zsetKey, 0, int64(limit)-1).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetWeeklyRankingFromRedis] ZRevRangeWithScores error: %v", err)
		return nil, 0, bizerror.DBError
	}

	// 获取用户排名
	var userRanking int64 = 1000 // 默认1000
	if userID > 0 {
		rank, err := redis.RedisClient.ZRevRank(ctx, zsetKey, fmt.Sprintf("%d", userID)).Result()
		if err == nil {
			userRanking = rank + 1
			if userRanking > 1000 {
				userRanking = 1000
			}
		}
	}

	// 根据userIDs获取用户详细信息
	var users []*model.User
	if len(result) > 0 {
		var userIDsInt64 []int64
		for _, item := range result {
			var id int64
			if _, err := fmt.Sscanf(item.Member.(string), "%d", &id); err == nil {
				userIDsInt64 = append(userIDsInt64, id)
			}
		}

		if len(userIDsInt64) > 0 {
			err = mysql.DB.Model(&model.User{}).Where("id IN ?", userIDsInt64).Find(&users).Error
			if err != nil {
				hlog.CtxErrorf(ctx, "[GetWeeklyRankingFromRedis] Get users error: %v", err)
				return nil, 0, bizerror.DBError
			}

			// 按照Redis中的顺序重新排列用户
			userMap := make(map[int64]*model.User)
			for _, u := range users {
				userMap[u.ID] = u
			}

			orderedUsers := make([]*model.User, 0, len(result))
			for _, item := range result {
				var id int64
				if _, err := fmt.Sscanf(item.Member.(string), "%d", &id); err == nil {
					if u, exists := userMap[id]; exists {
						// 设置周积分（从Redis分数获取）
						u.Points = int64(item.Score)
						orderedUsers = append(orderedUsers, u)
					}
				}
			}
			users = orderedUsers
		}
	}

	return users, userRanking, nil
}
