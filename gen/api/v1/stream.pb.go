// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: api/v1/stream.proto

package v1

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

type PingPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingPong) Reset() {
	*x = PingPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPong) ProtoMessage() {}

func (x *PingPong) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPong.ProtoReflect.Descriptor instead.
func (*PingPong) Descriptor() ([]byte, []int) {
	return file_api_v1_stream_proto_rawDescGZIP(), []int{0}
}

type PingPong_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PingPong_Request) Reset() {
	*x = PingPong_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPong_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPong_Request) ProtoMessage() {}

func (x *PingPong_Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPong_Request.ProtoReflect.Descriptor instead.
func (*PingPong_Request) Descriptor() ([]byte, []int) {
	return file_api_v1_stream_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PingPong_Request) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PingPong_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greet string `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
}

func (x *PingPong_Response) Reset() {
	*x = PingPong_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPong_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPong_Response) ProtoMessage() {}

func (x *PingPong_Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPong_Response.ProtoReflect.Descriptor instead.
func (*PingPong_Response) Descriptor() ([]byte, []int) {
	return file_api_v1_stream_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PingPong_Response) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

var File_api_v1_stream_proto protoreflect.FileDescriptor

var file_api_v1_stream_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x4d, 0x0a, 0x08, 0x50, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x1a, 0x1f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x32, 0x49, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x3b, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x74, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x42,
	0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x63, 0x6a, 0x2d, 0x79,
	0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x41, 0x70, 0x69, 0xca, 0x02, 0x03, 0x41,
	0x70, 0x69, 0xe2, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1_stream_proto_rawDescOnce sync.Once
	file_api_v1_stream_proto_rawDescData = file_api_v1_stream_proto_rawDesc
)

func file_api_v1_stream_proto_rawDescGZIP() []byte {
	file_api_v1_stream_proto_rawDescOnce.Do(func() {
		file_api_v1_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_stream_proto_rawDescData)
	})
	return file_api_v1_stream_proto_rawDescData
}

var file_api_v1_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_stream_proto_goTypes = []interface{}{
	(*PingPong)(nil),          // 0: api.PingPong
	(*PingPong_Request)(nil),  // 1: api.PingPong.Request
	(*PingPong_Response)(nil), // 2: api.PingPong.Response
}
var file_api_v1_stream_proto_depIdxs = []int32{
	1, // 0: api.StreamDemo.PingPong:input_type -> api.PingPong.Request
	2, // 1: api.StreamDemo.PingPong:output_type -> api.PingPong.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_stream_proto_init() }
func file_api_v1_stream_proto_init() {
	if File_api_v1_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPong); i {
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
		file_api_v1_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPong_Request); i {
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
		file_api_v1_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPong_Response); i {
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
			RawDescriptor: file_api_v1_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_stream_proto_goTypes,
		DependencyIndexes: file_api_v1_stream_proto_depIdxs,
		MessageInfos:      file_api_v1_stream_proto_msgTypes,
	}.Build()
	File_api_v1_stream_proto = out.File
	file_api_v1_stream_proto_rawDesc = nil
	file_api_v1_stream_proto_goTypes = nil
	file_api_v1_stream_proto_depIdxs = nil
}
