package logic

import (
	"context"

	"im_server/common/list_query"
	"im_server/common/models"
	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"
	"im_server/im_user/user_models"
	"im_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserValidListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友验证列表
func NewUserValidListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserValidListLogic {
	return &UserValidListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserValidListLogic) UserValidList(req *types.FriendValidRequest, token string) (resp *types.FriendValidResponse, err error) {
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return
	}
	my_id := claims.UserID
	fvs, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.FriendVerifyModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Where:   l.svcCtx.DB.Where("send_user_id = ? or rev_user_id = ?", my_id, my_id),
		Preload: []string{"RevUserModel.UserConfModel", "SendUserModel.UserConfModel"},
	})
	var list []types.FriendValidInfo
	for _, fv := range fvs {
		info := types.FriendValidInfo{
			AdditionalMessages: fv.AdditionalMessages,
			ID:                 fv.ID,
		}
		if fv.SendUserID == my_id {
			// 我是发起方
			info.UserID = fv.RevUserID
			info.Nickname = fv.RevUserModel.Nickname
			info.Avatar = fv.RevUserModel.Avatar
			info.Verification = fv.RevUserModel.UserConfModel.Verification
			info.Status = fv.SendStatus
			info.Flag = "send"
		}
		if fv.RevUserID == my_id {
			// 我是接收方
			info.UserID = fv.SendUserID
			info.Nickname = fv.SendUserModel.Nickname
			info.Avatar = fv.SendUserModel.Avatar
			info.Verification = fv.SendUserModel.UserConfModel.Verification
			info.Status = fv.RevStatus
			info.Flag = "rev"
		}
		if fv.VerificationQuestion != nil {
			info.VerificationQuestion = &types.VerificationQuestion{
				Problem1: fv.VerificationQuestion.Problem1,
				Problem2: fv.VerificationQuestion.Problem2,
				Problem3: fv.VerificationQuestion.Problem3,
				Answer1:  fv.VerificationQuestion.Answer1,
				Answer2:  fv.VerificationQuestion.Answer2,
				Answer3:  fv.VerificationQuestion.Answer3,
			}
		}
		list = append(list, info)
	}
	return &types.FriendValidResponse{
		List:  list,
		Count: count,
	}, nil

}
