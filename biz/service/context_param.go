package service

import (
	"Ephyra-genesis-api/biz/bizerror"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/cast"

	"reflect"
)

type ContextParams struct {
	UserID     int64
	Address    string
	Chain      string
	JWTSkipped bool
}

func GetContextParams(ctx context.Context, appCtx *app.RequestContext) (*ContextParams, error) {
	ctp := &ContextParams{
		UserID:     0,
		Address:    "",
		Chain:      "",
		JWTSkipped: false,
	}

	if v, ok := appCtx.Get("jwt_skipped"); ok {
		r := cast.ToBool(v)
		if r {
			ctp.JWTSkipped = true
			return ctp, nil
		}
	}

	v, ok := appCtx.Get("user_id")
	if !ok {
		hlog.CtxErrorf(ctx, "user_id not found in context")
		return ctp, bizerror.InvalidParamsError
	}
	userID, err := cast.ToInt64E(v)
	if err != nil {
		hlog.CtxErrorf(ctx, "user_id is not int64, user_id: %v, type %+v", v, reflect.TypeOf(v))
		return ctp, bizerror.InvalidParamsError
	}
	ctp.UserID = userID

	// 获取用户地址
	addr, ok := appCtx.Get("address")
	if !ok {
		hlog.CtxErrorf(ctx, "address not found in context: %+v", appCtx)
		err = bizerror.InvalidParamsError
		return ctp, bizerror.InvalidParamsError
	}
	address, err := cast.ToStringE(addr)
	if err != nil {
		hlog.CtxErrorf(ctx, "address is not string, address: %v ", addr)
		return ctp, bizerror.InvalidParamsError
	}
	ctp.Address = address
	// chain
	ch, ok := appCtx.Get("chain")
	if !ok {
		hlog.CtxErrorf(ctx, "chain not found in context: %+v", appCtx)
		err = bizerror.InvalidParamsError
		return ctp, bizerror.InvalidParamsError
	}
	chain, err := cast.ToStringE(ch)
	if err != nil {
		hlog.CtxErrorf(ctx, "chain is not string, chain: %v ", chain)
		return ctp, bizerror.InvalidParamsError
	}
	ctp.Chain = chain

	return ctp, nil
}
