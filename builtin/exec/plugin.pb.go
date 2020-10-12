// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: waypoint/builtin/exec/plugin.proto

package exec

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Input is the input type used for the exec plugin.
type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]*Input_Value `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_exec_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_exec_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_exec_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetData() map[string]*Input_Value {
	if x != nil {
		return x.Data
	}
	return nil
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_exec_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_exec_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_exec_plugin_proto_rawDescGZIP(), []int{1}
}

type Input_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*Input_Value_Text
	Value isInput_Value_Value `protobuf_oneof:"value"`
}

func (x *Input_Value) Reset() {
	*x = Input_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_exec_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input_Value) ProtoMessage() {}

func (x *Input_Value) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_exec_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input_Value.ProtoReflect.Descriptor instead.
func (*Input_Value) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_exec_plugin_proto_rawDescGZIP(), []int{0, 1}
}

func (m *Input_Value) GetValue() isInput_Value_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Input_Value) GetText() string {
	if x, ok := x.GetValue().(*Input_Value_Text); ok {
		return x.Text
	}
	return ""
}

type isInput_Value_Value interface {
	isInput_Value_Value()
}

type Input_Value_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

func (*Input_Value_Text) isInput_Value_Value() {}

var File_waypoint_builtin_exec_plugin_proto protoreflect.FileDescriptor

var file_waypoint_builtin_exec_plugin_proto_rawDesc = []byte{
	0x0a, 0x22, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74,
	0x69, 0x6e, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x78, 0x65, 0x63, 0x22, 0xa6, 0x01, 0x0a, 0x05, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x4a, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x26, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x17, 0x5a, 0x15, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x74, 0x69, 0x6e, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_waypoint_builtin_exec_plugin_proto_rawDescOnce sync.Once
	file_waypoint_builtin_exec_plugin_proto_rawDescData = file_waypoint_builtin_exec_plugin_proto_rawDesc
)

func file_waypoint_builtin_exec_plugin_proto_rawDescGZIP() []byte {
	file_waypoint_builtin_exec_plugin_proto_rawDescOnce.Do(func() {
		file_waypoint_builtin_exec_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_waypoint_builtin_exec_plugin_proto_rawDescData)
	})
	return file_waypoint_builtin_exec_plugin_proto_rawDescData
}

var file_waypoint_builtin_exec_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_waypoint_builtin_exec_plugin_proto_goTypes = []interface{}{
	(*Input)(nil),       // 0: exec.Input
	(*Deployment)(nil),  // 1: exec.Deployment
	nil,                 // 2: exec.Input.DataEntry
	(*Input_Value)(nil), // 3: exec.Input.Value
}
var file_waypoint_builtin_exec_plugin_proto_depIdxs = []int32{
	2, // 0: exec.Input.data:type_name -> exec.Input.DataEntry
	3, // 1: exec.Input.DataEntry.value:type_name -> exec.Input.Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_waypoint_builtin_exec_plugin_proto_init() }
func file_waypoint_builtin_exec_plugin_proto_init() {
	if File_waypoint_builtin_exec_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waypoint_builtin_exec_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_waypoint_builtin_exec_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_waypoint_builtin_exec_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input_Value); i {
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
	file_waypoint_builtin_exec_plugin_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Input_Value_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waypoint_builtin_exec_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waypoint_builtin_exec_plugin_proto_goTypes,
		DependencyIndexes: file_waypoint_builtin_exec_plugin_proto_depIdxs,
		MessageInfos:      file_waypoint_builtin_exec_plugin_proto_msgTypes,
	}.Build()
	File_waypoint_builtin_exec_plugin_proto = out.File
	file_waypoint_builtin_exec_plugin_proto_rawDesc = nil
	file_waypoint_builtin_exec_plugin_proto_goTypes = nil
	file_waypoint_builtin_exec_plugin_proto_depIdxs = nil
}
