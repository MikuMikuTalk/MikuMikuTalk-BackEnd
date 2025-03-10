// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"im_server/im_file/file_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 图片预览服务
				Method:  http.MethodGet,
				Path:    "/api/file/:imageName",
				Handler: ImagePreviewHandler(serverCtx),
			},
			{
				// 文件下载服务
				Method:  http.MethodGet,
				Path:    "/api/file/download/:fileName",
				Handler: FileDownloadHandler(serverCtx),
			},
			{
				// 文件上传服务
				Method:  http.MethodPost,
				Path:    "/api/file/file",
				Handler: FileHandler(serverCtx),
			},
			{
				// 图片上传服务
				Method:  http.MethodPost,
				Path:    "/api/file/image",
				Handler: ImageHandler(serverCtx),
			},
		},
	)
}
