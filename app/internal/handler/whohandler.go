package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myCicdTest/app/internal/logic"
	"myCicdTest/app/internal/svc"
)

func WhoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewWhoLogic(r.Context(), svcCtx)
		resp, err := l.Who()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
