package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero-mall/service/user/api/internal/logic"
	"gozero-mall/service/user/api/internal/svc"
	"gozero-mall/service/user/api/internal/types"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
