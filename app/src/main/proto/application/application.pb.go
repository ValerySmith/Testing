// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: application/application.proto

package application

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

type ApplicationInfo_ApplicationType int32

const (
	ApplicationInfo_AT_Unknown          ApplicationInfo_ApplicationType = 0
	ApplicationInfo_AT_TRINITY_Player   ApplicationInfo_ApplicationType = 1
	ApplicationInfo_AT_SWEET_TV_Player  ApplicationInfo_ApplicationType = 2
	ApplicationInfo_AT_CACTUS_Player    ApplicationInfo_ApplicationType = 3
	ApplicationInfo_AT_Alpha_IP         ApplicationInfo_ApplicationType = 4
	ApplicationInfo_AT_GRIZLI_Player    ApplicationInfo_ApplicationType = 5
	ApplicationInfo_AT_Bravis_TV_Player ApplicationInfo_ApplicationType = 6
	ApplicationInfo_AT_Startelecom_TV   ApplicationInfo_ApplicationType = 7
	ApplicationInfo_AT_Apelsin          ApplicationInfo_ApplicationType = 8
	ApplicationInfo_AT_IMPERIAL_TV      ApplicationInfo_ApplicationType = 9
	ApplicationInfo_AT_AIWA             ApplicationInfo_ApplicationType = 10
	ApplicationInfo_AT_Kahovka          ApplicationInfo_ApplicationType = 11
	ApplicationInfo_AT_VODAFONE_TV      ApplicationInfo_ApplicationType = 12
	ApplicationInfo_AT_DiaNet           ApplicationInfo_ApplicationType = 13
	ApplicationInfo_AT_Kopeika          ApplicationInfo_ApplicationType = 14
	ApplicationInfo_AT_RNet             ApplicationInfo_ApplicationType = 15
	ApplicationInfo_AT_HISENSE          ApplicationInfo_ApplicationType = 16
	ApplicationInfo_AT_UA_CITY          ApplicationInfo_ApplicationType = 17
	ApplicationInfo_AT_Enlider          ApplicationInfo_ApplicationType = 18
	ApplicationInfo_AT_Bestlink         ApplicationInfo_ApplicationType = 19
	ApplicationInfo_AT_Streamnetwork    ApplicationInfo_ApplicationType = 20
	ApplicationInfo_AT_Homenet          ApplicationInfo_ApplicationType = 21
	ApplicationInfo_AT_NetZahid         ApplicationInfo_ApplicationType = 22
)

// Enum value maps for ApplicationInfo_ApplicationType.
var (
	ApplicationInfo_ApplicationType_name = map[int32]string{
		0:  "AT_Unknown",
		1:  "AT_TRINITY_Player",
		2:  "AT_SWEET_TV_Player",
		3:  "AT_CACTUS_Player",
		4:  "AT_Alpha_IP",
		5:  "AT_GRIZLI_Player",
		6:  "AT_Bravis_TV_Player",
		7:  "AT_Startelecom_TV",
		8:  "AT_Apelsin",
		9:  "AT_IMPERIAL_TV",
		10: "AT_AIWA",
		11: "AT_Kahovka",
		12: "AT_VODAFONE_TV",
		13: "AT_DiaNet",
		14: "AT_Kopeika",
		15: "AT_RNet",
		16: "AT_HISENSE",
		17: "AT_UA_CITY",
		18: "AT_Enlider",
		19: "AT_Bestlink",
		20: "AT_Streamnetwork",
		21: "AT_Homenet",
		22: "AT_NetZahid",
	}
	ApplicationInfo_ApplicationType_value = map[string]int32{
		"AT_Unknown":          0,
		"AT_TRINITY_Player":   1,
		"AT_SWEET_TV_Player":  2,
		"AT_CACTUS_Player":    3,
		"AT_Alpha_IP":         4,
		"AT_GRIZLI_Player":    5,
		"AT_Bravis_TV_Player": 6,
		"AT_Startelecom_TV":   7,
		"AT_Apelsin":          8,
		"AT_IMPERIAL_TV":      9,
		"AT_AIWA":             10,
		"AT_Kahovka":          11,
		"AT_VODAFONE_TV":      12,
		"AT_DiaNet":           13,
		"AT_Kopeika":          14,
		"AT_RNet":             15,
		"AT_HISENSE":          16,
		"AT_UA_CITY":          17,
		"AT_Enlider":          18,
		"AT_Bestlink":         19,
		"AT_Streamnetwork":    20,
		"AT_Homenet":          21,
		"AT_NetZahid":         22,
	}
)

func (x ApplicationInfo_ApplicationType) Enum() *ApplicationInfo_ApplicationType {
	p := new(ApplicationInfo_ApplicationType)
	*p = x
	return p
}

func (x ApplicationInfo_ApplicationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplicationInfo_ApplicationType) Descriptor() protoreflect.EnumDescriptor {
	return file_application_application_proto_enumTypes[0].Descriptor()
}

func (ApplicationInfo_ApplicationType) Type() protoreflect.EnumType {
	return &file_application_application_proto_enumTypes[0]
}

