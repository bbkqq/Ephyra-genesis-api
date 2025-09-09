package task

import (
	"context"

	"Ephyra-genesis-api/biz/service"
	"Ephyra-genesis-api/biz/utils"
	task "Ephyra-genesis-api/hertz_gen/basic/task"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// TaskGetDailyQuestions .
// @router /api/task/daily-questions/ [GET]
func TaskGetDailyQuestions(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.GetDailyQuestionsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &task.GetDailyQuestionsResponse{}
	resp, err = service.NewTaskGetDailyQuestionsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// TaskCanSubmitAnswer .
// @router /api/task/can-submit-answer/ [POST]
func TaskCanSubmitAnswer(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.CanSubmitAnswerRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &task.CanSubmitAnswerResponse{}
	resp, err = service.NewTaskCanSubmitAnswerService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// TaskGetAnswerTxStatus .
// @router /api/task/get-answer-tx-status/ [POST]
func TaskGetAnswerTxStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.GetAnswerTxStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &task.GetAnswerTxStatusResponse{}
	resp, err = service.NewTaskGetAnswerTxStatusService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
