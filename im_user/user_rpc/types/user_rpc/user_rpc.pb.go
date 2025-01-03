// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.4
// source: user_rpc/user_rpc.proto

package user_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserCreateRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	NickName       string                 `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Password       string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role           int32                  `protobuf:"varint,3,opt,name=role,proto3" json:"role,omitempty"`
	Avatar         string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	RegisterSource string                 `protobuf:"bytes,5,opt,name=register_source,json=registerSource,proto3" json:"register_source,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *UserCreateRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCreateRequest) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *UserCreateRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserCreateRequest) GetRegisterSource() string {
	if x != nil {
		return x.RegisterSource
	}
	return ""
}

type UserCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserCreateResponse) Reset() {
	*x = UserCreateResponse{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateResponse) ProtoMessage() {}

func (x *UserCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateResponse.ProtoReflect.Descriptor instead.
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreateResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// 用户信息请求
type UserInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfoRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfoResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NickName      string                 `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Avatar        string                 `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UserListInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserIdList    []uint32               `protobuf:"varint,1,rep,packed,name=user_id_list,json=userIdList,proto3" json:"user_id_list,omitempty"` // 用户id列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListInfoRequest) Reset() {
	*x = UserListInfoRequest{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListInfoRequest) ProtoMessage() {}

func (x *UserListInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListInfoRequest.ProtoReflect.Descriptor instead.
func (*UserListInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *UserListInfoRequest) GetUserIdList() []uint32 {
	if x != nil {
		return x.UserIdList
	}
	return nil
}

type UserListInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserInfo      map[uint32]*UserInfo   `protobuf:"bytes,1,rep,name=user_info,json=userInfo,proto3" json:"user_info,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // 用户信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListInfoResponse) Reset() {
	*x = UserListInfoResponse{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListInfoResponse) ProtoMessage() {}

func (x *UserListInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListInfoResponse.ProtoReflect.Descriptor instead.
func (*UserListInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *UserListInfoResponse) GetUserInfo() map[uint32]*UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type IsFriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User1         uint32                 `protobuf:"varint,1,opt,name=user1,proto3" json:"user1,omitempty"`
	User2         uint32                 `protobuf:"varint,2,opt,name=user2,proto3" json:"user2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsFriendRequest) Reset() {
	*x = IsFriendRequest{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFriendRequest) ProtoMessage() {}

func (x *IsFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFriendRequest.ProtoReflect.Descriptor instead.
func (*IsFriendRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *IsFriendRequest) GetUser1() uint32 {
	if x != nil {
		return x.User1
	}
	return 0
}

func (x *IsFriendRequest) GetUser2() uint32 {
	if x != nil {
		return x.User2
	}
	return 0
}

type IsFriendResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsFriend      bool                   `protobuf:"varint,1,opt,name=is_friend,json=isFriend,proto3" json:"is_friend,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsFriendResponse) Reset() {
	*x = IsFriendResponse{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFriendResponse) ProtoMessage() {}

func (x *IsFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFriendResponse.ProtoReflect.Descriptor instead.
func (*IsFriendResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *IsFriendResponse) GetIsFriend() bool {
	if x != nil {
		return x.IsFriend
	}
	return false
}

type FriendListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          uint32                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FriendListRequest) Reset() {
	*x = FriendListRequest{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendListRequest) ProtoMessage() {}

func (x *FriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendListRequest.ProtoReflect.Descriptor instead.
func (*FriendListRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{9}
}

func (x *FriendListRequest) GetUser() uint32 {
	if x != nil {
		return x.User
	}
	return 0
}

type FriendInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NickName      string                 `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Avatar        string                 `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FriendInfo) Reset() {
	*x = FriendInfo{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendInfo) ProtoMessage() {}

func (x *FriendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendInfo.ProtoReflect.Descriptor instead.
func (*FriendInfo) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{10}
}

func (x *FriendInfo) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FriendInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *FriendInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type FriendListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FriendList    []*FriendInfo          `protobuf:"bytes,1,rep,name=friend_list,json=friendList,proto3" json:"friend_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FriendListResponse) Reset() {
	*x = FriendListResponse{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendListResponse) ProtoMessage() {}

func (x *FriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendListResponse.ProtoReflect.Descriptor instead.
func (*FriendListResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{11}
}

func (x *FriendListResponse) GetFriendList() []*FriendInfo {
	if x != nil {
		return x.FriendList
	}
	return nil
}

type UserBaseInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserBaseInfoRequest) Reset() {
	*x = UserBaseInfoRequest{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserBaseInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBaseInfoRequest) ProtoMessage() {}

func (x *UserBaseInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBaseInfoRequest.ProtoReflect.Descriptor instead.
func (*UserBaseInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{12}
}

func (x *UserBaseInfoRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserBaseInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NickName      string                 `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Avatar        string                 `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserBaseInfoResponse) Reset() {
	*x = UserBaseInfoResponse{}
	mi := &file_user_rpc_user_rpc_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserBaseInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBaseInfoResponse) ProtoMessage() {}

func (x *UserBaseInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_rpc_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBaseInfoResponse.ProtoReflect.Descriptor instead.
func (*UserBaseInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_rpc_proto_rawDescGZIP(), []int{13}
}

func (x *UserBaseInfoResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserBaseInfoResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserBaseInfoResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_user_rpc_user_rpc_proto protoreflect.FileDescriptor

var file_user_rpc_user_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x31, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3f,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69,
	0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22,
	0x37, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x14, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x4f, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a,
	0x0f, 0x49, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32, 0x22, 0x2f, 0x0a, 0x10,
	0x49, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x22, 0x27, 0x0a,
	0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x22, 0x4b, 0x0a, 0x12, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x2e, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x64, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x32, 0xbd, 0x03, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x47, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x49, 0x73, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x49, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_rpc_user_rpc_proto_rawDescOnce sync.Once
	file_user_rpc_user_rpc_proto_rawDescData = file_user_rpc_user_rpc_proto_rawDesc
)

func file_user_rpc_user_rpc_proto_rawDescGZIP() []byte {
	file_user_rpc_user_rpc_proto_rawDescOnce.Do(func() {
		file_user_rpc_user_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_rpc_user_rpc_proto_rawDescData)
	})
	return file_user_rpc_user_rpc_proto_rawDescData
}

var file_user_rpc_user_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_user_rpc_user_rpc_proto_goTypes = []any{
	(*UserCreateRequest)(nil),    // 0: user_rpc.UserCreateRequest
	(*UserCreateResponse)(nil),   // 1: user_rpc.UserCreateResponse
	(*UserInfoRequest)(nil),      // 2: user_rpc.UserInfoRequest
	(*UserInfoResponse)(nil),     // 3: user_rpc.UserInfoResponse
	(*UserInfo)(nil),             // 4: user_rpc.UserInfo
	(*UserListInfoRequest)(nil),  // 5: user_rpc.UserListInfoRequest
	(*UserListInfoResponse)(nil), // 6: user_rpc.UserListInfoResponse
	(*IsFriendRequest)(nil),      // 7: user_rpc.IsFriendRequest
	(*IsFriendResponse)(nil),     // 8: user_rpc.IsFriendResponse
	(*FriendListRequest)(nil),    // 9: user_rpc.FriendListRequest
	(*FriendInfo)(nil),           // 10: user_rpc.FriendInfo
	(*FriendListResponse)(nil),   // 11: user_rpc.FriendListResponse
	(*UserBaseInfoRequest)(nil),  // 12: user_rpc.UserBaseInfoRequest
	(*UserBaseInfoResponse)(nil), // 13: user_rpc.UserBaseInfoResponse
	nil,                          // 14: user_rpc.UserListInfoResponse.UserInfoEntry
}
var file_user_rpc_user_rpc_proto_depIdxs = []int32{
	14, // 0: user_rpc.UserListInfoResponse.user_info:type_name -> user_rpc.UserListInfoResponse.UserInfoEntry
	10, // 1: user_rpc.FriendListResponse.friend_list:type_name -> user_rpc.FriendInfo
	4,  // 2: user_rpc.UserListInfoResponse.UserInfoEntry.value:type_name -> user_rpc.UserInfo
	0,  // 3: user_rpc.Users.UserCreate:input_type -> user_rpc.UserCreateRequest
	2,  // 4: user_rpc.Users.UserInfo:input_type -> user_rpc.UserInfoRequest
	12, // 5: user_rpc.Users.UserBaseInfo:input_type -> user_rpc.UserBaseInfoRequest
	5,  // 6: user_rpc.Users.UserListInfo:input_type -> user_rpc.UserListInfoRequest
	7,  // 7: user_rpc.Users.IsFriend:input_type -> user_rpc.IsFriendRequest
	9,  // 8: user_rpc.Users.FriendList:input_type -> user_rpc.FriendListRequest
	1,  // 9: user_rpc.Users.UserCreate:output_type -> user_rpc.UserCreateResponse
	3,  // 10: user_rpc.Users.UserInfo:output_type -> user_rpc.UserInfoResponse
	13, // 11: user_rpc.Users.UserBaseInfo:output_type -> user_rpc.UserBaseInfoResponse
	6,  // 12: user_rpc.Users.UserListInfo:output_type -> user_rpc.UserListInfoResponse
	8,  // 13: user_rpc.Users.IsFriend:output_type -> user_rpc.IsFriendResponse
	11, // 14: user_rpc.Users.FriendList:output_type -> user_rpc.FriendListResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_user_rpc_user_rpc_proto_init() }
func file_user_rpc_user_rpc_proto_init() {
	if File_user_rpc_user_rpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_rpc_user_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_rpc_user_rpc_proto_goTypes,
		DependencyIndexes: file_user_rpc_user_rpc_proto_depIdxs,
		MessageInfos:      file_user_rpc_user_rpc_proto_msgTypes,
	}.Build()
	File_user_rpc_user_rpc_proto = out.File
	file_user_rpc_user_rpc_proto_rawDesc = nil
	file_user_rpc_user_rpc_proto_goTypes = nil
	file_user_rpc_user_rpc_proto_depIdxs = nil
}