func (x ApplicationInfo_ApplicationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplicationInfo_ApplicationType.Descriptor instead.
func (ApplicationInfo_ApplicationType) EnumDescriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{0, 0}
}

type ApplicationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ApplicationInfo_ApplicationType `protobuf:"varint,1,opt,name=type,proto3,enum=application.ApplicationInfo_ApplicationType" json:"type,omitempty"`
}

func (x *ApplicationInfo) Reset() {
	*x = ApplicationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationInfo) ProtoMessage() {}

func (x *ApplicationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationInfo.ProtoReflect.Descriptor instead.
func (*ApplicationInfo) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationInfo) GetType() ApplicationInfo_ApplicationType {
	if x != nil {
		return x.Type
	}
	return ApplicationInfo_AT_Unknown
}

var File_application_application_proto protoreflect.FileDescriptor

var file_application_application_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x04, 0x0a,
	0x0f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0xb6, 0x03, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x54, 0x5f, 0x54, 0x52, 0x49,
	0x4e, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x41, 0x54, 0x5f, 0x53, 0x57, 0x45, 0x45, 0x54, 0x5f, 0x54, 0x56, 0x5f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x54, 0x5f, 0x43, 0x41, 0x43, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x54, 0x5f, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x49, 0x50, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x54, 0x5f, 0x47, 0x52, 0x49, 0x5a, 0x4c, 0x49, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x54, 0x5f, 0x42, 0x72, 0x61, 0x76, 0x69, 0x73, 0x5f,
	0x54, 0x56, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x41,
	0x54, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x6c, 0x65, 0x63, 0x6f, 0x6d, 0x5f, 0x54, 0x56,
	0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x41, 0x70, 0x65, 0x6c, 0x73, 0x69, 0x6e,
	0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x45, 0x52, 0x49, 0x41,
	0x4c, 0x5f, 0x54, 0x56, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x54, 0x5f, 0x41, 0x49, 0x57,
	0x41, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x4b, 0x61, 0x68, 0x6f, 0x76, 0x6b,
	0x61, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x54, 0x5f, 0x56, 0x4f, 0x44, 0x41, 0x46, 0x4f,
	0x4e, 0x45, 0x5f, 0x54, 0x56, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x54, 0x5f, 0x44, 0x69,
	0x61, 0x4e, 0x65, 0x74, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x4b, 0x6f, 0x70,
	0x65, 0x69, 0x6b, 0x61, 0x10, 0x0e, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x54, 0x5f, 0x52, 0x4e, 0x65,
	0x74, 0x10, 0x0f, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x48, 0x49, 0x53, 0x45, 0x4e, 0x53,
	0x45, 0x10, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x55, 0x41, 0x5f, 0x43, 0x49, 0x54,
	0x59, 0x10, 0x11, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x5f, 0x45, 0x6e, 0x6c, 0x69, 0x64, 0x65,
	0x72, 0x10, 0x12, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x54, 0x5f, 0x42, 0x65, 0x73, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x10, 0x13, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x54, 0x5f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x14, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54,
	0x5f, 0x48, 0x6f, 0x6d, 0x65, 0x6e, 0x65, 0x74, 0x10, 0x15, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x54,
	0x5f, 0x4e, 0x65, 0x74, 0x5a, 0x61, 0x68, 0x69, 0x64, 0x10, 0x16, 0x42, 0x45, 0x0a, 0x20, 0x63,
	0x6f, 0x6d, 0x2e, 0x75, 0x61, 0x2e, 0x6d, 0x79, 0x74, 0x72, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x74, 0x76, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x74, 0x76,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_application_proto_rawDescOnce sync.Once
	file_application_application_proto_rawDescData = file_application_application_proto_rawDesc
)

func file_application_application_proto_rawDescGZIP() []byte {
	file_application_application_proto_rawDescOnce.Do(func() {
		file_application_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_application_proto_rawDescData)
	})
	return file_application_application_proto_rawDescData
}

var file_application_application_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_application_application_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_application_application_proto_goTypes = []interface{}{
	(ApplicationInfo_ApplicationType)(0), // 0: application.ApplicationInfo.ApplicationType
	(*ApplicationInfo)(nil),              // 1: application.ApplicationInfo
}
var file_application_application_proto_depIdxs = []int32{
	0, // 0: application.ApplicationInfo.type:type_name -> application.ApplicationInfo.ApplicationType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_application_application_proto_init() }
func file_application_application_proto_init() {
	if File_application_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_application_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationInfo); i {
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
			RawDescriptor: file_application_application_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_application_application_proto_goTypes,
		DependencyIndexes: file_application_application_proto_depIdxs,
		EnumInfos:         file_application_application_proto_enumTypes,
		MessageInfos:      file_application_application_proto_msgTypes,
	}.Build()
	File_application_application_proto = out.File
	file_application_application_proto_rawDesc = nil
	file_application_application_proto_goTypes = nil
	file_application_application_proto_depIdxs = nil
}
