package service

import (
	"context"

	"Ephyra-genesis-api/biz/dal"
	user "Ephyra-genesis-api/hertz_gen/basic/user"
	"Ephyra-genesis-api/storage/mysql/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserGetUserRankingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserGetUserRankingService(Context context.Context, RequestContext *app.RequestContext) *UserGetUserRankingService {
	return &UserGetUserRankingService{RequestContext: RequestContext, Context: Context}
}

func (h *UserGetUserRankingService) Run(req *user.GetUserRankingRequest) (resp *user.GetUserRankingResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 从context获取用户参数
	var userID int64 = 0
	contextParams, err := GetContextParams(h.Context, h.RequestContext)
	if err == nil {
		userID = contextParams.UserID
	}
	// 排名接口允许未登录用户访问，所以不返回错误

	// 验证排名类型参数
	rankType := req.Type
	if rankType != "daily" && rankType != "weekly" && rankType != "all" {
		rankType = "all" // 默认为总排名
	}

	var users []*model.User
	var userRanking int64 = 0

	// 根据排名类型使用不同的数据源
	switch rankType {
	case "daily":
		// 日榜：从Redis获取当天积分排名
		users, userRanking, err = dal.GetDailyRankingFromRedis(h.Context, userID, 100)
		if err != nil {
			return nil, err
		}
	case "weekly":
		// 周榜：从Redis获取周积分排名
		users, userRanking, err = dal.GetWeeklyRankingFromRedis(h.Context, userID, 100)
		if err != nil {
			return nil, err
		}
	case "all":
		// 总榜：从MySQL获取总积分排名
		users, err = dal.GetUserRanking(h.Context, rankType, 100)
		if err != nil {
			return nil, err
		}

		// 获取用户在总榜中的排名
		if userID > 0 {
			userRanking, err = dal.GetUserRankingPosition(h.Context, userID, rankType)
			if err != nil {
				userRanking = 1000 // 如果获取失败，设置为1000
			}
			if userRanking > 1000 {
				userRanking = 1000
			}
		}
	}

	// 转换为响应格式
	var rankingList []*user.UserRanking
	for _, u := range users {
		rankingList = append(rankingList, &user.UserRanking{
			Address:  u.Address,
			UserName: u.UserName,
			Points:   u.Points,
		})
	}

	resp = &user.GetUserRankingResponse{
		UserRanking: userRanking,
		RankingList: rankingList,
	}
	return
}
