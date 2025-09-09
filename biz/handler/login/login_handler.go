package login

import (
	"context"

	"Ephyra-genesis-api/biz/service"
	"Ephyra-genesis-api/biz/utils"
	login "Ephyra-genesis-api/hertz_gen/basic/login"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// LoginUserLogin .
// @router /api/login/ [POST]
func LoginUserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req login.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &login.UserLoginResponse{}
	resp, err = service.NewLoginUserLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
