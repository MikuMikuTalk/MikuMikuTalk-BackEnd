package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"im_server/common/response"
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileRequest
		if err := httpx.Parse(r, &req); err != nil {
			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
			responseError(r, w, err)
			return
		}
		files := r.MultipartForm.File["files"]
		var image_lists []string = make([]string, 0)
		var dirName string
		var filePath string
		for _, fileHeader := range files {
			fileName := fileHeader.Filename
			if !validateFileSize(fileHeader.Size, svcCtx.Config.FileSize) {
				responseError(r, w, fmt.Errorf("图片大小超过限制，最大只能上传%.2fMB大小的图片", svcCtx.Config.FileSize))
				return
			}

			// 硬编码，存储到uploads/files目录
			dirName = filepath.Join("uploads", "files")
			filePath = filepath.Join(dirName, fileName)
			if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
				responseError(r, w, errors.New("文件夹创建失败"))
				return
			}
			// 打开文件
			file, err := fileHeader.Open()
			if err != nil {
				responseError(r, w, err)
				return
			}
			if isFileInDirectory(dirName, fileName) {
				// 检查文件内容是否重复
				if isDuplicateFile(filePath, file) {
					responseError(r, w, errors.New("不要上传重复图片"))
					return
				}

				// 如果文件内容不同，生成新文件名
				fileName = renameFile(fileName)
				filePath = filepath.Join(dirName, fileName)
			}
			// 保存图片
			if err := saveFile(filePath, file); err != nil {
				responseError(r, w, errors.New("文件保存失败"))
				return
			}
			image_lists = append(image_lists, "/"+filePath)
		}

		l := logic.NewFileLogic(r.Context(), svcCtx)
		resp, err := l.File(&req)
		if err == nil {
			resp.Src = image_lists
		}
		// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		response.Response(r, w, resp, err)
	}
}
