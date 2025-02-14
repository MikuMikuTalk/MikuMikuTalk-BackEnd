package mqs

import (
	"context"
	"encoding/json"
	"im_server/common/contexts"
	"im_server/im_log/logs_api/internal/svc"
	"im_server/im_log/logs_model"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/addr"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogEvent struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *LogEvent {
	return &LogEvent{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Request struct {
	LogType int8   `json:"logType"`
	IP      string `json:"ip"`
	UserID  uint   `json:"userID"`
	Level   string `json:"level"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Service string `json:"service"`
}

func (l *LogEvent) Consume(ctx context.Context, key, val string) error {
	var req Request
	err := json.Unmarshal([]byte(val), &req)
	if err != nil {
		logx.Errorf("json解析错误 %s %s	", err.Error(), val)
		return err
	}

	// 查询ip对应的地址
	var info = logs_model.LogModel{
		LogType: req.LogType,
		IP:      req.IP,
		UserID:  req.UserID,
		Addr:    addr.GetAddr(req.IP),
		Level:   req.Level,
		Title:   req.Title,
		Content: req.Content,
		Service: req.Service,
	}
	l.ctx = context.WithValue(l.ctx, contexts.ContextKeyClientIP, info.IP)
	l.ctx = context.WithValue(l.ctx, contexts.ContextKeyUserID, info.UserID)
	if req.UserID != 0 {
		baseInfo, err := l.svcCtx.UserRpc.UserBaseInfo(l.ctx, &user_rpc.UserBaseInfoRequest{
			UserId: uint32(req.UserID),
		})
		if err != nil {
			logx.Error("logEventPayment Error on get user base info:", err)
			return err
		}
		info.UserNickname = baseInfo.NickName
		info.UserAvatar = baseInfo.Avatar
	}
	// 判断是不是运行日志
	if info.LogType == 3 {
		// 运行日志
		// 先查一下 今天这个服务有没有日志  有没有，有的话就更新，没有再创建
		mutex := sync.Mutex{}

		mutex.Lock()
		var logModel logs_model.LogModel
		err = l.svcCtx.DB.Take(&logModel, "log_type = ? and service = ? and to_days(created_at) = to_days(now())", 3, info.Service).Error
		mutex.Unlock()
		if err == nil {
			// 找到了
			l.svcCtx.DB.Model(&logModel).Update("content", logModel.Content+"\n"+info.Content)
			logx.Infof("运行日志 %s 更新成功", req.Title)
			return nil
		}
	}
	mutex := sync.Mutex{}

	mutex.Lock()
	err = l.svcCtx.DB.Create(&info).Error
	mutex.Unlock()
	if err != nil {
		logx.Error(err)
		return err
	}
	logx.Infof("日志 %s 保存成功", req.Title)
	// logx.Info(req)
	return nil
}
