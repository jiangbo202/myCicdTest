// Code generated by goctl. DO NOT EDIT!
// Source: ops-rpc.proto

package server

import (
	"context"

	"myCicdTest/rpc/internal/logic"
	"myCicdTest/rpc/internal/svc"
	"myCicdTest/rpc/types/ops"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	ops.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetWho(ctx context.Context, in *ops.GetWhoReq) (*ops.GetWhoReply, error) {
	l := logic.NewGetWhoLogic(ctx, s.svcCtx)
	return l.GetWho(in)
}
