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

// 事件报告
type MsgEventPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   Event    `protobuf:"varint,1,opt,name=event,proto3,enum=event.Event" json:"event,omitempty"` // 事件类型
	Ts      int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`                        // 发生时间(unix-time-毫秒)
	Machine *Machine `protobuf:"bytes,3,opt,name=machine,proto3" json:"machine,omitempty"`               // 机器信息(必须唯一)
	Node    *Node    `protobuf:"bytes,4,opt,name=node,proto3" json:"node,omitempty"`                     // 节点消息(软件版本，自己的角色类型等)
}

func (x *MsgEventPost) Reset() {
	*x = MsgEventPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgEventPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgEventPost) ProtoMessage() {}

func (x *MsgEventPost) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MsgEventPost.ProtoReflect.Descriptor instead.
func (*MsgEventPost) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
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

func (x *MsgEventPost) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

func (x *MsgEventPost) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

// RepeaterServerInfoReq 请求代理服务器的信息
type MsgRepeaterServerInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"` // 自己的node信息
}

func (x *MsgRepeaterServerInfoReq) Reset() {
	*x = MsgRepeaterServerInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRepeaterServerInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRepeaterServerInfoReq) ProtoMessage() {}

func (x *MsgRepeaterServerInfoReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MsgRepeaterServerInfoReq.ProtoReflect.Descriptor instead.
func (*MsgRepeaterServerInfoReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *MsgRepeaterServerInfoReq) GetMachine() *Machine {
	if x != nil {
		return x.Machine
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
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRepeaterServerInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRepeaterServerInfoRsp) ProtoMessage() {}

func (x *MsgRepeaterServerInfoRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MsgRepeaterServerInfoRsp.ProtoReflect.Descriptor instead.
func (*MsgRepeaterServerInfoRsp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *MsgRepeaterServerInfoRsp) GetServers() []*RepeaterServerNode {
	if x != nil {
		return x.Servers
	}
	return nil
}

// KeepAlive 消息
type MsgKeepAlive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`       // 角色 pac 还是 repeater
	Machine *Machine `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"` // 机器信息(必须唯一)
	Node    *Node    `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`       // 节点消息(软件版本，自己的角色类型等)
}

func (x *MsgKeepAlive) Reset() {
	*x = MsgKeepAlive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgKeepAlive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgKeepAlive) ProtoMessage() {}

func (x *MsgKeepAlive) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgKeepAlive.ProtoReflect.Descriptor instead.
func (*MsgKeepAlive) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *MsgKeepAlive) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *MsgKeepAlive) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

func (x *MsgKeepAlive) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x22, 0x46, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a,
	0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x22, 0x52, 0x0a, 0x18, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x6e, 0x0a,
	0x0c, 0x4d, 0x73, 0x67, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x1e, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_proto_goTypes = []interface{}{
	(*MsgEventPost)(nil),             // 0: message.MsgEventPost
	(*MsgRepeaterServerInfoReq)(nil), // 1: message.MsgRepeaterServerInfoReq
	(*MsgRepeaterServerInfoRsp)(nil), // 2: message.MsgRepeaterServerInfoRsp
	(*MsgKeepAlive)(nil),             // 3: message.MsgKeepAlive
	(Event)(0),                       // 4: event.Event
	(*Machine)(nil),                  // 5: machine.Machine
	(*Node)(nil),                     // 6: node.Node
	(*RepeaterServerNode)(nil),       // 7: repeater.RepeaterServerNode
}
var file_message_proto_depIdxs = []int32{
	4, // 0: message.MsgEventPost.event:type_name -> event.Event
	5, // 1: message.MsgEventPost.machine:type_name -> machine.Machine
	6, // 2: message.MsgEventPost.node:type_name -> node.Node
	5, // 3: message.MsgRepeaterServerInfoReq.machine:type_name -> machine.Machine
	7, // 4: message.MsgRepeaterServerInfoRsp.servers:type_name -> repeater.RepeaterServerNode
	5, // 5: message.MsgKeepAlive.machine:type_name -> machine.Machine
	6, // 6: message.MsgKeepAlive.node:type_name -> node.Node
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_event_proto_init()
	file_machine_proto_init()
	file_repeater_proto_init()
	file_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgKeepAlive); i {
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
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
