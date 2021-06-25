// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: user_message.proto

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

type Role int32

const (
	Role_UNKNOWN Role = 0
	Role_USER    Role = 1
	Role_AGENT   Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "UNKNOWN",
		1: "USER",
		2: "AGENT",
	}
	Role_value = map[string]int32{
		"UNKNOWN": 0,
		"USER":    1,
		"AGENT":   2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_user_message_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_user_message_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_user_message_proto_rawDescGZIP(), []int{0}
}

type UserMessage_Gender int32

const (
	UserMessage_UNKNOWN UserMessage_Gender = 0
	UserMessage_MALE    UserMessage_Gender = 1
	UserMessage_FEMALE  UserMessage_Gender = 2
)

// Enum value maps for UserMessage_Gender.
var (
	UserMessage_Gender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MALE",
		2: "FEMALE",
	}
	UserMessage_Gender_value = map[string]int32{
		"UNKNOWN": 0,
		"MALE":    1,
		"FEMALE":  2,
	}
)

func (x UserMessage_Gender) Enum() *UserMessage_Gender {
	p := new(UserMessage_Gender)
	*p = x
	return p
}

func (x UserMessage_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserMessage_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_user_message_proto_enumTypes[1].Descriptor()
}

func (UserMessage_Gender) Type() protoreflect.EnumType {
	return &file_user_message_proto_enumTypes[1]
}

func (x UserMessage_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserMessage_Gender.Descriptor instead.
func (UserMessage_Gender) EnumDescriptor() ([]byte, []int) {
	return file_user_message_proto_rawDescGZIP(), []int{0, 0}
}

type UserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string             `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password  string             `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name      string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Surname   string             `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	BirthDate int64              `protobuf:"varint,5,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Website   string             `protobuf:"bytes,6,opt,name=website,proto3" json:"website,omitempty"`
	Biography string             `protobuf:"bytes,7,opt,name=biography,proto3" json:"biography,omitempty"`
	Gender    UserMessage_Gender `protobuf:"varint,8,opt,name=gender,proto3,enum=proto.UserMessage_Gender" json:"gender,omitempty"`
	Public    bool               `protobuf:"varint,9,opt,name=public,proto3" json:"public,omitempty"`
	Taggable  bool               `protobuf:"varint,10,opt,name=taggable,proto3" json:"taggable,omitempty"`
	Active    bool               `protobuf:"varint,11,opt,name=active,proto3" json:"active,omitempty"`
	Email     string             `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	Role      Role               `protobuf:"varint,13,opt,name=role,proto3,enum=proto.Role" json:"role,omitempty"`
	Phone     string             `protobuf:"bytes,14,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserMessage) Reset() {
	*x = UserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessage) ProtoMessage() {}

func (x *UserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_user_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessage.ProtoReflect.Descriptor instead.
func (*UserMessage) Descriptor() ([]byte, []int) {
	return file_user_message_proto_rawDescGZIP(), []int{0}
}

func (x *UserMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserMessage) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UserMessage) GetBirthDate() int64 {
	if x != nil {
		return x.BirthDate
	}
	return 0
}

func (x *UserMessage) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *UserMessage) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *UserMessage) GetGender() UserMessage_Gender {
	if x != nil {
		return x.Gender
	}
	return UserMessage_UNKNOWN
}

func (x *UserMessage) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *UserMessage) GetTaggable() bool {
	if x != nil {
		return x.Taggable
	}
	return false
}

func (x *UserMessage) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *UserMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserMessage) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_UNKNOWN
}

func (x *UserMessage) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_user_message_proto protoreflect.FileDescriptor

var file_user_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x03, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x61, 0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10,
	0x02, 0x2a, 0x28, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x73, 0x74, 0x61, 0x67,
	0x72, 0x61, 0x6d, 0x2d, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_message_proto_rawDescOnce sync.Once
	file_user_message_proto_rawDescData = file_user_message_proto_rawDesc
)

func file_user_message_proto_rawDescGZIP() []byte {
	file_user_message_proto_rawDescOnce.Do(func() {
		file_user_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_message_proto_rawDescData)
	})
	return file_user_message_proto_rawDescData
}

var file_user_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_message_proto_goTypes = []interface{}{
	(Role)(0),               // 0: proto.Role
	(UserMessage_Gender)(0), // 1: proto.UserMessage.Gender
	(*UserMessage)(nil),     // 2: proto.UserMessage
}
var file_user_message_proto_depIdxs = []int32{
	1, // 0: proto.UserMessage.gender:type_name -> proto.UserMessage.Gender
	0, // 1: proto.UserMessage.role:type_name -> proto.Role
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_message_proto_init() }
func file_user_message_proto_init() {
	if File_user_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMessage); i {
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
			RawDescriptor: file_user_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_message_proto_goTypes,
		DependencyIndexes: file_user_message_proto_depIdxs,
		EnumInfos:         file_user_message_proto_enumTypes,
		MessageInfos:      file_user_message_proto_msgTypes,
	}.Build()
	File_user_message_proto = out.File
	file_user_message_proto_rawDesc = nil
	file_user_message_proto_goTypes = nil
	file_user_message_proto_depIdxs = nil
}
