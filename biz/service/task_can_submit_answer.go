package service

import (
	"context"
	"time"

	"Ephyra-genesis-api/biz/dal"
	task "Ephyra-genesis-api/hertz_gen/basic/task"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type TaskCanSubmitAnswerService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewTaskCanSubmitAnswerService(Context context.Context, RequestContext *app.RequestContext) *TaskCanSubmitAnswerService {
	return &TaskCanSubmitAnswerService{RequestContext: RequestContext, Context: Context}
}

func (h *TaskCanSubmitAnswerService) Run(req *task.CanSubmitAnswerRequest) (resp *task.CanSubmitAnswerResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 从context获取用户参数
	contextParams, err := GetContextParams(h.Context, h.RequestContext)
	if err != nil {
		hlog.CtxErrorf(h.Context, "GetContextParams error: %v", err)
		// 用户未认证或参数错误
		resp = &task.CanSubmitAnswerResponse{
			Status: false,
		}
		return
	}
	userID := contextParams.UserID

	// 使用新的计算天数偏移函数
	currentTime := time.Now()
	daysDiff := dal.CalculateDayOffset(currentTime)

	// 检查用户今天是否已经提交过答案
	existingLog, err := dal.GetUserDailyTaskLogByUserAndDay(h.Context, userID, daysDiff)
	if err != nil {
		hlog.CtxErrorf(h.Context, "GetUserDailyTaskLogByUserAndDay error: %v", err)
		resp = &task.CanSubmitAnswerResponse{
			Status: false,
		}
		return resp, nil
	}

	// 如果今天已经提交过答案，则不能再提交
	canSubmit := existingLog == nil
	resp = &task.CanSubmitAnswerResponse{
		Status: canSubmit,
	}

	if canSubmit {
		user, err := dal.GetUserByID(h.Context, userID)
		if err != nil {
			hlog.CtxErrorf(h.Context, "GetUserByID error: %v", err)
			resp = &task.CanSubmitAnswerResponse{
				Status: false,
			}
			return resp, nil
		}
		reward, err := dal.CalculateTaskReward(h.Context, user, daysDiff)
		if err != nil {
			hlog.CtxErrorf(h.Context, "CalculateTaskReward error: %v", err)
			resp = &task.CanSubmitAnswerResponse{
				Status: false,
			}
			return resp, nil
		}

		resp.EstimatedPoints = int64(reward.Points)
		resp.EstimatedExtraPoints = int64(reward.ExtraPoints)
	}

	return
}
