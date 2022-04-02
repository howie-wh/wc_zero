// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"wczero/services/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/user/wechat_login",
				Handler: WeChatLoginHandler(serverCtx),
			},
		},
	)
}
