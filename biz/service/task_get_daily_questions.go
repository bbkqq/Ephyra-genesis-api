package service

import (
	"context"

	"Ephyra-genesis-api/biz/dal"
	task "Ephyra-genesis-api/hertz_gen/basic/task"

	"github.com/cloudwego/hertz/pkg/app"
)

type TaskGetDailyQuestionsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewTaskGetDailyQuestionsService(Context context.Context, RequestContext *app.RequestContext) *TaskGetDailyQuestionsService {
	return &TaskGetDailyQuestionsService{RequestContext: RequestContext, Context: Context}
}

func (h *TaskGetDailyQuestionsService) Run(req *task.GetDailyQuestionsRequest) (resp *task.GetDailyQuestionsResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 获取每日问题，这里获取所有问题，实际项目中可能需要根据日期或其他条件过滤
	questions, err := dal.GetDailyQuestions(h.Context, 0) // 0表示不限制数量
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	var responseQuestions []*task.Question
	for _, q := range questions {
		responseQuestions = append(responseQuestions, &task.Question{
			Id:      q.ID,
			Title:   q.Title,
			Content: q.Content,
			Type:    int64(q.Type),
		})
	}

	resp = &task.GetDailyQuestionsResponse{
		Questions: responseQuestions,
	}
	return
}
