package startup

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	startupModel "metaLand/data/model/startup"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartupTokenContractIsExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据TokenContract查询starup是否存在
func NewStartupTokenContractIsExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartupTokenContractIsExistLogic {
	return &StartupTokenContractIsExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartupTokenContractIsExistLogic) StartupTokenContractIsExist(req *types.TokenContractReqest) (resp *types.ExistResponse, err error) {
	exist, err := startupModel.StartupTokenContractIsExist(l.svcCtx.DB, req.TokenContract)
	if err != nil {
		return nil, err
	}
	resp = &types.ExistResponse{
		IsExist: exist,
	}
	return resp, nil
}
