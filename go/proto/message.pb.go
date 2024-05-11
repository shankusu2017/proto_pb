// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: message.proto

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

// 消息ID-唯一
type MsgID int32

const (
	MsgID_RepeaterServerInfoReq MsgID = 0
	MsgID_RepeaterServerInfoRsp MsgID = 1
	MsgID_EventPost             MsgID = 2
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0: "RepeaterServerInfoReq",
		1: "RepeaterServerInfoRsp",
		2: "EventPost",
	}
	MsgID_value = map[string]int32{
		"RepeaterServerInfoReq": 0,
		"RepeaterServerInfoRsp": 1,
		"EventPost":             2,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}

func (x MsgID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgID) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MsgID) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MsgID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgID.Descriptor instead.
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

// RepeaterServerInfoReq 请求代理服务器的信息
type MsgRepeaterServerInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *Machine `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 自己的node信息
}

func (x *MsgRepeaterServerInfoReq) Reset() {
	*x = MsgRepeaterServerInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRepeaterServerInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRepeaterServerInfoReq) ProtoMessage() {}

func (x *MsgRepeaterServerInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRepeaterServerInfoReq.ProtoReflect.Descriptor instead.
func (*MsgRepeaterServerInfoReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *MsgRepeaterServerInfoReq) GetId() *Machine {
	if x != nil {
		return x.Id
	}
	return nil
}

// RepeaterServerInfoRsp 代理服务器信息返回
type MsgRepeaterServerInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*RepeaterServerNode `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"` // 当前所有 serverNode 的信息
}

func (x *MsgRepeaterServerInfoRsp) Reset() {
	*x = MsgRepeaterServerInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRepeaterServerInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRepeaterServerInfoRsp) ProtoMessage() {}

func (x *MsgRepeaterServerInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRepeaterServerInfoRsp.ProtoReflect.Descriptor instead.
func (*MsgRepeaterServerInfoRsp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *MsgRepeaterServerInfoRsp) GetServers() []*RepeaterServerNode {
	if x != nil {
		return x.Servers
	}
	return nil
}

// 事件报告
type MsgEventPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event Event `protobuf:"varint,1,opt,name=event,proto3,enum=event.Event" json:"event,omitempty"`
	Ts    int64 `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"` // 发生时间(unix-time)
}

func (x *MsgEventPost) Reset() {
	*x = MsgEventPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgEventPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgEventPost) ProtoMessage() {}

func (x *MsgEventPost) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgEventPost.ProtoReflect.Descriptor instead.
func (*MsgEventPost) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *MsgEventPost) GetEvent() Event {
	if x != nil {
		return x.Event
	}
	return Event_STARTED
}

func (x *MsgEventPost) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x20, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x52, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x36,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x42, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x2a, 0x4c, 0x0a, 0x05, 0x4d, 0x73,
	0x67, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x10, 0x02, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_proto_goTypes = []interface{}{
	(MsgID)(0),                       // 0: message.MsgID
	(*MsgRepeaterServerInfoReq)(nil), // 1: message.MsgRepeaterServerInfoReq
	(*MsgRepeaterServerInfoRsp)(nil), // 2: message.MsgRepeaterServerInfoRsp
	(*MsgEventPost)(nil),             // 3: message.MsgEventPost
	(*Machine)(nil),                  // 4: machine.Machine
	(*RepeaterServerNode)(nil),       // 5: repeater.RepeaterServerNode
	(Event)(0),                       // 6: event.Event
}
var file_message_proto_depIdxs = []int32{
	4, // 0: message.MsgRepeaterServerInfoReq.id:type_name -> machine.Machine
	5, // 1: message.MsgRepeaterServerInfoRsp.servers:type_name -> repeater.RepeaterServerNode
	6, // 2: message.MsgEventPost.event:type_name -> event.Event
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_event_proto_init()
	file_machine_proto_init()
	file_repeater_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRepeaterServerInfoReq); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRepeaterServerInfoRsp); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgEventPost); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
