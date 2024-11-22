package file_models

import (
	"im_server/common/models"
)

type FileModel struct {
	models.Models
	FileName string `gorm:"type:varchar(255);uniqueIndex:idx_filename_md5;not null;comment:文件名" json:"file_name"`
	Hash     string `gorm:"type:varchar(32);uniqueIndex:idx_filename_md5;not null;comment:文件md5值" json:"file_md5"`
}
