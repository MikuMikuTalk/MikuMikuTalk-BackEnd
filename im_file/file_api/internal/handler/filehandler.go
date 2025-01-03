package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"im_server/common/response"
	"im_server/im_file/file_api/file_models"
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"
	"im_server/utils/jwts"
	"im_server/utils/md5_util"

	"github.com/google/uuid"
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
		token := r.Header.Get("Authorization")
		claims, err := jwts.ParseToken(token, svcCtx.Config.Auth.AuthSecret)
		if err != nil {
			responseError(r, w, err)
			return
		}
		my_id := claims.UserID

		files := r.MultipartForm.File["files"]
		srcLists := make([]string, 0)
		var dirName string
		var filePath string
		for _, fileHeader := range files {

			fileName := fileHeader.Filename
			if !validateFileSize(fileHeader.Size, svcCtx.Config.FileSize) {
				responseError(r, w, fmt.Errorf("图片大小超过限制，最大只能上传%.2fMB大小的文件", svcCtx.Config.FileSize))
				return
			}

			dirName = filepath.Join(svcCtx.Config.UploadDir, "files")
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
					responseError(r, w, errors.New("不要上传重复文件"))
					return
				}

				// 如果文件内容不同，生成新文件名
				fileName = renameFile(fileName)
				filePath = filepath.Join(dirName, fileName)
			}
			// 解决因为文件指针问题导致的存储文件占用空间大小为0且hash值一样的问题
			buf := new(bytes.Buffer)
			teeReader := io.TeeReader(file, buf)
			// 保存文件
			if err := saveFile(filePath, teeReader); err != nil {
				responseError(r, w, errors.New("文件保存失败"))
				return
			}
			hash, err := md5_util.ComputeMD5(buf)
			if err != nil {
				responseError(r, w, errors.New("计算Hash失败"))
				return
			}
			fileModel := file_models.FileModel{
				UserID:   my_id,
				FileName: fileName,
				Size:     fileHeader.Size,
				Path:     filePath,
				Hash:     hash,
				Uid:      uuid.New(),
			}
			// 创建数据库记录
			err = svcCtx.DB.Create(&fileModel).Error
			if err != nil {
				response.Response(r, w, nil, errors.New("创建数据库记录失败"))
				return
			}
			srcLists = append(srcLists, fileModel.FileWebPath())
		}

		l := logic.NewFileLogic(r.Context(), svcCtx)
		resp, err := l.File(&req)
		if err == nil {
			resp.Src = srcLists
		}
		// 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
		response.Response(r, w, resp, err)
	}
}
