package logic

import (
	"context"
	"myCicdTest/rpc/user"

	"myCicdTest/app/internal/svc"
	"myCicdTest/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WhoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWhoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhoLogic {
	return &WhoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WhoLogic) Who() (resp *types.UserWhoReply, err error) {
	res, err := l.svcCtx.MyRpc.GetWho(l.ctx, &user.GetWhoReq{})
	return &types.UserWhoReply{Msg: res.Msg}, nil
}
