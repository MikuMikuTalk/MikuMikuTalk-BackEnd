package logic

import (
	"context"
	"errors"
	"strings"

	"im_server/im_file/file_api/file_models"
	"im_server/im_file/file_rpc/internal/svc"
	"im_server/im_file/file_rpc/types/file_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoLogic {
	return &FileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileInfoLogic) FileInfo(in *file_rpc.FileInfoRequest) (*file_rpc.FileInfoResponse, error) {
	var file_model file_models.FileModel
	err := l.svcCtx.DB.Take(&file_model, "uid = ?", in.FildId).Error
	if err != nil {
		return nil, errors.New("文件不存在")
	}
	var file_type string
	nameList := strings.Split(file_model.FileName, ".")
	if len(nameList) > 1 {
		file_type = nameList[len(nameList)-1]
	}
	return &file_rpc.FileInfoResponse{
		FileName: file_model.FileName,
		FileHash: file_model.Hash,
		FileSize: file_model.Size,
		FileType: file_type,
	}, nil
}
