package service

import (
	"context"

	"Ephyra-genesis-api/biz/dal"
	user "Ephyra-genesis-api/hertz_gen/basic/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserGetUserProfileService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserGetUserProfileService(Context context.Context, RequestContext *app.RequestContext) *UserGetUserProfileService {
	return &UserGetUserProfileService{RequestContext: RequestContext, Context: Context}
}

func (h *UserGetUserProfileService) Run(req *user.GetUserProfileRequest) (resp *user.GetUserProfileResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 从context获取用户参数
	contextParams, err := GetContextParams(h.Context, h.RequestContext)
	if err != nil {
		return nil, err
	}
	userID := contextParams.UserID

	// 获取用户信息
	userInfo, err := dal.GetUserByID(h.Context, userID)
	if err != nil {
		return nil, err
	}

	// 解析用户徽章
	badgeIDs, err := dal.ParseUserBadges(userInfo.Badges)
	if err != nil {
		return nil, err
	}

	// 获取所有徽章信息
	allBadges, err := dal.GetAllBadges(h.Context)
	if err != nil {
		return nil, err
	}

	// 构建用户徽章响应，标记已解锁的徽章
	var userBadges []*user.Badge
	badgeMap := make(map[int64]bool)
	for _, badgeID := range badgeIDs {
		badgeMap[badgeID] = true
	}

	for _, badge := range allBadges {
		unlocked := badgeMap[badge.ID]
		userBadges = append(userBadges, &user.Badge{
			Id:               badge.ID,
			Name:             badge.Name,
			ImageUrl:         badge.ImageURL,
			Unlocked:         unlocked,
			UnlockConditions: badge.UnlockConditions,
		})
	}

	// 计算用户贡献统计
	totalDays, consecutiveDays, maxConsecutiveDays, err := dal.CalculateUserContributions(h.Context, userID, userInfo.TaskStatus)
	if err != nil {
		return nil, err
	}

	contributions := &user.UserContributions{
		TotalDays:          totalDays,
		ConsecutiveDays:    consecutiveDays,
		MaxConsecutiveDays: maxConsecutiveDays,
	}

	// 获取用户最近5次积分历史
	pointsLogs, err := dal.GetUserPointsLogsByUser(h.Context, userID, 5)
	if err != nil {
		return nil, err
	}

	var userPointsLogs []*user.UserPointsLogs
	for _, log := range pointsLogs {
		userPointsLogs = append(userPointsLogs, &user.UserPointsLogs{
			TotalPoints: int64(log.TotalPoints),
			BasePoints:  int64(log.BasePoints),
			ExtraPoints: int64(log.ExtraPoints),
			Type:        log.Type,
			CreateTime:  log.CreateAt.Unix(),
		})
	}

	resp = &user.GetUserProfileResponse{
		Uid:           userInfo.ID,
		Address:       userInfo.Address,
		UserName:      userInfo.UserName,
		AvatarUrl:     userInfo.AvatarURL,
		Points:        userInfo.Points,
		Badges:        userBadges,
		Contributions: contributions,
		PointsLogs:    userPointsLogs,
		SbtTokenId:    userInfo.SbtTokenID,
	}
	return
}
