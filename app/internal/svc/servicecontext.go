package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"myCicdTest/app/internal/config"
	"myCicdTest/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	MyRpc  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MyRpc:  user.NewUser(zrpc.MustNewClient(c.MyRpc)),
	}
}
