// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: invoice/v1/invoice_service.proto

package invoicepb

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

type SendDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Invoice `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendDataRequest) Reset() {
	*x = SendDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_v1_invoice_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDataRequest) ProtoMessage() {}

func (x *SendDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_v1_invoice_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDataRequest.ProtoReflect.Descriptor instead.
func (*SendDataRequest) Descriptor() ([]byte, []int) {
	return file_invoice_v1_invoice_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendDataRequest) GetData() *Invoice {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SendDataResponse) Reset() {
	*x = SendDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_v1_invoice_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDataResponse) ProtoMessage() {}

func (x *SendDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_v1_invoice_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDataResponse.ProtoReflect.Descriptor instead.
func (*SendDataResponse) Descriptor() ([]byte, []int) {
	return file_invoice_v1_invoice_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendDataResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_invoice_v1_invoice_service_proto protoreflect.FileDescriptor

var file_invoice_v1_invoice_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x32, 0x59, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b,
	0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x6e, 0x69, 0x63,
	0x68, 0x38, 0x39, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_invoice_v1_invoice_service_proto_rawDescOnce sync.Once
	file_invoice_v1_invoice_service_proto_rawDescData = file_invoice_v1_invoice_service_proto_rawDesc
)

func file_invoice_v1_invoice_service_proto_rawDescGZIP() []byte {
	file_invoice_v1_invoice_service_proto_rawDescOnce.Do(func() {
		file_invoice_v1_invoice_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoice_v1_invoice_service_proto_rawDescData)
	})
	return file_invoice_v1_invoice_service_proto_rawDescData
}

var file_invoice_v1_invoice_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_invoice_v1_invoice_service_proto_goTypes = []interface{}{
	(*SendDataRequest)(nil),  // 0: invoice.v1.SendDataRequest
	(*SendDataResponse)(nil), // 1: invoice.v1.SendDataResponse
	(*Invoice)(nil),          // 2: invoice.v1.Invoice
}
var file_invoice_v1_invoice_service_proto_depIdxs = []int32{
	2, // 0: invoice.v1.SendDataRequest.data:type_name -> invoice.v1.Invoice
	0, // 1: invoice.v1.InvoiceService.SendData:input_type -> invoice.v1.SendDataRequest
	1, // 2: invoice.v1.InvoiceService.SendData:output_type -> invoice.v1.SendDataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_invoice_v1_invoice_service_proto_init() }
func file_invoice_v1_invoice_service_proto_init() {
	if File_invoice_v1_invoice_service_proto != nil {
		return
	}
	file_invoice_v1_invoice_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_invoice_v1_invoice_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDataRequest); i {
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
		file_invoice_v1_invoice_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDataResponse); i {
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
			RawDescriptor: file_invoice_v1_invoice_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invoice_v1_invoice_service_proto_goTypes,
		DependencyIndexes: file_invoice_v1_invoice_service_proto_depIdxs,
		MessageInfos:      file_invoice_v1_invoice_service_proto_msgTypes,
	}.Build()
	File_invoice_v1_invoice_service_proto = out.File
	file_invoice_v1_invoice_service_proto_rawDesc = nil
	file_invoice_v1_invoice_service_proto_goTypes = nil
	file_invoice_v1_invoice_service_proto_depIdxs = nil
}