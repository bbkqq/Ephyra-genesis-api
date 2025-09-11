package service

import (
	"context"

	"Ephyra-genesis-api/biz/dal"
	common "Ephyra-genesis-api/hertz_gen/basic/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type CommonNFTMintedInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCommonNFTMintedInfoService(Context context.Context, RequestContext *app.RequestContext) *CommonNFTMintedInfoService {
	return &CommonNFTMintedInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *CommonNFTMintedInfoService) Run(req *common.NFTMintedInfoRequest) (resp *common.NFTMintedInfoResponse, err error) {

	users, err := dal.GetLatestUsersWithSBTTokenID(h.Context, 10)
	if err != nil {
		hlog.CtxErrorf(h.Context, "GetLatestUsersWithSBTTokenID error: %v", err)
		return nil, err
	}

	var infos []*common.NFTInfo
	for _, user := range users {
		infos = append(infos, &common.NFTInfo{
			UserAddress: user.Address,
			NftTokenId:  user.SbtTokenID,
			MintTime:    user.CreateAt.Unix(),
		})
	}

	resp = &common.NFTMintedInfoResponse{}
	resp.Info = infos
	return
}
