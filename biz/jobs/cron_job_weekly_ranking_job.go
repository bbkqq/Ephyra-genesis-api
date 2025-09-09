package jobs

import (
	"Ephyra-genesis-api/biz/dal"
	"Ephyra-genesis-api/biz/dal/redis"
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	redisLib "github.com/redis/go-redis/v9"
)

type WeeklyRankingJob struct{}

func NewWeeklyRankingJob() *WeeklyRankingJob {
	return &WeeklyRankingJob{}
}

func (w *WeeklyRankingJob) Do(ctx context.Context) {
	hlog.CtxInfof(ctx, "[WeeklyRankingJob] Start processing weekly ranking")

	// 获取本周开始时间（周一）
	now := time.Now()
	weekStart := now.AddDate(0, 0, -int(now.Weekday())+1)
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())

	// 计算前7天（不包括今天）
	today := time.Now()
	weeklyPoints := make(map[string]float64) // userID -> totalPoints

	// 遍历前6天的数据（不包括今天）
	for i := 1; i <= 6; i++ {
		targetDate := today.AddDate(0, 0, -i)

		// 获取该天的所有用户积分数据
		userIDs, scores, err := dal.GetTopUsersByDay(ctx, targetDate, 10000) // 获取大量数据避免遗漏
		if err != nil {
			hlog.CtxErrorf(ctx, "[WeeklyRankingJob] Get daily points for %s error: %v", targetDate.Format(time.DateOnly), err)
			continue
		}

		// 累加到周积分中
		for j, userID := range userIDs {
			if j < len(scores) {
				weeklyPoints[userID] += scores[j]
			}
		}

		hlog.CtxInfof(ctx, "[WeeklyRankingJob] Processed %d users for date %s", len(userIDs), targetDate.Format(time.DateOnly))
	}

	if len(weeklyPoints) == 0 {
		hlog.CtxInfof(ctx, "[WeeklyRankingJob] No weekly points data found")
		return
	}

	// 生成周排名ZSet key
	weeklyZSetKey := redis.UserWeeklyPointsZSetKey(weekStart)

	// 清空之前的周排名数据
	err := redis.RedisClient.Del(ctx, weeklyZSetKey).Err()
	if err != nil {
		hlog.CtxErrorf(ctx, "[WeeklyRankingJob] Clear weekly ranking error: %v", err)
		return
	}

	// 批量写入前1000名用户的周积分数据
	count := 0
	for userID, totalPoints := range weeklyPoints {
		if count >= 1000 {
			break
		}

		err = redis.RedisClient.ZAdd(ctx, weeklyZSetKey, redisLib.Z{
			Score:  totalPoints,
			Member: userID,
		}).Err()

		if err != nil {
			hlog.CtxErrorf(ctx, "[WeeklyRankingJob] ZAdd error for user %s: %v", userID, err)
			continue
		}

		count++
	}

	// 设置过期时间为14天
	redis.RedisClient.Expire(ctx, weeklyZSetKey, 14*24*time.Hour)

	hlog.CtxInfof(ctx, "[WeeklyRankingJob] Successfully updated weekly ranking with %d users", count)
}

func (w *WeeklyRankingJob) Interval() time.Duration {
	return 12 * time.Hour // 每12小时执行一次
}

func (w *WeeklyRankingJob) DisableSerializable() bool {
	return false
}
