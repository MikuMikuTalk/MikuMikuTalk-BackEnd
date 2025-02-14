// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: group_rpc.proto

package group_rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Groups_IsInGroup_FullMethodName       = "/group_rpc.Groups/IsInGroup"
	Groups_UserGroupSearch_FullMethodName = "/group_rpc.Groups/UserGroupSearch"
)

// GroupsClient is the client API for Groups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupsClient interface {
	IsInGroup(ctx context.Context, in *IsInGroupRequest, opts ...grpc.CallOption) (*IsInGroupResponse, error)
	UserGroupSearch(ctx context.Context, in *UserGroupSearchRequest, opts ...grpc.CallOption) (*UserGroupSearchResponse, error)
}

type groupsClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupsClient(cc grpc.ClientConnInterface) GroupsClient {
	return &groupsClient{cc}
}

func (c *groupsClient) IsInGroup(ctx context.Context, in *IsInGroupRequest, opts ...grpc.CallOption) (*IsInGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsInGroupResponse)
	err := c.cc.Invoke(ctx, Groups_IsInGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) UserGroupSearch(ctx context.Context, in *UserGroupSearchRequest, opts ...grpc.CallOption) (*UserGroupSearchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserGroupSearchResponse)
	err := c.cc.Invoke(ctx, Groups_UserGroupSearch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupsServer is the server API for Groups service.
// All implementations must embed UnimplementedGroupsServer
// for forward compatibility.
type GroupsServer interface {
	IsInGroup(context.Context, *IsInGroupRequest) (*IsInGroupResponse, error)
	UserGroupSearch(context.Context, *UserGroupSearchRequest) (*UserGroupSearchResponse, error)
	mustEmbedUnimplementedGroupsServer()
}

// UnimplementedGroupsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGroupsServer struct{}

func (UnimplementedGroupsServer) IsInGroup(context.Context, *IsInGroupRequest) (*IsInGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInGroup not implemented")
}
func (UnimplementedGroupsServer) UserGroupSearch(context.Context, *UserGroupSearchRequest) (*UserGroupSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGroupSearch not implemented")
}
func (UnimplementedGroupsServer) mustEmbedUnimplementedGroupsServer() {}
func (UnimplementedGroupsServer) testEmbeddedByValue()                {}

// UnsafeGroupsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupsServer will
// result in compilation errors.
type UnsafeGroupsServer interface {
	mustEmbedUnimplementedGroupsServer()
}

func RegisterGroupsServer(s grpc.ServiceRegistrar, srv GroupsServer) {
	// If the following call pancis, it indicates UnimplementedGroupsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Groups_ServiceDesc, srv)
}

func _Groups_IsInGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsInGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).IsInGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Groups_IsInGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).IsInGroup(ctx, req.(*IsInGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_UserGroupSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).UserGroupSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Groups_UserGroupSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).UserGroupSearch(ctx, req.(*UserGroupSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Groups_ServiceDesc is the grpc.ServiceDesc for Groups service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Groups_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "group_rpc.Groups",
	HandlerType: (*GroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsInGroup",
			Handler:    _Groups_IsInGroup_Handler,
		},
		{
			MethodName: "UserGroupSearch",
			Handler:    _Groups_UserGroupSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group_rpc.proto",
}
