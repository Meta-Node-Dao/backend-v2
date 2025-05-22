package startup

import (
	"context"
	"database/sql"
	"errors"
	"metaLand/data/model/startup"
	"time"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStartupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取项目详情

func NewGetStartupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStartupInfoLogic {
	return &GetStartupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStartupInfoLogic) GetStartupInfo(req *types.EmptyRequest) (resp *types.StartupInfoResponse, err error) {

	//TODO 从路径参数中获取 startupId  startupId  一直没有取到值

	startupId, ok := l.ctx.Value("startupId").(uint64)
	if !ok {
		return nil, errors.New("invalid startupId")
	}

	// 查询项目详情
	startupInfo, err := startup.GetStartupInfo(l.svcCtx.DB, &startupId)

	if err != nil {
		return nil, err
	}

	data := &types.StartupRes{
		ComerID:              startupInfo.ComerID,
		Name:                 startupInfo.Name,
		Mode:                 startupInfo.Mode,
		Logo:                 startupInfo.Logo,
		Cover:                startupInfo.Cover,
		Mission:              startupInfo.Mission,
		TokenContractAddress: startupInfo.TokenContractAddress,
		Overview:             startupInfo.Overview,
		TxHash:               startupInfo.TxHash,
		OnChain:              startupInfo.OnChain,
		KYC:                  startupInfo.KYC,
		ContractAudit:        startupInfo.ContractAudit,
		Website:              startupInfo.Website,
		Discord:              startupInfo.Discord,
		Twitter:              startupInfo.Twitter,
		Telegram:             startupInfo.Telegram,
		Docs:                 startupInfo.Docs,
		Email:                startupInfo.Email,
		Facebook:             startupInfo.Facebook,
		Medium:               startupInfo.Medium,
		Linktree:             startupInfo.Linktree,
		LaunchNetwork:        startupInfo.LaunchNetwork,
		TokenName:            startupInfo.TokenName,
		TokenSymbol:          startupInfo.TokenSymbol,
		TotalSupply:          startupInfo.TotalSupply,
		PresaleStart:         nullTimeToString(startupInfo.PresaleStart),
		PresaleEnd:           nullTimeToString(startupInfo.PresaleEnd),
		LaunchDate:           nullTimeToString(startupInfo.LaunchDate),
		TabSequence:          startupInfo.TabSequence,
		IsDeleted:            startupInfo.IsDeleted,
	}

	// 返回成功响应
	return &types.StartupInfoResponse{
		Code:    200,
		Message: "查询成功",
		Data:    *data,
	}, nil
}

// 将sql.NullTime转换为字符串时间格式（RFC3339）
func nullTimeToString(nt sql.NullTime) string {
	if nt.Valid {
		return nt.Time.Format(time.RFC3339)
	}
	return ""
}
