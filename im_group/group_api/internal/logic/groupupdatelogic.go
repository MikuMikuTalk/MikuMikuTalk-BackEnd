package logic

import (
	"context"
	"errors"

	"im_server/common/ctype"
	"im_server/im_group/group_models"
	"im_server/utils/jwts"
	"im_server/utils/ref_map"

	"im_server/im_group/group_api/internal/svc"
	"im_server/im_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUpdateLogic {
	return &GroupUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUpdateLogic) GroupUpdate(req *types.GroupUpdateRequest) (resp *types.GroupUpdateResponse, err error) {
	// 只能是群主或者管理员才可以调用群信息更新
	claims, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		return nil, err
	}
	my_id := claims.UserID

	var groupMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Preload("GroupModel").Take(&groupMember, "group_id = ? and user_id = ?", req.ID, my_id).Error
	if err != nil {
		return nil, errors.New("群不存在或者群用户不是群成员")
	}
	if !(groupMember.Role == 1 || groupMember.Role == 0) {
		return nil, errors.New("群信息只能由群主或者是管理员进行修改")
	}
	groupMaps := ref_map.RefToMap(*req, "conf")
	if len(groupMaps) != 0 {
		verificationQuestion, ok := groupMaps["verification_question"]
		if ok {
			delete(groupMaps, "verification_question")
			data := ctype.VerificationQuestion{}
			ref_map.MapToStruct(verificationQuestion.(map[string]any), &data)
			l.svcCtx.DB.Model(&groupMember.GroupModel).Updates(&group_models.GroupModel{
				VerificationQuestion: &data,
			})
		}
		err = l.svcCtx.DB.Model(&groupMember.GroupModel).Updates(groupMaps).Error
		if err != nil {
			logx.Error(err)
			return nil, errors.New("群信息更新失败")
		}
	}
	return
}
