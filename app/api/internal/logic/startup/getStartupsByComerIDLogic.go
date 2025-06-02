package startup

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStartupsByComerIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据ComerID查询项目
func NewGetStartupsByComerIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStartupsByComerIDLogic {
	return &GetStartupsByComerIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStartupsByComerIDLogic) GetStartupsByComerID(req *types.GetStartupsByComerIDRequest) (resp *types.GetStartupsByComerIDResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
