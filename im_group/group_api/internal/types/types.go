// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5

package types

type GroupMemberInfo struct {
	UserID         uint   `json:"userId"`
	UserNickname   string `json:"userNickname"`
	Avatar         string `json:"avatar"`
	IsOnline       bool   `json:"isOnline"`
	Role           int8   `json:"role"`
	MemberNickname string `json:"memberNickname"`
	CreatedAt      string `json:"createdAt"`
	NewMsgDate     string `json:"newMsgDate"`
}

type GroupSearchResponse struct {
	GroupID         uint   `json:"groupId"`
	Title           string `json:"title"`
	Abstract        string `json:"abstract"`
	Avatar          string `json:"avatar"`
	IsInGroup       bool   `json:"isInGroup"`       // 我是否在群里面
	UserCount       int    `json:"userCount"`       // 群用户总数
	UserOnlineCount int    `json:"userOnlineCount"` // 群用户在线总数
}

type GroupfriendsResponse struct {
	UserId    uint   `json:"userId"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	IsInGroup bool   `json:"isInGroup"` // 是否在群里面
}

type UserInfo struct {
	UserID   uint   `json:"userId"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type VerificationQuestion struct {
	Problem1 *string `json:"problem1,optional" conf:"problem1"`
	Problem2 *string `json:"problem2,optional" conf:"problem2"`
	Problem3 *string `json:"problem3,optional" conf:"problem3"`
	Answer1  *string `json:"answer1,optional" conf:"answer1"`
	Answer2  *string `json:"answer2,optional" conf:"answer2"`
	Answer3  *string `json:"answer3,optional" conf:"answer3"`
}

type GroupCreateRequest struct {
	Token      string `header:"Authorization"`
	Mode       int8   `json:"mode,optional"` // 模式  1 直接创建模式 2 选人创建模式
	Name       string `json:"name,optional"`
	IsSearch   bool   `json:"isSearch,optional"`   // 是否可以搜到
	Size       int    `json:"size,optional"`       // 群规模
	UserIDList []uint `json:"userIdList,optional"` // 用户id列表
}

type GroupCreateResponse struct {
}

type GroupInfoRequest struct {
	Token string `header:"Authorization"`
	ID    uint   `path:"id"` //群id
}

type GroupInfoResponse struct {
	GroupID           uint       `json:"groupId"`
	Title             string     `json:"title"`
	Abstract          string     `json:"abstract"`
	MemberCount       int        `json:"memberCount"`
	MemberOnlineCount int        `json:"memberOnlineCount"`
	Avatar            string     `json:"avatar"`
	Creator           UserInfo   `json:"creator"`
	AdminList         []UserInfo `json:"adminList"`
	Role              int8       `json:"role"` // 角色  1 群主 2 群管理员 3 群成员
}

type GroupMemberAddRequest struct {
	Token        string `header:"Authorization"`
	ID           uint   `json:"id"`           // 群id
	MemberIDList []uint `json:"memberIdList"` // 成员id列表
}

type GroupMemberAddResponse struct {
}

type GroupMemberNicknameUpdateRequest struct {
	Token    string `header:"Authorization"`
	ID       uint   `json:"id"`       // 群id
	MemberID uint   `json:"memberId"` // 成员id
	Nickname string `json:"nickname"` // 昵称
}

type GroupMemberNicknameUpdateResponse struct {
}

type GroupMemberRemoveRequest struct {
	Token    string `header:"Authorization"`
	ID       uint   `form:"id"`       // 群id
	MemberID uint   `form:"memberId"` // 成员id
}

type GroupMemberRemoveResponse struct {
}

type GroupMemberRequest struct {
	Token string `header:"Authorization"`
	ID    uint   `form:"id"`
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
	Sort  string `form:"sort,optional"`
}

type GroupMemberResponse struct {
	List  []GroupMemberInfo `json:"list"`
	Count int               `json:"count"`
}

type GroupMemberRoleUpdateRequest struct {
	Token    string `header:"Authorization"`
	ID       uint   `json:"id"`       // 群id
	MemberID uint   `json:"memberId"` // 成员id
	Role     int8   `json:"role"`
}

type GroupMemberRoleUpdateResponse struct {
}

type GroupRemoveRequest struct {
	Token string `header:"Authorization"`
	ID    uint   `path:"id"` // 群id
}

type GroupRemoveResponse struct {
}

type GroupSearchListResponse struct {
	List  []GroupSearchResponse `json:"list"`
	Count int                   `json:"count"`
}

type GroupSearchRequest struct {
	Token string `header:"Authorization"`
	Key   string `form:"key,optional"` // 用户id和昵称
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
}

type GroupUpdateRequest struct {
	Token                string                `header:"Authorization"`
	ID                   uint                  `json:"id"`                                                       // 群id
	IsSearch             *bool                 `json:"isSearch,optional"  conf:"is_search"`                      // 是否可以被搜索
	Avatar               *string               `json:"avatar,optional"  conf:"avatar"`                           // 群头像
	Abstract             *string               `json:"abstract,optional"  conf:"abstract"`                       // 群简介
	Title                *string               `json:"title,optional"  conf:"title"`                             // 群名
	Verification         *int8                 `json:"verification,optional"  conf:"verification"`               // 群验证
	IsInvite             *bool                 `json:"isInvite,optional"  conf:"is_invite"`                      // 是否可邀请好友
	IsTemporarySession   *bool                 `json:"isTemporarySession,optional"  conf:"is_temporary_session"` // 是否开启临时会话
	IsProhibition        *bool                 `json:"isProhibition,optional" conf:"is_prohibition"`             // 是否开启全员禁言
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion,optional" conf:"verification_question"`
}

type GroupUpdateResponse struct {
}

type GroupfriendsListRequest struct {
	Token string `header:"Authorization"`
	ID    uint   `form:"id"` // 群id
}

type GroupfriendsListResponse struct {
	List  []GroupfriendsResponse `json:"list"`
	Count int                    `json:"count"`
}
