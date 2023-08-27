// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: shorter.proto

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

type UrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullURL string `protobuf:"bytes,1,opt,name=fullURL,proto3" json:"fullURL,omitempty"`
}

func (x *UrlRequest) Reset() {
	*x = UrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlRequest) ProtoMessage() {}

func (x *UrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlRequest.ProtoReflect.Descriptor instead.
func (*UrlRequest) Descriptor() ([]byte, []int) {
	return file_shorter_proto_rawDescGZIP(), []int{0}
}

func (x *UrlRequest) GetFullURL() string {
	if x != nil {
		return x.FullURL
	}
	return ""
}

type UrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UrlResponse) Reset() {
	*x = UrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResponse) ProtoMessage() {}

func (x *UrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shorter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResponse.ProtoReflect.Descriptor instead.
func (*UrlResponse) Descriptor() ([]byte, []int) {
	return file_shorter_proto_rawDescGZIP(), []int{1}
}

func (x *UrlResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *UrlResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_shorter_proto protoreflect.FileDescriptor

var file_shorter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x75, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6c, 0x6c, 0x55, 0x52, 0x4c, 0x22, 0x3d,
	0x0a, 0x0b, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x6c, 0x0a,
	0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x68, 0x65, 0x72, 0x61, 0x73,
	0x79, 0x6d, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shorter_proto_rawDescOnce sync.Once
	file_shorter_proto_rawDescData = file_shorter_proto_rawDesc
)

func file_shorter_proto_rawDescGZIP() []byte {
	file_shorter_proto_rawDescOnce.Do(func() {
		file_shorter_proto_rawDescData = protoimpl.X.CompressGZIP(file_shorter_proto_rawDescData)
	})
	return file_shorter_proto_rawDescData
}

var file_shorter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shorter_proto_goTypes = []interface{}{
	(*UrlRequest)(nil),  // 0: proto.UrlRequest
	(*UrlResponse)(nil), // 1: proto.UrlResponse
}
var file_shorter_proto_depIdxs = []int32{
	0, // 0: proto.Shorter.Create:input_type -> proto.UrlRequest
	0, // 1: proto.Shorter.Get:input_type -> proto.UrlRequest
	1, // 2: proto.Shorter.Create:output_type -> proto.UrlResponse
	1, // 3: proto.Shorter.Get:output_type -> proto.UrlResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shorter_proto_init() }
func file_shorter_proto_init() {
	if File_shorter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shorter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlRequest); i {
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
		file_shorter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlResponse); i {
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
			RawDescriptor: file_shorter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shorter_proto_goTypes,
		DependencyIndexes: file_shorter_proto_depIdxs,
		MessageInfos:      file_shorter_proto_msgTypes,
	}.Build()
	File_shorter_proto = out.File
	file_shorter_proto_rawDesc = nil
	file_shorter_proto_goTypes = nil
	file_shorter_proto_depIdxs = nil
}