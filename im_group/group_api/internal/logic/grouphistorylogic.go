package logic

import (
	"context"
	"errors"
	"time"

	"im_server/common/ctype"
	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"
	"im_server/im_group/group_models"
	"im_server/im_user/user_rpc/types/user_rpc"
	"im_server/utils/jwts"
	"im_server/utils/list_util"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupHistoryLogic {
	return &GroupHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HistoryResponse struct {
	UserID       uint      `json:"userID"`
	UserNickname string    `json:"userNickname"`
	UserAvatar   string    `json:"userAvatar"`
	Msg          ctype.Msg `json:"msg"`
	ID           uint      `json:"id"`
	MsgType      int8      `json:"msgType"`
	CreatedAt    time.Time `json:"createdAt"`
}

type HistoryListResponse struct {
	List  []HistoryResponse `json:"list"`
	Count int               `json:"count"`
}

func (l *GroupHistoryLogic) GroupHistory(req *types.GroupHistoryRequest) (resp *HistoryListResponse, err error) {
	// 解析 JWT 获取当前用户 ID
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, errors.New("无效的 Token")
	}
	myID := claims.UserID
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.ID, myID).Error
	if err != nil {
		return nil, errors.New("该用户不是群成员")
	}

	groupMsgList, count, err := list_query.ListQuery(l.svcCtx.DB, group_models.GroupMsgModel{GroupID: req.ID}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	var userIDList []uint32
	for _, model := range groupMsgList {
		userIDList = append(userIDList, uint32(model.SendUserID))
	}
	// 去重
	userIDList = list_util.DeduplicationList(userIDList)

	userListResponse, err1 := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	var list = make([]HistoryResponse, 0)
	for _, model := range groupMsgList {
		info := HistoryResponse{
			UserID:    model.SendUserID,
			Msg:       model.Msg,
			ID:        model.ID,
			MsgType:   model.MsgType,
			CreatedAt: model.CreatedAt,
		}
		if err1 == nil {
			info.UserNickname = userListResponse.UserInfo[uint32(info.UserID)].NickName
			info.UserAvatar = userListResponse.UserInfo[uint32(info.UserID)].Avatar
		}
		list = append(list, info)
	}

	resp = new(HistoryListResponse)
	resp.List = list
	resp.Count = int(count)
	return
}
