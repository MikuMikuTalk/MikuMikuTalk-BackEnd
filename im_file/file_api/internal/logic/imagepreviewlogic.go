package logic

import (
	"context"

	"im_server/im_file/file_api/internal/svc"
	"im_server/im_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImagePreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 图片预览服务
func NewImagePreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImagePreviewLogic {
	return &ImagePreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImagePreviewLogic) ImagePreview(req *types.ImagePreviewRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
