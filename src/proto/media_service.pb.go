// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: media_service.proto

package proto

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

type SaveMediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *MediaMessage `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *SaveMediaRequest) Reset() {
	*x = SaveMediaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMediaRequest) ProtoMessage() {}

func (x *SaveMediaRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SaveMediaRequest.ProtoReflect.Descriptor instead.
func (*SaveMediaRequest) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{0}
}

func (x *SaveMediaRequest) GetImage() *MediaMessage {
	if x != nil {
		return x.Image
	}
	return nil
}

type SaveMediaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveMediaResponse) Reset() {
	*x = SaveMediaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMediaResponse) ProtoMessage() {}

func (x *SaveMediaResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SaveMediaResponse.ProtoReflect.Descriptor instead.
func (*SaveMediaResponse) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{1}
}

func (x *SaveMediaResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMediaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMediaRequest) Reset() {
	*x = GetMediaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMediaRequest) ProtoMessage() {}

func (x *GetMediaRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetMediaRequest.ProtoReflect.Descriptor instead.
func (*GetMediaRequest) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMediaRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMediaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *MediaMessage `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetMediaResponse) Reset() {
	*x = GetMediaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMediaResponse) ProtoMessage() {}

func (x *GetMediaResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetMediaResponse.ProtoReflect.Descriptor instead.
func (*GetMediaResponse) Descriptor() ([]byte, []int) {
	return file_media_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetMediaResponse) GetImage() *MediaMessage {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_media_service_proto protoreflect.FileDescriptor

var file_media_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3d, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x23, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x32, 0x8b, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2d, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x69, 0x73, 0x74, 0x61,
	0x67, 0x72, 0x61, 0x6d, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*SaveMediaRequest)(nil),  // 0: proto.SaveMediaRequest
	(*SaveMediaResponse)(nil), // 1: proto.SaveMediaResponse
	(*GetMediaRequest)(nil),   // 2: proto.GetMediaRequest
	(*GetMediaResponse)(nil),  // 3: proto.GetMediaResponse
	(*MediaMessage)(nil),      // 4: proto.MediaMessage
}
var file_media_service_proto_depIdxs = []int32{
	4, // 0: proto.SaveMediaRequest.image:type_name -> proto.MediaMessage
	4, // 1: proto.GetMediaResponse.image:type_name -> proto.MediaMessage
	0, // 2: proto.MediaService.SaveMedia:input_type -> proto.SaveMediaRequest
	2, // 3: proto.MediaService.GetMedia:input_type -> proto.GetMediaRequest
	1, // 4: proto.MediaService.SaveMedia:output_type -> proto.SaveMediaResponse
	3, // 5: proto.MediaService.GetMedia:output_type -> proto.GetMediaResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_media_service_proto_init() }
func file_media_service_proto_init() {
	if File_media_service_proto != nil {
		return
	}
	file_media_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_media_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMediaRequest); i {
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
			switch v := v.(*SaveMediaResponse); i {
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
			switch v := v.(*GetMediaRequest); i {
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
			switch v := v.(*GetMediaResponse); i {
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
