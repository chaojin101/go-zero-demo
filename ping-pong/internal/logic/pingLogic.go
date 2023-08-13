package logic

import (
	"context"

	"ping-pong/internal/svc"
	"ping-pong/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.Req) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return &types.Resp{
		Message: "pong",
	}, nil
}
