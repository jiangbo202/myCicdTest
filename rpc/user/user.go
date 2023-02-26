// Code generated by goctl. DO NOT EDIT!
// Source: ops-rpc.proto

package user

import (
	"context"

	"myCicdTest/rpc/types/ops"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetWhoReply = ops.GetWhoReply
	GetWhoReq   = ops.GetWhoReq

	User interface {
		GetWho(ctx context.Context, in *GetWhoReq, opts ...grpc.CallOption) (*GetWhoReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetWho(ctx context.Context, in *GetWhoReq, opts ...grpc.CallOption) (*GetWhoReply, error) {
	client := ops.NewUserClient(m.cli.Conn())
	return client.GetWho(ctx, in, opts...)
}