// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user_rpc.proto

package users

import (
	"context"

	"im_server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FriendInfo           = user_rpc.FriendInfo
	FriendListRequest    = user_rpc.FriendListRequest
	FriendListResponse   = user_rpc.FriendListResponse
	IsFriendRequest      = user_rpc.IsFriendRequest
	IsFriendResponse     = user_rpc.IsFriendResponse
	UserBaseInfoRequest  = user_rpc.UserBaseInfoRequest
	UserBaseInfoResponse = user_rpc.UserBaseInfoResponse
	UserCreateRequest    = user_rpc.UserCreateRequest
	UserCreateResponse   = user_rpc.UserCreateResponse
	UserInfo             = user_rpc.UserInfo
	UserInfoRequest      = user_rpc.UserInfoRequest
	UserInfoResponse     = user_rpc.UserInfoResponse
	UserListInfoRequest  = user_rpc.UserListInfoRequest
	UserListInfoResponse = user_rpc.UserListInfoResponse

	Users interface {
		UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		UserBaseInfo(ctx context.Context, in *UserBaseInfoRequest, opts ...grpc.CallOption) (*UserBaseInfoResponse, error)
		UserListInfo(ctx context.Context, in *UserListInfoRequest, opts ...grpc.CallOption) (*UserListInfoResponse, error)
		IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*IsFriendResponse, error)
		FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	}

	defaultUsers struct {
		cli zrpc.Client
	}
)

func NewUsers(cli zrpc.Client) Users {
	return &defaultUsers{
		cli: cli,
	}
}

func (m *defaultUsers) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.UserCreate(ctx, in, opts...)
}

func (m *defaultUsers) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUsers) UserBaseInfo(ctx context.Context, in *UserBaseInfoRequest, opts ...grpc.CallOption) (*UserBaseInfoResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.UserBaseInfo(ctx, in, opts...)
}

func (m *defaultUsers) UserListInfo(ctx context.Context, in *UserListInfoRequest, opts ...grpc.CallOption) (*UserListInfoResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.UserListInfo(ctx, in, opts...)
}

func (m *defaultUsers) IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*IsFriendResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.IsFriend(ctx, in, opts...)
}

func (m *defaultUsers) FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}
