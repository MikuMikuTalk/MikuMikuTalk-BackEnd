package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"im_server/common/response"
	"im_server/im_file/file_api/file_models"
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"

	"im_server/utils/file_utils"
	"im_server/utils/jwts"
	"im_server/utils/md5_util"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	var mu sync.Mutex

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest

		if err := httpx.Parse(r, &req); err != nil {
			responseError(r, w, err)
			return
		}
		token := r.Header.Get("Authorization")
		claims, err := jwts.ParseToken(token, svcCtx.Config.Auth.AuthSecret)
		my_id := claims.UserID
		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			responseError(r, w, err)
			return
		}
		defer file.Close()

		imageType := r.FormValue("imageType")
		if imageType == "" {
			responseError(r, w, errors.New("imageType不能为空"))
			return
		}
		// 图片类型白名单
		switch imageType {
		case "avatar", "group_avatar", "chat":
		default:
			response.Response(r, w, nil, errors.New("imageType只能为 avatar,group_avatar,chat"))
			return
		}

		//文件后缀白名单
		imageName := fileHeader.Filename
		imageExtName := file_utils.GetFileExtName(imageName)
		if !validateFileExtension(imageExtName, svcCtx.Config.WhiteList) {
			responseError(r, w, errors.New("不可以上传这种格式的图片！"))
			return
		}
		// 文件大小限制
		if !validateFileSize(fileHeader.Size, svcCtx.Config.FileSize) {
			responseError(r, w, fmt.Errorf("图片大小超过限制，最大只能上传%.2fMB大小的图片", svcCtx.Config.ImageSize))
			return
		}

		dirPath := filepath.Join(svcCtx.Config.UploadDir, imageType)
		_, err = os.ReadDir(dirPath)
		if err != nil {
			logx.Info("文件夹不存在，创建文件夹", dirPath)
			os.MkdirAll(dirPath, os.ModePerm)
		}

		imagePath := filepath.Join(dirPath, imageName)
		mu.Lock() // Lock for concurrent access
		defer mu.Unlock()
		if isFileInDirectory(dirPath, imageName) {
			// 检查文件内容是否重复
			if isDuplicateFile(imagePath, file) {
				responseError(r, w, errors.New("不要上传重复图片"))
				return
			}

			// 如果文件内容不同，生成新文件名
			imageName = renameFile(imageName)
			imagePath = filepath.Join(dirPath, imageName)
		}
		// 保存图片
		if err := saveFile(imagePath, file); err != nil {
			responseError(r, w, errors.New("文件保存失败"))
			return
		}
		imageData, _ := io.ReadAll(file)
		fileModel := file_models.FileModel{
			UserID:   my_id,
			FileName: imageName,
			Size:     fileHeader.Size,
			Path:     imagePath,
			Hash:     md5_util.MD5(imageData),
			Uid:      uuid.New(),
		}

		// 逻辑
		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		if err == nil {
			resp.Url = fileModel.WebPath()
		}
		//创建表
		err = svcCtx.DB.Create(&fileModel).Error
		if err != nil {
			logx.Error(err)
			response.Response(r, w, resp, err)
			return
		}
		response.Response(r, w, resp, err)
	}
}
