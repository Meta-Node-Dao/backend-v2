package startup

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	startupModel "metaLand/data/model/startup"
	tagModel "metaLand/data/model/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStartupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建项目
func NewCreateStartupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStartupLogic {
	return &CreateStartupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStartupLogic) CreateStartup(req *types.CreateStartupRequest) (resp *types.CreateStartupResponse, err error) {
	// todo: add your logic here and delete this line
	var tag tagModel.Tag
	hashTags := []tagModel.Tag{tag}

	request := &startupModel.Startup{
		Logo:     req.Logo,
		Mode:     startupModel.Mode(req.Mode),
		Name:     req.Name,
		Mission:  req.Mission,
		Overview: req.Overview,
		TxHash:   req.TxHash,
		ChainID:  req.ChainID,
		HashTags: hashTags,
	}

	err = startupModel.CreateStartup(l.svcCtx.DB, request)
	if err != nil {
		return nil, err
	}
	// TODO: 不太清楚ID是如何生成的，这里先hard code一个ID
	var id uint64 = 123
	return &types.CreateStartupResponse{
		ID: id,
	}, nil
}
