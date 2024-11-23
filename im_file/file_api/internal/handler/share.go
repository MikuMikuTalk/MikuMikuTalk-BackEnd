package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"im_server/common/response"
	direcotry "im_server/utils/directory"
	"im_server/utils/file_utils"
	"im_server/utils/md5_util"
	"im_server/utils/str_util"
	"im_server/utils/whitelist"

	"github.com/zeromicro/go-zero/core/logx"
)

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
