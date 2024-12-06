package logic

import (
	"context"
	"errors"

	"im_server/common/ctype"
	"im_server/im_user/user_models"
	"im_server/utils/jwts"
	"im_server/utils/ref_map"

	"im_server/im_user/user_api/internal/svc"
	"im_server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户信息更新接口
func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(token string, req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {
	claims, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AuthSecret)
	if err != nil {
		logx.Error("error: ", err)
		return nil, err
	}
	user_id := claims.UserID
	userMaps := ref_map.RefToMap(*req, "user")
	userConfMaps := ref_map.RefToMap(*req, "user_conf")
	logx.Info("userMaps: ", userMaps)
	logx.Info("userConfMaps:", userConfMaps)
	if len(userMaps) != 0 {
		var user user_models.UserModel
		err = l.svcCtx.DB.Take(&user, user_id).Error
		if err != nil {
			return nil, errors.New("用户不存在")
		}
		if nick, ok := userMaps["nickname"].(string); ok && nick != user.Nickname {
			var existingUser user_models.UserModel
			err = l.svcCtx.DB.Where("nickname = ?", nick).First(&existingUser).Error
			if err == nil {
				return nil, errors.New("用户名已存在")
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				logx.Error("检查用户名是否存在时发生错误: ", err)
				return nil, errors.New("用户信息更新失败")
			}
		}

		if *(req.Nickname) == user.Nickname {
			err = errors.New("已经是原来的用户名了哦")
			return nil, err
		}
		err = l.svcCtx.DB.Model(&user).Updates(userMaps).Error
		if err != nil {

			return nil, errors.New("用户信息更新失败")
		}
	}
	if len(userConfMaps) != 0 {
		var userConf user_models.UserConfModel
		err = l.svcCtx.DB.Take(&userConf, "user_id = ?", user_id).Error
		if err != nil {
			return nil, errors.New("用户配置信息不存在")
		}
		verificationQuestion, ok := userConfMaps["verification_question"]
		if ok {
			delete(userConfMaps, "verification_question")
			data := ctype.VerificationQuestion{}
			if val, ok := verificationQuestion.(map[string]any)["problem1"]; ok {
				s := val.(string)
				data.Problem1 = &s
			}
			if val, ok := verificationQuestion.(map[string]any)["problem2"]; ok {
				s := val.(string)
				data.Problem2 = &s
			}
			if val, ok := verificationQuestion.(map[string]any)["problem3"]; ok {
				s := val.(string)
				data.Problem3 = &s
			}
			if val, ok := verificationQuestion.(map[string]any)["answer1"]; ok {
				s := val.(string)
				data.Answer1 = &s
			}
			if val, ok := verificationQuestion.(map[string]any)["answer2"]; ok {
				s := val.(string)
				data.Answer2 = &s
			}
			if val, ok := verificationQuestion.(map[string]any)["answer3"]; ok {
				s := val.(string)
				data.Answer3 = &s
			}
			err = l.svcCtx.DB.Model(&userConf).Updates(&user_models.UserConfModel{
				VerificationQuestion: &data,
			}).Error
			if err != nil {
				logx.Error("要更新的用户配置信息: ", data)
				logx.Error(err)
				return nil, errors.New("用户配置信息更新失败")
			}
		}
		err = l.svcCtx.DB.Model(&userConf).Updates(&userConfMaps).Error
		if err != nil {
			logx.Error("要更新的用户配置信息: ", userConfMaps)
			logx.Error(err)
			return nil, errors.New("用户配置信息更新失败")
		}
	}
	return
}
