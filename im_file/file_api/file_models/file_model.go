package file_models

import (
	"im_server/common/models"

	"github.com/google/uuid"
)

type FileModel struct {
	models.Models
	Uid      uuid.UUID `json:"uuid"`
	UserID   uint      `json:"userID"`
	FileName string    `gorm:"type:varchar(255);uniqueIndex:idx_filename_md5;not null;comment:文件名" json:"file_name"`
	Size     int64     `json:"size"`
	Path     string    `json:"path"`
	Hash     string    `gorm:"type:varchar(32);uniqueIndex:idx_filename_md5;not null;comment:文件md5值" json:"file_md5"`
}

func (file *FileModel) ImageWebPath() string {
	return "/api/file/" + file.Uid.String()
}
func (file *FileModel) FileWebPath() string {
	return "/api/file/download/" + file.Uid.String()
}
