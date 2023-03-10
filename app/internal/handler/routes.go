// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"myCicdTest/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/who",
				Handler: WhoHandler(serverCtx),
			},
		},
	)
}
