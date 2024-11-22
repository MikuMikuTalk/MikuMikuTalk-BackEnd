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
	"im_server/im_file/file_api/internal/logic"
	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"
	direcotry "im_server/utils/directory"
	"im_server/utils/file_utils"
	"im_server/utils/md5_util"
	"im_server/utils/str_util"
	"im_server/utils/whitelist"

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

		imageName := fileHeader.Filename
		fileExtName := file_utils.GetFileExtName(imageName)
		if !validateFileExtension(fileExtName, svcCtx.Config.WhiteList) {
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

		filePath := filepath.Join(dirName, imageName)
		mu.Lock() // Lock for concurrent access
		defer mu.Unlock()
		/*
			FIXME: 用户上传图片，文件名相同，但是md5值不同，然后重命名上传的新文件并进行存储。
					对于文件名相同，但是md5值也相同的，执行拒绝上传响应。
					我感觉这个逻辑有个问题，就是只要这个图片和第一张图片不是一张图(md5值不通)，但是名字相同，服务就会重新命名并存储，导致这个图片反而可以重复上传。
					这个问题怎么解决？用哈希表存hash的话，那不就上传图片太多，哈希表一直变大，存内存里面服务爆炸了么？但是持久化存储到数据库好像也不太好吧，怎么搞？
			TODO： 使用数据库存储文件名和hash值,上传的时候进行查询校验
		*/
		if isFileInDirectory(dirName, imageName) {
			// 检查文件内容是否重复
			if isDuplicateFile(filePath, file) {
				responseError(r, w, errors.New("不要上传重复图片"))
				return
			}

			// 如果文件内容不同，生成新文件名
			imageName = renameFile(imageName)
			filePath = filepath.Join(dirName, imageName)
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

func validateFileExtension(suffix string, whiteList []string) bool {
	return whitelist.IsInList(suffix, whiteList)
}

func validateFileSize(size int64, maxSize float64) bool {
	fileSizeMB := float64(size) / 1024 / 1024
	return fileSizeMB <= maxSize
}

func isFileInDirectory(dirName string, fileName string) bool {
	dirs, _ := os.ReadDir(dirName)
	return direcotry.InDir(dirs, fileName)
}

// 优化后的文件读取逻辑
func isDuplicateFile(filePath string, uploadedFile io.ReadSeeker) bool {
	// 重置文件流位置
	uploadedFile.Seek(0, io.SeekStart)
	defer uploadedFile.Seek(0, io.SeekStart) // 确保后续可用

	existingData, err := os.ReadFile(filePath)
	if err != nil {
		// 文件读取失败，认为不是重复文件
		return false
	}

	uploadedData, _ := io.ReadAll(uploadedFile)

	oldFileHash := md5_util.MD5(existingData)
	newFileHash := md5_util.MD5(uploadedData)
	logx.Info("oldFileHash: ", oldFileHash)
	logx.Info("newFileHash: ", newFileHash)
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

func renameFile(fileName string) string {
	fileNameWithoudExt := file_utils.GetFileNameWithoutExt(fileName)
	logx.Info("fileNameWithoudExt: ", fileNameWithoudExt)
	fileExtName := file_utils.GetFileExtName(fileName)
	random_str := str_util.GenerateRandomStr(8)
	newFileName := fmt.Sprintf("%s%s.%s", fileNameWithoudExt, random_str, fileExtName)
	return newFileName
}
