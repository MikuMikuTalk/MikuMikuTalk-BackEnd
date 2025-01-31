package mqs

import (
	"context"
	"im_server/im_log/logs_api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Updated Consume method to match the kq.ConsumeHandler interface
func (l *PaymentSuccess) Consume(ctx context.Context, key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)
	return nil
}
