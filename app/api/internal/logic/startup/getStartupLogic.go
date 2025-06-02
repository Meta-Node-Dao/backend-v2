package startup

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	startupModel "metaLand/data/model/startup"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStartupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据StartupID查询starup信息
func NewGetStartupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStartupLogic {
	return &GetStartupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStartupLogic) GetStartup(req *types.GetStartupRequest) (resp *types.GetStartupResponse, err error) {
	// todo: add your logic here and delete this line
	var startupId uint64 = req.ID
	var startup startupModel.Startup
	err = startupModel.GetStartup(l.svcCtx.DB, startupId, &startup)
	if err != nil {
		return nil, err
	}
	Startup := &types.Startup{

		Name:     startup.Name,
		Mission:  startup.Mission,
		Overview: startup.Overview,
		Logo:     startup.Logo,
		Mode:     uint8(startup.Mode),
		TxHash:   startup.TxHash,
		ChainID:  startup.ChainID,
		// HashTags: startup.HashTags,
		// CreatedAt:   startup.CreatedAt,
		// UpdatedAt:   startup.UpdatedAt,
		// FollowCount: startup.FollowCount,
		// Followed: false, // 这里需要根据实际情况设置followed状态
	}

	return &types.GetStartupResponse{
		Startup: Startup,
	}, nil
}
