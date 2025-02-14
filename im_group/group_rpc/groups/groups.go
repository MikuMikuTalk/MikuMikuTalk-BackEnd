// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: group_rpc.proto

package groups

import (
	"context"

	"im_server/im_group/group_rpc/types/group_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IsInGroupRequest        = group_rpc.IsInGroupRequest
	IsInGroupResponse       = group_rpc.IsInGroupResponse
	UserGroupSearchRequest  = group_rpc.UserGroupSearchRequest
	UserGroupSearchResponse = group_rpc.UserGroupSearchResponse

	Groups interface {
		IsInGroup(ctx context.Context, in *IsInGroupRequest, opts ...grpc.CallOption) (*IsInGroupResponse, error)
		UserGroupSearch(ctx context.Context, in *UserGroupSearchRequest, opts ...grpc.CallOption) (*UserGroupSearchResponse, error)
	}

	defaultGroups struct {
		cli zrpc.Client
	}
)

func NewGroups(cli zrpc.Client) Groups {
	return &defaultGroups{
		cli: cli,
	}
}

func (m *defaultGroups) IsInGroup(ctx context.Context, in *IsInGroupRequest, opts ...grpc.CallOption) (*IsInGroupResponse, error) {
	client := group_rpc.NewGroupsClient(m.cli.Conn())
	return client.IsInGroup(ctx, in, opts...)
}

func (m *defaultGroups) UserGroupSearch(ctx context.Context, in *UserGroupSearchRequest, opts ...grpc.CallOption) (*UserGroupSearchResponse, error) {
	client := group_rpc.NewGroupsClient(m.cli.Conn())
	return client.UserGroupSearch(ctx, in, opts...)
}
