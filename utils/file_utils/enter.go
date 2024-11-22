package file_utils

import (
	"path/filepath"
	"strings"
)

func GetFileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
func GetFileExtName(fileName string) string {
	return strings.TrimPrefix(filepath.Ext(fileName), ".")
}
