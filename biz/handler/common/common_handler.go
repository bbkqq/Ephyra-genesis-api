package common

import (
	"context"

	"Ephyra-genesis-api/biz/service"
	"Ephyra-genesis-api/biz/utils"
	common "Ephyra-genesis-api/hertz_gen/basic/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommonContributions .
// @router /api/common/contributions/ [GET]
func CommonContributions(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.ContributionsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &common.ContributionsResponse{}
	resp, err = service.NewCommonContributionsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CommonNFTMintedInfo .
// @router /api/common/nft-minted-info/ [GET]
func CommonNFTMintedInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.NFTMintedInfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCommonNFTMintedInfoService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
