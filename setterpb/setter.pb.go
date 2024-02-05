// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: setterpb/setter.proto

package setterpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetterFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Include bool `protobuf:"varint,1,opt,name=include,proto3" json:"include,omitempty"`
	Exclude bool `protobuf:"varint,2,opt,name=exclude,proto3" json:"exclude,omitempty"`
}

func (x *SetterFieldOptions) Reset() {
	*x = SetterFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setterpb_setter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetterFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetterFieldOptions) ProtoMessage() {}

func (x *SetterFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_setterpb_setter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetterFieldOptions.ProtoReflect.Descriptor instead.
func (*SetterFieldOptions) Descriptor() ([]byte, []int) {
	return file_setterpb_setter_proto_rawDescGZIP(), []int{0}
}

func (x *SetterFieldOptions) GetInclude() bool {
	if x != nil {
		return x.Include
	}
	return false
}

func (x *SetterFieldOptions) GetExclude() bool {
	if x != nil {
		return x.Exclude
	}
	return false
}

var file_setterpb_setter_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         94070,
		Name:          "setter.setter_messages",
		Tag:           "varint,94070,opt,name=setter_messages",
		Filename:      "setterpb/setter.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         94071,
		Name:          "setter.all_fields",
		Tag:           "varint,94071,opt,name=all_fields",
		Filename:      "setterpb/setter.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*SetterFieldOptions)(nil),
		Field:         94073,
		Name:          "setter.field",
		Tag:           "bytes,94073,opt,name=field",
		Filename:      "setterpb/setter.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*SetterFieldOptions)(nil),
		Field:         94074,
		Name:          "setter.oneof_field",
		Tag:           "bytes,94074,opt,name=oneof_field",
		Filename:      "setterpb/setter.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional bool setter_messages = 94070;
	E_SetterMessages = &file_setterpb_setter_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional bool all_fields = 94071;
	E_AllFields = &file_setterpb_setter_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional setter.SetterFieldOptions field = 94073;
	E_Field = &file_setterpb_setter_proto_extTypes[2]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional setter.SetterFieldOptions oneof_field = 94074;
	E_OneofField = &file_setterpb_setter_proto_extTypes[3]
)

var File_setterpb_setter_proto protoreflect.FileDescriptor

var file_setterpb_setter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x74, 0x74, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x48, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x3a, 0x47, 0x0a, 0x0f, 0x73,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf6, 0xde, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x3a, 0x40, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf7, 0xde, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x51, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf9,
	0xde, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x5c, 0x0a, 0x0b, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfa, 0xde, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x6f, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_setterpb_setter_proto_rawDescOnce sync.Once
	file_setterpb_setter_proto_rawDescData = file_setterpb_setter_proto_rawDesc
)

func file_setterpb_setter_proto_rawDescGZIP() []byte {
	file_setterpb_setter_proto_rawDescOnce.Do(func() {
		file_setterpb_setter_proto_rawDescData = protoimpl.X.CompressGZIP(file_setterpb_setter_proto_rawDescData)
	})
	return file_setterpb_setter_proto_rawDescData
}

var file_setterpb_setter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_setterpb_setter_proto_goTypes = []interface{}{
	(*SetterFieldOptions)(nil),          // 0: setter.SetterFieldOptions
	(*descriptorpb.FileOptions)(nil),    // 1: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil),   // 4: google.protobuf.OneofOptions
}
var file_setterpb_setter_proto_depIdxs = []int32{
	1, // 0: setter.setter_messages:extendee -> google.protobuf.FileOptions
	2, // 1: setter.all_fields:extendee -> google.protobuf.MessageOptions
	3, // 2: setter.field:extendee -> google.protobuf.FieldOptions
	4, // 3: setter.oneof_field:extendee -> google.protobuf.OneofOptions
	0, // 4: setter.field:type_name -> setter.SetterFieldOptions
	0, // 5: setter.oneof_field:type_name -> setter.SetterFieldOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	4, // [4:6] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_setterpb_setter_proto_init() }
func file_setterpb_setter_proto_init() {
	if File_setterpb_setter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_setterpb_setter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetterFieldOptions); i {
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
			RawDescriptor: file_setterpb_setter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_setterpb_setter_proto_goTypes,
		DependencyIndexes: file_setterpb_setter_proto_depIdxs,
		MessageInfos:      file_setterpb_setter_proto_msgTypes,
		ExtensionInfos:    file_setterpb_setter_proto_extTypes,
	}.Build()
	File_setterpb_setter_proto = out.File
	file_setterpb_setter_proto_rawDesc = nil
	file_setterpb_setter_proto_goTypes = nil
	file_setterpb_setter_proto_depIdxs = nil
}
