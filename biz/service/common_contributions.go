package service

import (
	"context"

	"Ephyra-genesis-api/biz/dal"
	common "Ephyra-genesis-api/hertz_gen/basic/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type CommonContributionsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCommonContributionsService(Context context.Context, RequestContext *app.RequestContext) *CommonContributionsService {
	return &CommonContributionsService{RequestContext: RequestContext, Context: Context}
}

func (h *CommonContributionsService) Run(req *common.ContributionsRequest) (resp *common.ContributionsResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 获取 user_daily_task_log 表的总数
	count, err := dal.GetUserDailyTaskLogCount(h.Context)
	if err != nil {
		return nil, err
	}

	resp = &common.ContributionsResponse{
		ContributionNum: count,
	}
	return
}
