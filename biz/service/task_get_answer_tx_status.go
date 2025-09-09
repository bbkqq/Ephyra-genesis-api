package service

import (
	"Ephyra-genesis-api/biz/dal"
	task "Ephyra-genesis-api/hertz_gen/basic/task"
	user "Ephyra-genesis-api/hertz_gen/basic/user"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type TaskGetAnswerTxStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewTaskGetAnswerTxStatusService(Context context.Context, RequestContext *app.RequestContext) *TaskGetAnswerTxStatusService {
	return &TaskGetAnswerTxStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *TaskGetAnswerTxStatusService) Run(req *task.GetAnswerTxStatusRequest) (resp *task.GetAnswerTxStatusResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 检查 TxHash 参数
	if len(req.TxHash) == 0 {
		resp = &task.GetAnswerTxStatusResponse{
			Status:      0, // 未处理
			Badges:      []*user.Badge{},
			Points:      0,
			ExtraPoints: 0,
		}
		return
	}

	// 根据 txhash 获取任务日志
	taskLog, err := dal.GetUserDailyTaskLogByTxHash(h.Context, req.TxHash)
	if err != nil {
		return nil, err
	}

	// 如果没有任务日志，说明用户今天没有参与
	if taskLog == nil {
		resp = &task.GetAnswerTxStatusResponse{
			Status:      0, // 未处理
			Badges:      []*user.Badge{},
			Points:      0,
			ExtraPoints: 0,
		}
		return
	}

	// 解析用户获得的徽章
	badgeIDs, err := dal.ParseUserBadges(taskLog.Badges)
	if err != nil {
		hlog.CtxErrorf(h.Context, "ParseUserBadges error: %v", err)
		return nil, err
	}

	// 获取徽章详情
	badges, err := dal.GetBadgesByIDs(h.Context, badgeIDs)
	if err != nil {
		hlog.CtxErrorf(h.Context, "GetBadgesByIDs error: %v", err)
		return nil, err
	}

	// 转换为响应格式
	var responseBadges []*user.Badge
	for _, badge := range badges {
		responseBadges = append(responseBadges, &user.Badge{
			Id:               badge.ID,
			Name:             badge.Name,
			ImageUrl:         badge.ImageURL,
			Unlocked:         true, // 既然在任务日志中，说明已解锁
			UnlockConditions: badge.UnlockConditions,
		})
	}

	// 从任务日志中直接获取积分信息
	basePoints := int64(taskLog.Points)
	extraPoints := int64(taskLog.ExtraPoints)

	resp = &task.GetAnswerTxStatusResponse{
		Status:      1, // 成功
		Badges:      responseBadges,
		Points:      basePoints,
		ExtraPoints: extraPoints,
	}
	return
}
