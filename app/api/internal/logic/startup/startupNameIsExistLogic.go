package startup

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	startupModel "metaLand/data/model/startup"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartupNameIsExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据StartupName查询starup是否存在
func NewStartupNameIsExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartupNameIsExistLogic {
	return &StartupNameIsExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartupNameIsExistLogic) StartupNameIsExist(req *types.StartupNameRequest) (resp *types.ExistResponse, err error) {
	// todo: add your logic here and delete this line
	exist, err := startupModel.StartupNameIsExist(l.svcCtx.DB, req.Name)
	if err != nil {
		return nil, err
	}
	resp = &types.ExistResponse{
		IsExist: exist,
	}
	return resp, nil
}
