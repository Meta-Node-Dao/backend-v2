package startup

import (
	"context"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartupFollowedByComerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据startup是否被follow
func NewStartupFollowedByComerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartupFollowedByComerLogic {
	return &StartupFollowedByComerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartupFollowedByComerLogic) StartupFollowedByComer(req *types.FollowRelationRequest) (resp *types.ExistResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
