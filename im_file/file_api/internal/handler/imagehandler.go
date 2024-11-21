package handler

import (
	"errors"
	"fmt"
	"im_server/common/response"
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
			response.Response(r, w, nil, err)
			return
		}
		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		imageType := r.FormValue("imageType")
		if imageType == "" {
			response.Response(r, w, nil, errors.New("imageType不能为空"))
			return
		}
		fileName := fileHeader.Filename
		err = os.MkdirAll(fmt.Sprintf("uploads/%s", imageType), os.ModePerm)
		if err != nil {
			response.Response(r, w, nil, errors.New("文件夹创建失败"))
			return
		}
		filePath := path.Join("uploads", imageType, fileName)
		outFile, err := os.Create(filePath)
		defer outFile.Close() // 确保文件在函数结束时关闭
		if err != nil {
			response.Response(r, w, nil, errors.New("文件创建失败"))
			return
		}
		_, err = io.Copy(outFile, file)
		if err != nil {
			response.Response(r, w, nil, errors.New("文件创建失败"))
			return
		}

		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		resp.Url = "/" + filePath
		// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		response.Response(r, w, resp, err)

	}
}
