package service

import (
	"context"

	"Ephyra-genesis-api/biz/dal"
	user "Ephyra-genesis-api/hertz_gen/basic/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserSetUserProfileService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserSetUserProfileService(Context context.Context, RequestContext *app.RequestContext) *UserSetUserProfileService {
	return &UserSetUserProfileService{RequestContext: RequestContext, Context: Context}
}

func (h *UserSetUserProfileService) Run(req *user.SetUserProfileRequest) (resp *user.SetUserProfileResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// 从context获取用户参数
	contextParams, err := GetContextParams(h.Context, h.RequestContext)
	if err != nil {
		resp = &user.SetUserProfileResponse{
			Status: false,
		}
		return
	}
	userID := contextParams.UserID

	// 更新用户档案
	err = dal.UpdateUserProfile(h.Context, userID, req.UserName, req.AvatarUrl)
	if err != nil {
		resp = &user.SetUserProfileResponse{
			Status: false,
		}
		return
	}

	resp = &user.SetUserProfileResponse{
		Status: true,
	}
	return
}
