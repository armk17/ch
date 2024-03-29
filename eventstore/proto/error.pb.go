// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: proto/error.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Errors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code    string       `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Status  int32        `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Title   string       `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Details []*anypb.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Errors) Reset() {
	*x = Errors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errors) ProtoMessage() {}

func (x *Errors) ProtoReflect() protoreflect.Message {
	mi := &file_proto_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errors.ProtoReflect.Descriptor instead.
func (*Errors) Descriptor() ([]byte, []int) {
	return file_proto_error_proto_rawDescGZIP(), []int{0}
}

func (x *Errors) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Errors) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Errors) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Errors) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Errors) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type Errors_Detail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Messages []string `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Errors_Detail) Reset() {
	*x = Errors_Detail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errors_Detail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errors_Detail) ProtoMessage() {}

func (x *Errors_Detail) ProtoReflect() protoreflect.Message {
	mi := &file_proto_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errors_Detail.ProtoReflect.Descriptor instead.
func (*Errors_Detail) Descriptor() ([]byte, []int) {
	return file_proto_error_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Errors_Detail) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Errors_Detail) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_proto_error_proto protoreflect.FileDescriptor

var file_proto_error_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72,
	0x6b, 0x31, 0x37, 0x39, 0x30, 0x2f, 0x63, 0x68, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_error_proto_rawDescOnce sync.Once
	file_proto_error_proto_rawDescData = file_proto_error_proto_rawDesc
)

func file_proto_error_proto_rawDescGZIP() []byte {
	file_proto_error_proto_rawDescOnce.Do(func() {
		file_proto_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_error_proto_rawDescData)
	})
	return file_proto_error_proto_rawDescData
}

var file_proto_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_error_proto_goTypes = []interface{}{
	(*Errors)(nil),        // 0: common.Errors
	(*Errors_Detail)(nil), // 1: common.Errors.Detail
	(*anypb.Any)(nil),     // 2: google.protobuf.Any
}
var file_proto_error_proto_depIdxs = []int32{
	2, // 0: common.Errors.details:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_error_proto_init() }
func file_proto_error_proto_init() {
	if File_proto_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errors); i {
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
		file_proto_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errors_Detail); i {
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
			RawDescriptor: file_proto_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_error_proto_goTypes,
		DependencyIndexes: file_proto_error_proto_depIdxs,
		MessageInfos:      file_proto_error_proto_msgTypes,
	}.Build()
	File_proto_error_proto = out.File
	file_proto_error_proto_rawDesc = nil
	file_proto_error_proto_goTypes = nil
	file_proto_error_proto_depIdxs = nil
}
