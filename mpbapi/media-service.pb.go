// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.14.0
// source: media-service.proto

package mpbapi

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

type AddPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddPostRequest) Reset() {
	*x = AddPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostRequest) ProtoMessage() {}

func (x *AddPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostRequest.ProtoReflect.Descriptor instead.
func (*AddPostRequest) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddPostRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
}

func (x *AddPostResponse) Reset() {
	*x = AddPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostResponse) ProtoMessage() {}

func (x *AddPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostResponse.ProtoReflect.Descriptor instead.
func (*AddPostResponse) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddPostResponse) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostRequest) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_media_service_proto protoreflect.FileDescriptor

var file_media_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x70, 0x62, 0x61, 0x70, 0x69, 0x22, 0x24, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64,
	0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x8a, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x6d, 0x70, 0x62, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x70,
	0x62, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x6d, 0x70, 0x62, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x70, 0x62, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x70, 0x62, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_media_service_proto_rawDescOnce sync.Once
	file_media_service_proto_rawDescData = file_media_service_proto_rawDesc
)

func file_media_service_proto_rawDescGZIP() []byte {
	file_media_service_proto_rawDescOnce.Do(func() {
		file_media_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_media_service_proto_rawDescData)
	})
	return file_media_service_proto_rawDescData
}

var file_media_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_media_service_proto_goTypes = []interface{}{
	(*AddPostRequest)(nil),  // 0: mpbapi.AddPostRequest
	(*AddPostResponse)(nil), // 1: mpbapi.AddPostResponse
	(*GetPostRequest)(nil),  // 2: mpbapi.GetPostRequest
	(*GetPostResponse)(nil), // 3: mpbapi.GetPostResponse
}
var file_media_service_proto_depIdxs = []int32{
	0, // 0: mpbapi.MediaService.AddPost:input_type -> mpbapi.AddPostRequest
	2, // 1: mpbapi.MediaService.GetPost:input_type -> mpbapi.GetPostRequest
	1, // 2: mpbapi.MediaService.AddPost:output_type -> mpbapi.AddPostResponse
	3, // 3: mpbapi.MediaService.GetPost:output_type -> mpbapi.GetPostResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_media_service_proto_init() }
func file_media_service_proto_init() {
	if File_media_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_media_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_media_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_media_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_media_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_media_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_media_service_proto_goTypes,
		DependencyIndexes: file_media_service_proto_depIdxs,
		MessageInfos:      file_media_service_proto_msgTypes,
	}.Build()
	File_media_service_proto = out.File
	file_media_service_proto_rawDesc = nil
	file_media_service_proto_goTypes = nil
	file_media_service_proto_depIdxs = nil
}
