package handler

import (
	"errors"
	"net/http"
	"os"

	"im_server/common/response"
	"im_server/im_file/file_api/file_models"
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImagePreviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImagePreviewRequest
		if err := httpx.Parse(r, &req); err != nil {
			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
			response.Response(r, w, nil, err)
			return
		}
		var fileModel file_models.FileModel
		err := svcCtx.DB.Take(&fileModel, "uid = ?", req.ImageName).Error
		if err != nil {
			response.Response(r, w, nil, errors.New("文件失败"))
			return
		}
		byteData, err := os.ReadFile(fileModel.Path)
		if err != nil {
			//读取文件失败
			response.Response(r, w, nil, errors.New("读取文件失败"))
			return
		}
		w.Write(byteData)
		l := logic.NewImagePreviewLogic(r.Context(), svcCtx)
		err = l.ImagePreview(&req)
		if err != nil {
			response.Response(r, w, nil, errors.New("读取文件失败"))
			return
		}
	}
}
