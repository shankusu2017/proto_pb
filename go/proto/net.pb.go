// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: net.proto

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

type Net struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubId int32 `protobuf:"varint,1,opt,name=SubId,proto3" json:"SubId,omitempty"` // 子网号
}

func (x *Net) Reset() {
	*x = Net{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Net) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Net) ProtoMessage() {}

func (x *Net) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Net.ProtoReflect.Descriptor instead.
func (*Net) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{0}
}

func (x *Net) GetSubId() int32 {
	if x != nil {
		return x.SubId
	}
	return 0
}

var File_net_proto protoreflect.FileDescriptor

var file_net_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x74,
	0x22, 0x1b, 0x0a, 0x03, 0x4e, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x75, 0x62, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x75, 0x62, 0x49, 0x64, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_net_proto_rawDescOnce sync.Once
	file_net_proto_rawDescData = file_net_proto_rawDesc
)

func file_net_proto_rawDescGZIP() []byte {
	file_net_proto_rawDescOnce.Do(func() {
		file_net_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_proto_rawDescData)
	})
	return file_net_proto_rawDescData
}

var file_net_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_net_proto_goTypes = []interface{}{
	(*Net)(nil), // 0: net.Net
}
var file_net_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_net_proto_init() }
func file_net_proto_init() {
	if File_net_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_net_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Net); i {
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
			RawDescriptor: file_net_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_net_proto_goTypes,
		DependencyIndexes: file_net_proto_depIdxs,
		MessageInfos:      file_net_proto_msgTypes,
	}.Build()
	File_net_proto = out.File
	file_net_proto_rawDesc = nil
	file_net_proto_goTypes = nil
	file_net_proto_depIdxs = nil
}
