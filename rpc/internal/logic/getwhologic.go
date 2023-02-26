package logic

import (
	"context"

	"myCicdTest/rpc/internal/svc"
	"myCicdTest/rpc/types/ops"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWhoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWhoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWhoLogic {
	return &GetWhoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWhoLogic) GetWho(in *ops.GetWhoReq) (*ops.GetWhoReply, error) {
	msg := "我是大哥大!"

	return &ops.GetWhoReply{Msg: msg}, nil
}
