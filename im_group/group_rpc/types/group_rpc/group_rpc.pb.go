// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v3.19.4
// source: group_rpc.proto

package group_rpc

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

type IsInGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId       uint32                 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsInGroupRequest) Reset() {
	*x = IsInGroupRequest{}
	mi := &file_group_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsInGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsInGroupRequest) ProtoMessage() {}

func (x *IsInGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsInGroupRequest.ProtoReflect.Descriptor instead.
func (*IsInGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *IsInGroupRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsInGroupRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type IsInGroupResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsInGroup     bool                   `protobuf:"varint,1,opt,name=is_in_group,json=isInGroup,proto3" json:"is_in_group,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsInGroupResponse) Reset() {
	*x = IsInGroupResponse{}
	mi := &file_group_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsInGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsInGroupResponse) ProtoMessage() {}

func (x *IsInGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsInGroupResponse.ProtoReflect.Descriptor instead.
func (*IsInGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *IsInGroupResponse) GetIsInGroup() bool {
	if x != nil {
		return x.IsInGroup
	}
	return false
}

var File_group_rpc_proto protoreflect.FileDescriptor

var file_group_rpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x46, 0x0a, 0x10,
	0x49, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x11, 0x49, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x32, 0x50, 0x0a, 0x06, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x46, 0x0a, 0x09, 0x49, 0x73, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x49,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x73, 0x49, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_group_rpc_proto_rawDescOnce sync.Once
	file_group_rpc_proto_rawDescData = file_group_rpc_proto_rawDesc
)

func file_group_rpc_proto_rawDescGZIP() []byte {
	file_group_rpc_proto_rawDescOnce.Do(func() {
		file_group_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_rpc_proto_rawDescData)
	})
	return file_group_rpc_proto_rawDescData
}

var file_group_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_group_rpc_proto_goTypes = []any{
	(*IsInGroupRequest)(nil),  // 0: group_rpc.IsInGroupRequest
	(*IsInGroupResponse)(nil), // 1: group_rpc.IsInGroupResponse
}
var file_group_rpc_proto_depIdxs = []int32{
	0, // 0: group_rpc.Groups.IsInGroup:input_type -> group_rpc.IsInGroupRequest
	1, // 1: group_rpc.Groups.IsInGroup:output_type -> group_rpc.IsInGroupResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_group_rpc_proto_init() }
func file_group_rpc_proto_init() {
	if File_group_rpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_group_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_rpc_proto_goTypes,
		DependencyIndexes: file_group_rpc_proto_depIdxs,
		MessageInfos:      file_group_rpc_proto_msgTypes,
	}.Build()
	File_group_rpc_proto = out.File
	file_group_rpc_proto_rawDesc = nil
	file_group_rpc_proto_goTypes = nil
	file_group_rpc_proto_depIdxs = nil
}
