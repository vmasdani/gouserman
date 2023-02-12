// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: gouserman.proto

package main

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

type GousermanUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GormModel *GousermanGormModel `protobuf:"bytes,1,opt,name=gorm_model,json=gormModel,proto3,oneof" json:"gorm_model,omitempty"`
	BaseModel *GousermanBaseModel `protobuf:"bytes,2,opt,name=base_model,json=baseModel,proto3,oneof" json:"base_model,omitempty"`
	Name      *string             `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Username  *string             `protobuf:"bytes,4,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password  *string             `protobuf:"bytes,5,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Email     *string             `protobuf:"bytes,6,opt,name=email,proto3,oneof" json:"email,omitempty"`
}

func (x *GousermanUser) Reset() {
	*x = GousermanUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gouserman_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GousermanUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GousermanUser) ProtoMessage() {}

func (x *GousermanUser) ProtoReflect() protoreflect.Message {
	mi := &file_gouserman_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GousermanUser.ProtoReflect.Descriptor instead.
func (*GousermanUser) Descriptor() ([]byte, []int) {
	return file_gouserman_proto_rawDescGZIP(), []int{0}
}

func (x *GousermanUser) GetGormModel() *GousermanGormModel {
	if x != nil {
		return x.GormModel
	}
	return nil
}

func (x *GousermanUser) GetBaseModel() *GousermanBaseModel {
	if x != nil {
		return x.BaseModel
	}
	return nil
}

func (x *GousermanUser) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GousermanUser) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *GousermanUser) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *GousermanUser) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

type GousermanUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*GousermanUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GousermanUsers) Reset() {
	*x = GousermanUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gouserman_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GousermanUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GousermanUsers) ProtoMessage() {}

func (x *GousermanUsers) ProtoReflect() protoreflect.Message {
	mi := &file_gouserman_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GousermanUsers.ProtoReflect.Descriptor instead.
func (*GousermanUsers) Descriptor() ([]byte, []int) {
	return file_gouserman_proto_rawDescGZIP(), []int{1}
}

func (x *GousermanUsers) GetUsers() []*GousermanUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type GousermanGormModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *uint64 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CreatedAt *string `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt *string `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	DeletedAt *string `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
}

func (x *GousermanGormModel) Reset() {
	*x = GousermanGormModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gouserman_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GousermanGormModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GousermanGormModel) ProtoMessage() {}

func (x *GousermanGormModel) ProtoReflect() protoreflect.Message {
	mi := &file_gouserman_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GousermanGormModel.ProtoReflect.Descriptor instead.
func (*GousermanGormModel) Descriptor() ([]byte, []int) {
	return file_gouserman_proto_rawDescGZIP(), []int{2}
}

func (x *GousermanGormModel) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GousermanGormModel) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

func (x *GousermanGormModel) GetUpdatedAt() string {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return ""
}

func (x *GousermanGormModel) GetDeletedAt() string {
	if x != nil && x.DeletedAt != nil {
		return *x.DeletedAt
	}
	return ""
}

type GousermanBaseModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GousermanBaseModel) Reset() {
	*x = GousermanBaseModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gouserman_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GousermanBaseModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GousermanBaseModel) ProtoMessage() {}

func (x *GousermanBaseModel) ProtoReflect() protoreflect.Message {
	mi := &file_gouserman_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GousermanBaseModel.ProtoReflect.Descriptor instead.
func (*GousermanBaseModel) Descriptor() ([]byte, []int) {
	return file_gouserman_proto_rawDescGZIP(), []int{3}
}

var File_gouserman_proto protoreflect.FileDescriptor

var file_gouserman_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc2, 0x02, 0x0a, 0x0d, 0x47, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x6f, 0x75, 0x73, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x47, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x09,
	0x67, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x0a,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x47, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x42, 0x61, 0x73, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x48, 0x01, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x05, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x36, 0x0a, 0x0e, 0x47, 0x6f, 0x75, 0x73, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x47, 0x6f, 0x75, 0x73, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xc9,
	0x01, 0x0a, 0x12, 0x47, 0x6f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x47, 0x6f, 0x72, 0x6d,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x6f,
	0x75, 0x73, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gouserman_proto_rawDescOnce sync.Once
	file_gouserman_proto_rawDescData = file_gouserman_proto_rawDesc
)

func file_gouserman_proto_rawDescGZIP() []byte {
	file_gouserman_proto_rawDescOnce.Do(func() {
		file_gouserman_proto_rawDescData = protoimpl.X.CompressGZIP(file_gouserman_proto_rawDescData)
	})
	return file_gouserman_proto_rawDescData
}

var file_gouserman_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gouserman_proto_goTypes = []interface{}{
	(*GousermanUser)(nil),      // 0: GousermanUser
	(*GousermanUsers)(nil),     // 1: GousermanUsers
	(*GousermanGormModel)(nil), // 2: GousermanGormModel
	(*GousermanBaseModel)(nil), // 3: GousermanBaseModel
}
var file_gouserman_proto_depIdxs = []int32{
	2, // 0: GousermanUser.gorm_model:type_name -> GousermanGormModel
	3, // 1: GousermanUser.base_model:type_name -> GousermanBaseModel
	0, // 2: GousermanUsers.users:type_name -> GousermanUser
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gouserman_proto_init() }
func file_gouserman_proto_init() {
	if File_gouserman_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gouserman_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GousermanUser); i {
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
		file_gouserman_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GousermanUsers); i {
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
		file_gouserman_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GousermanGormModel); i {
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
		file_gouserman_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GousermanBaseModel); i {
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
	file_gouserman_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_gouserman_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gouserman_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gouserman_proto_goTypes,
		DependencyIndexes: file_gouserman_proto_depIdxs,
		MessageInfos:      file_gouserman_proto_msgTypes,
	}.Build()
	File_gouserman_proto = out.File
	file_gouserman_proto_rawDesc = nil
	file_gouserman_proto_goTypes = nil
	file_gouserman_proto_depIdxs = nil
}