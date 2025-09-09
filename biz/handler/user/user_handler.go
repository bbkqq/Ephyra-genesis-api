package user

import (
	"context"

	"Ephyra-genesis-api/biz/service"
	"Ephyra-genesis-api/biz/utils"
	user "Ephyra-genesis-api/hertz_gen/basic/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserGetUserProfile .
// @router /api/user/get-user-profile [GET]
func UserGetUserProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.GetUserProfileResponse{}
	resp, err = service.NewUserGetUserProfileService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserSetUserProfile .
// @router /api/user/set-user-profile [POST]
func UserSetUserProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.SetUserProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.SetUserProfileResponse{}
	resp, err = service.NewUserSetUserProfileService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserGetUserRanking .
// @router /api/user/get-user-ranking [GET]
func UserGetUserRanking(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserRankingRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &user.GetUserRankingResponse{}
	resp, err = service.NewUserGetUserRankingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
