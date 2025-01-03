package handler

import (
	"errors"
	"io"
	"net/http"
	"os"

	"im_server/common/response"
	"im_server/im_file/file_api/file_models"
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
			response.Response(r, w, nil, err)
			return
		}
		var fileModel file_models.FileModel
		err := svcCtx.DB.Take(&fileModel, "uid = ?", req.FileName).Error
		if err != nil {
			response.Response(r, w, nil, errors.New("文件不存在"))
			return
		}
		file, err := os.Open(fileModel.Path)
		if err != nil {
			response.Response(r, w, nil, errors.New("文件打开失败"))
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			response.Response(r, w, nil, errors.New("文件写入失败"))
		}
		l := logic.NewFileDownloadLogic(r.Context(), svcCtx)
		resp, err := l.FileDownload(&req)
		// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		if err != nil {
			response.Response(r, w, resp, err)
			return
		}
	}
}
