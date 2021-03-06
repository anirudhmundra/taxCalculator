// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: taxCalculatorService.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TaxCalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tds                    float32 `protobuf:"fixed32,1,opt,name=tds,proto3" json:"tds,omitempty"`
	RentalAllowance        float32 `protobuf:"fixed32,2,opt,name=rentalAllowance,proto3" json:"rentalAllowance,omitempty"`
	IncomeFromOtherSources float32 `protobuf:"fixed32,3,opt,name=incomeFromOtherSources,proto3" json:"incomeFromOtherSources,omitempty"`
	LoanAmount             float32 `protobuf:"fixed32,4,opt,name=loanAmount,proto3" json:"loanAmount,omitempty"`
	TotalIncome            float32 `protobuf:"fixed32,5,opt,name=totalIncome,proto3" json:"totalIncome,omitempty"`
}

func (x *TaxCalculatorRequest) Reset() {
	*x = TaxCalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxCalculatorService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxCalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxCalculatorRequest) ProtoMessage() {}

func (x *TaxCalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taxCalculatorService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxCalculatorRequest.ProtoReflect.Descriptor instead.
func (*TaxCalculatorRequest) Descriptor() ([]byte, []int) {
	return file_taxCalculatorService_proto_rawDescGZIP(), []int{0}
}

func (x *TaxCalculatorRequest) GetTds() float32 {
	if x != nil {
		return x.Tds
	}
	return 0
}

func (x *TaxCalculatorRequest) GetRentalAllowance() float32 {
	if x != nil {
		return x.RentalAllowance
	}
	return 0
}

func (x *TaxCalculatorRequest) GetIncomeFromOtherSources() float32 {
	if x != nil {
		return x.IncomeFromOtherSources
	}
	return 0
}

func (x *TaxCalculatorRequest) GetLoanAmount() float32 {
	if x != nil {
		return x.LoanAmount
	}
	return 0
}

func (x *TaxCalculatorRequest) GetTotalIncome() float32 {
	if x != nil {
		return x.TotalIncome
	}
	return 0
}

type TaxCalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TaxCalculatorResponse) Reset() {
	*x = TaxCalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taxCalculatorService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaxCalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaxCalculatorResponse) ProtoMessage() {}

func (x *TaxCalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taxCalculatorService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaxCalculatorResponse.ProtoReflect.Descriptor instead.
func (*TaxCalculatorResponse) Descriptor() ([]byte, []int) {
	return file_taxCalculatorService_proto_rawDescGZIP(), []int{1}
}

func (x *TaxCalculatorResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_taxCalculatorService_proto protoreflect.FileDescriptor

var file_taxCalculatorService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x61, 0x78, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x14, 0x54, 0x61, 0x78, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x0f,
	0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x16, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x6c, 0x6f, 0x61, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x6c, 0x6f, 0x61, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65,
	0x22, 0x2f, 0x0a, 0x15, 0x54, 0x61, 0x78, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x72, 0x0a, 0x14, 0x54, 0x61, 0x78, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x61, 0x78, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x61, 0x78, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x61, 0x78, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x78, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taxCalculatorService_proto_rawDescOnce sync.Once
	file_taxCalculatorService_proto_rawDescData = file_taxCalculatorService_proto_rawDesc
)

func file_taxCalculatorService_proto_rawDescGZIP() []byte {
	file_taxCalculatorService_proto_rawDescOnce.Do(func() {
		file_taxCalculatorService_proto_rawDescData = protoimpl.X.CompressGZIP(file_taxCalculatorService_proto_rawDescData)
	})
	return file_taxCalculatorService_proto_rawDescData
}

var file_taxCalculatorService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_taxCalculatorService_proto_goTypes = []interface{}{
	(*TaxCalculatorRequest)(nil),  // 0: proto.TaxCalculatorRequest
	(*TaxCalculatorResponse)(nil), // 1: proto.TaxCalculatorResponse
}
var file_taxCalculatorService_proto_depIdxs = []int32{
	0, // 0: proto.TaxCalculatorService.CalculateTax:input_type -> proto.TaxCalculatorRequest
	1, // 1: proto.TaxCalculatorService.CalculateTax:output_type -> proto.TaxCalculatorResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_taxCalculatorService_proto_init() }
func file_taxCalculatorService_proto_init() {
	if File_taxCalculatorService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_taxCalculatorService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxCalculatorRequest); i {
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
		file_taxCalculatorService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaxCalculatorResponse); i {
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
			RawDescriptor: file_taxCalculatorService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taxCalculatorService_proto_goTypes,
		DependencyIndexes: file_taxCalculatorService_proto_depIdxs,
		MessageInfos:      file_taxCalculatorService_proto_msgTypes,
	}.Build()
	File_taxCalculatorService_proto = out.File
	file_taxCalculatorService_proto_rawDesc = nil
	file_taxCalculatorService_proto_goTypes = nil
	file_taxCalculatorService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaxCalculatorServiceClient is the client API for TaxCalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaxCalculatorServiceClient interface {
	CalculateTax(ctx context.Context, in *TaxCalculatorRequest, opts ...grpc.CallOption) (*TaxCalculatorResponse, error)
}

type taxCalculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaxCalculatorServiceClient(cc grpc.ClientConnInterface) TaxCalculatorServiceClient {
	return &taxCalculatorServiceClient{cc}
}

func (c *taxCalculatorServiceClient) CalculateTax(ctx context.Context, in *TaxCalculatorRequest, opts ...grpc.CallOption) (*TaxCalculatorResponse, error) {
	out := new(TaxCalculatorResponse)
	err := c.cc.Invoke(ctx, "/proto.TaxCalculatorService/CalculateTax", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaxCalculatorServiceServer is the server API for TaxCalculatorService service.
type TaxCalculatorServiceServer interface {
	CalculateTax(context.Context, *TaxCalculatorRequest) (*TaxCalculatorResponse, error)
}

// UnimplementedTaxCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaxCalculatorServiceServer struct {
}

func (*UnimplementedTaxCalculatorServiceServer) CalculateTax(context.Context, *TaxCalculatorRequest) (*TaxCalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateTax not implemented")
}

func RegisterTaxCalculatorServiceServer(s *grpc.Server, srv TaxCalculatorServiceServer) {
	s.RegisterService(&_TaxCalculatorService_serviceDesc, srv)
}

func _TaxCalculatorService_CalculateTax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaxCalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxCalculatorServiceServer).CalculateTax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TaxCalculatorService/CalculateTax",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxCalculatorServiceServer).CalculateTax(ctx, req.(*TaxCalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaxCalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TaxCalculatorService",
	HandlerType: (*TaxCalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateTax",
			Handler:    _TaxCalculatorService_CalculateTax_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taxCalculatorService.proto",
}
