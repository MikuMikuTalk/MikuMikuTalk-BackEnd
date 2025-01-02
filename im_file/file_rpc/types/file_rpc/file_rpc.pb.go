// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.4
// source: file_rpc.proto

package file_rpc

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

type FileInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FildId        string                 `protobuf:"bytes,1,opt,name=fild_id,json=fildId,proto3" json:"fild_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileInfoRequest) Reset() {
	*x = FileInfoRequest{}
	mi := &file_file_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoRequest) ProtoMessage() {}

func (x *FileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoRequest.ProtoReflect.Descriptor instead.
func (*FileInfoRequest) Descriptor() ([]byte, []int) {
	return file_file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfoRequest) GetFildId() string {
	if x != nil {
		return x.FildId
	}
	return ""
}

type FileInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath      string                 `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileSize      int64                  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	FileType      string                 `protobuf:"bytes,4,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileInfoResponse) Reset() {
	*x = FileInfoResponse{}
	mi := &file_file_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoResponse) ProtoMessage() {}

func (x *FileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoResponse.ProtoReflect.Descriptor instead.
func (*FileInfoResponse) Descriptor() ([]byte, []int) {
	return file_file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfoResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfoResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileInfoResponse) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileInfoResponse) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

var File_file_rpc_proto protoreflect.FileDescriptor

var file_file_rpc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x2a, 0x0a, 0x0f, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32,
	0x4a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_file_rpc_proto_rawDescOnce sync.Once
	file_file_rpc_proto_rawDescData = file_file_rpc_proto_rawDesc
)

func file_file_rpc_proto_rawDescGZIP() []byte {
	file_file_rpc_proto_rawDescOnce.Do(func() {
		file_file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_rpc_proto_rawDescData)
	})
	return file_file_rpc_proto_rawDescData
}

var file_file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_file_rpc_proto_goTypes = []any{
	(*FileInfoRequest)(nil),  // 0: file_rpc.FileInfoRequest
	(*FileInfoResponse)(nil), // 1: file_rpc.FileInfoResponse
}
var file_file_rpc_proto_depIdxs = []int32{
	0, // 0: file_rpc.files.FileInfo:input_type -> file_rpc.FileInfoRequest
	1, // 1: file_rpc.files.FileInfo:output_type -> file_rpc.FileInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_rpc_proto_init() }
func file_file_rpc_proto_init() {
	if File_file_rpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_rpc_proto_goTypes,
		DependencyIndexes: file_file_rpc_proto_depIdxs,
		MessageInfos:      file_file_rpc_proto_msgTypes,
	}.Build()
	File_file_rpc_proto = out.File
	file_file_rpc_proto_rawDesc = nil
	file_file_rpc_proto_goTypes = nil
	file_file_rpc_proto_depIdxs = nil
}
