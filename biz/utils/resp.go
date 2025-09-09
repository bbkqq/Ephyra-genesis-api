package utils

import (
	"Ephyra-genesis-api/biz/bizerror"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Status bool               `json:"status"`
	Data   interface{}        `json:"data,omitempty"`
	Error  *bizerror.BizError `json:"error,omitempty"`
}

func (r *Response) Reset() {
	r.Status = false
	r.Data = nil
	r.Error = nil
}

var ResponsePool = sync.Pool{
	New: func() interface{} {
		return &Response{}
	},
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	hlog.CtxErrorf(ctx, "handler error:%v", err)

	resp := ResponsePool.Get().(*Response)
	defer func() {
		resp.Reset()
		ResponsePool.Put(resp)
	}()

	resp.Status = false

	var bizErr *bizerror.BizError
	ok := errors.As(err, &bizErr)
	if !ok {
		resp.Error = bizerror.InvalidRequestError
	} else {
		resp.Error = bizErr
	}

	c.JSON(code, resp)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	resp := ResponsePool.Get().(*Response)
	defer func() {
		resp.Reset()
		ResponsePool.Put(resp)
	}()

	resp.Status = true
	resp.Data = data

	c.JSON(code, resp)
}
