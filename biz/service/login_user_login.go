package service

import (
	"Ephyra-genesis-api/biz/bizerror"
	"Ephyra-genesis-api/biz/dal"
	jwt "Ephyra-genesis-api/biz/mw"
	"Ephyra-genesis-api/biz/utils"
	"Ephyra-genesis-api/pkg/web3"
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	login "Ephyra-genesis-api/hertz_gen/basic/login"

	"github.com/cloudwego/hertz/pkg/app"
)

type LoginUserLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginUserLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginUserLoginService {
	return &LoginUserLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginUserLoginService) Run(req *login.UserLoginRequest) (resp *login.UserLoginResponse, err error) {

	// 验证钱包签名消息正确
	ret, err := web3.VerifyEthSignature(h.Context, req.SignMessage, req.Signature, req.Address)
	if !ret || err != nil {
		return nil, bizerror.InvalidSignatureError
	}

	// 查询用户钱包是否已经注册
	userModel, err := dal.GetUserByAddress(h.Context, req.Address)
	if err != nil {
		return nil, bizerror.DBError
	}

	isNew := false
	// 用户不存在
	if userModel == nil {
		// 创建跟用户相关的表记录
		userModel, err = dal.CreateNewUser(h.Context, req.Address)
		if err != nil {
			return nil, err
		}
		isNew = true
	}

	// 生成jwt token
	token, err := jwt.JwtMiddleware.CreateTokenWithVersion(
		jwt.BscVersion,
		jwt.NewJWTPayload(userModel.ID, req.Address, userModel.CreateAt.Unix(), req.ChainName))
	if err != nil {
		// 生成jwt token 不应该出错
		return nil, bizerror.InvalidRequestError
	}

	// 生成md5
	tokenMd5 := utils.MD5(token)

	// jwt token 放到缓存中
	if _, err = jwt.JwtMiddleware.SetJwtToken(h.Context, userModel.ID, req.Address, tokenMd5); err != nil {
		return nil, bizerror.DBError
	}

	// 返回
	resp = &login.UserLoginResponse{
		AuthToken: token,
		UserId:    userModel.ID,
		Address:   userModel.Address,
		IsNew:     isNew,
	}
	hlog.CtxInfof(h.Context, "[UserLoginService] userId: %v, tokenMd5: %v, addressFmt: %v, signMessage: %v, signature: %v, chain: %v",
		userModel.ID, tokenMd5, req.Address, req.SignMessage, req.Signature, req.ChainName)

	return
}
