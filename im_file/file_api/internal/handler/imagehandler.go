package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"im_server/common/response"
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"
	"im_server/utils/md5_util"
	"im_server/utils/whitelist"

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

		fileName := fileHeader.Filename
		if !validateFileExtension(fileName, svcCtx.Config.WhiteList) {
			responseError(r, w, errors.New("不可以上传这种格式的图片！"))
			return
		}

		if !validateFileSize(fileHeader.Size, svcCtx.Config.FileSize) {
			responseError(r, w, fmt.Errorf("图片大小超过限制，最大只能上传%.2fMB大小的图片", svcCtx.Config.FileSize))
			return
		}

		dirName := filepath.Join("uploads", imageType)
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			responseError(r, w, errors.New("文件夹创建失败"))
			return
		}

		filePath := filepath.Join(dirName, fileName)
		mu.Lock() // Lock for concurrent access
		defer mu.Unlock()

		// 查看是否为相同图片
		if isDuplicateFile(filePath, file) {
			responseError(r, w, errors.New("不要上传重复图片"))
			return
		}

		// 保存图片
		if err := saveFile(filePath, file); err != nil {
			responseError(r, w, errors.New("文件保存失败"))
			return
		}

		// 逻辑
		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		if err == nil {
			resp.Url = "/" + filePath
		}

		response.Response(r, w, resp, err)
	}
}

func responseError(r *http.Request, w http.ResponseWriter, err error) {
	response.Response(r, w, nil, err)
}

func validateFileExtension(fileName string, whiteList []string) bool {
	suffix := strings.ToLower(filepath.Ext(fileName))
	if len(suffix) > 1 {
		suffix = suffix[1:] // Remove the leading dot
	}
	return whitelist.IsInList(suffix, whiteList)
}

func validateFileSize(size int64, maxSize float64) bool {
	fileSizeMB := float64(size) / 1024 / 1024
	return fileSizeMB <= maxSize
}

func isDuplicateFile(filePath string, uploadedFile io.ReadSeeker) bool {
	existingData, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}

	uploadedFile.Seek(0, io.SeekStart)
	uploadedData, _ := io.ReadAll(uploadedFile)

	oldFileHash := md5_util.MD5(existingData)
	newFileHash := md5_util.MD5(uploadedData)

	return oldFileHash == newFileHash
}

func saveFile(filePath string, file io.Reader) error {
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	return err
}
