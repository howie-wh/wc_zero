package handler

import (
	"net/http"
	"wczero/common/response" //1

	"github.com/zeromicro/go-zero/rest/httpx"
	"wczero/services/wallpaper/api/internal/logic"
	"wczero/services/wallpaper/api/internal/svc"
	"wczero/services/wallpaper/api/internal/types"
)

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRemoveLogic(r.Context(), svcCtx)
		resp, err := l.Remove(req)
		response.Response(w, resp, err) //2
	}
}
