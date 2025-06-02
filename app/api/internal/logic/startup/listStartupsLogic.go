package startup

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	startupModel "metaLand/data/model/startup"

	"github.com/qiniu/x/log"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListStartupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询项目列表
func NewListStartupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStartupsLogic {
	return &ListStartupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStartupsLogic) ListStartups(req *types.ListStartupsRequest) (resp *types.ListStartupsResponse, err error) {
	const comerID = 1001
	request := &startupModel.ListStartupRequest{
		Limit:     req.Limit,
		Offset:    req.Offset,
		Keyword:   req.Keyword,
		IsDeleted: req.IsDeleted,
		Mode:      req.Mode,
	}
	var startups []startupModel.Startup
	total, err := startupModel.ListStartups(l.svcCtx.DB, comerID, request, &startups)
	if err != nil {
		log.Warn(err)
		return
	}
	if total == 0 {
		resp.List = make([]*types.Startup, 0)
		return
	}
	var list []*types.Startup
	for _, startup := range startups {
		list = append(list, &types.Startup{
			ComerID:              startup.ComerID,
			Name:                 startup.Name,
			Mode:                 uint8(startup.Mode),
			Logo:                 startup.Logo,
			Cover:                startup.Cover,
			Mission:              startup.Mission,
			TokenContractAddress: startup.TokenContractAddress,
			Overview:             startup.Overview,
			ChainID:              startup.ChainID,
			TxHash:               startup.TxHash,
			OnChain:              startup.OnChain,
			KYC:                  startup.KYC,
			ContractAudit:        startup.ContractAudit,
		})
	}
	return &types.ListStartupsResponse{
		List:  list,
		Total: total,
	}, nil
}
