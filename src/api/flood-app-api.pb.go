// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flood-app-api.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckAlertsRequest struct {
	AlertMode            uint64   `protobuf:"varint,1,opt,name=alertMode,proto3" json:"alertMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAlertsRequest) Reset()         { *m = CheckAlertsRequest{} }
func (m *CheckAlertsRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAlertsRequest) ProtoMessage()    {}
func (*CheckAlertsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{0}
}
func (m *CheckAlertsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAlertsRequest.Unmarshal(m, b)
}
func (m *CheckAlertsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAlertsRequest.Marshal(b, m, deterministic)
}
func (dst *CheckAlertsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAlertsRequest.Merge(dst, src)
}
func (m *CheckAlertsRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAlertsRequest.Size(m)
}
func (m *CheckAlertsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAlertsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAlertsRequest proto.InternalMessageInfo

func (m *CheckAlertsRequest) GetAlertMode() uint64 {
	if m != nil {
		return m.AlertMode
	}
	return 0
}

type CheckAlertsResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAlertsResponse) Reset()         { *m = CheckAlertsResponse{} }
func (m *CheckAlertsResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAlertsResponse) ProtoMessage()    {}
func (*CheckAlertsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{1}
}
func (m *CheckAlertsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAlertsResponse.Unmarshal(m, b)
}
func (m *CheckAlertsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAlertsResponse.Marshal(b, m, deterministic)
}
func (dst *CheckAlertsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAlertsResponse.Merge(dst, src)
}
func (m *CheckAlertsResponse) XXX_Size() int {
	return xxx_messageInfo_CheckAlertsResponse.Size(m)
}
func (m *CheckAlertsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAlertsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAlertsResponse proto.InternalMessageInfo

func (m *CheckAlertsResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type TestAlertsRequest struct {
	TestMessage          string   `protobuf:"bytes,1,opt,name=testMessage,proto3" json:"testMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestAlertsRequest) Reset()         { *m = TestAlertsRequest{} }
func (m *TestAlertsRequest) String() string { return proto.CompactTextString(m) }
func (*TestAlertsRequest) ProtoMessage()    {}
func (*TestAlertsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{2}
}
func (m *TestAlertsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAlertsRequest.Unmarshal(m, b)
}
func (m *TestAlertsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAlertsRequest.Marshal(b, m, deterministic)
}
func (dst *TestAlertsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAlertsRequest.Merge(dst, src)
}
func (m *TestAlertsRequest) XXX_Size() int {
	return xxx_messageInfo_TestAlertsRequest.Size(m)
}
func (m *TestAlertsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAlertsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestAlertsRequest proto.InternalMessageInfo

func (m *TestAlertsRequest) GetTestMessage() string {
	if m != nil {
		return m.TestMessage
	}
	return ""
}

type TestAlertsResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestAlertsResponse) Reset()         { *m = TestAlertsResponse{} }
func (m *TestAlertsResponse) String() string { return proto.CompactTextString(m) }
func (*TestAlertsResponse) ProtoMessage()    {}
func (*TestAlertsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{3}
}
func (m *TestAlertsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAlertsResponse.Unmarshal(m, b)
}
func (m *TestAlertsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAlertsResponse.Marshal(b, m, deterministic)
}
func (dst *TestAlertsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAlertsResponse.Merge(dst, src)
}
func (m *TestAlertsResponse) XXX_Size() int {
	return xxx_messageInfo_TestAlertsResponse.Size(m)
}
func (m *TestAlertsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAlertsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestAlertsResponse proto.InternalMessageInfo

func (m *TestAlertsResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

// Mute SMS = "SMS True/False"
// Mute Email = "Email True/False"
type MuteAlertRequest struct {
	MuteAlertType        string   `protobuf:"bytes,1,opt,name=muteAlertType,proto3" json:"muteAlertType,omitempty"`
	MuteAlertStatus      bool     `protobuf:"varint,2,opt,name=muteAlertStatus,proto3" json:"muteAlertStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MuteAlertRequest) Reset()         { *m = MuteAlertRequest{} }
func (m *MuteAlertRequest) String() string { return proto.CompactTextString(m) }
func (*MuteAlertRequest) ProtoMessage()    {}
func (*MuteAlertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{4}
}
func (m *MuteAlertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MuteAlertRequest.Unmarshal(m, b)
}
func (m *MuteAlertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MuteAlertRequest.Marshal(b, m, deterministic)
}
func (dst *MuteAlertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MuteAlertRequest.Merge(dst, src)
}
func (m *MuteAlertRequest) XXX_Size() int {
	return xxx_messageInfo_MuteAlertRequest.Size(m)
}
func (m *MuteAlertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MuteAlertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MuteAlertRequest proto.InternalMessageInfo

func (m *MuteAlertRequest) GetMuteAlertType() string {
	if m != nil {
		return m.MuteAlertType
	}
	return ""
}

func (m *MuteAlertRequest) GetMuteAlertStatus() bool {
	if m != nil {
		return m.MuteAlertStatus
	}
	return false
}

// Returns what was sucessfully muted
type MuteAlertResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MuteAlertResponse) Reset()         { *m = MuteAlertResponse{} }
func (m *MuteAlertResponse) String() string { return proto.CompactTextString(m) }
func (*MuteAlertResponse) ProtoMessage()    {}
func (*MuteAlertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{5}
}
func (m *MuteAlertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MuteAlertResponse.Unmarshal(m, b)
}
func (m *MuteAlertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MuteAlertResponse.Marshal(b, m, deterministic)
}
func (dst *MuteAlertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MuteAlertResponse.Merge(dst, src)
}
func (m *MuteAlertResponse) XXX_Size() int {
	return xxx_messageInfo_MuteAlertResponse.Size(m)
}
func (m *MuteAlertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MuteAlertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MuteAlertResponse proto.InternalMessageInfo

func (m *MuteAlertResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type GetForcastRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetForcastRequest) Reset()         { *m = GetForcastRequest{} }
func (m *GetForcastRequest) String() string { return proto.CompactTextString(m) }
func (*GetForcastRequest) ProtoMessage()    {}
func (*GetForcastRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{6}
}
func (m *GetForcastRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetForcastRequest.Unmarshal(m, b)
}
func (m *GetForcastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetForcastRequest.Marshal(b, m, deterministic)
}
func (dst *GetForcastRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetForcastRequest.Merge(dst, src)
}
func (m *GetForcastRequest) XXX_Size() int {
	return xxx_messageInfo_GetForcastRequest.Size(m)
}
func (m *GetForcastRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetForcastRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetForcastRequest proto.InternalMessageInfo

type GetForcastResponse struct {
	Forcast              string   `protobuf:"bytes,1,opt,name=forcast,proto3" json:"forcast,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetForcastResponse) Reset()         { *m = GetForcastResponse{} }
func (m *GetForcastResponse) String() string { return proto.CompactTextString(m) }
func (*GetForcastResponse) ProtoMessage()    {}
func (*GetForcastResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{7}
}
func (m *GetForcastResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetForcastResponse.Unmarshal(m, b)
}
func (m *GetForcastResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetForcastResponse.Marshal(b, m, deterministic)
}
func (dst *GetForcastResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetForcastResponse.Merge(dst, src)
}
func (m *GetForcastResponse) XXX_Size() int {
	return xxx_messageInfo_GetForcastResponse.Size(m)
}
func (m *GetForcastResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetForcastResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetForcastResponse proto.InternalMessageInfo

func (m *GetForcastResponse) GetForcast() string {
	if m != nil {
		return m.Forcast
	}
	return ""
}

type UpdateForcastRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateForcastRequest) Reset()         { *m = UpdateForcastRequest{} }
func (m *UpdateForcastRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateForcastRequest) ProtoMessage()    {}
func (*UpdateForcastRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{8}
}
func (m *UpdateForcastRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateForcastRequest.Unmarshal(m, b)
}
func (m *UpdateForcastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateForcastRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateForcastRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateForcastRequest.Merge(dst, src)
}
func (m *UpdateForcastRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateForcastRequest.Size(m)
}
func (m *UpdateForcastRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateForcastRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateForcastRequest proto.InternalMessageInfo

type UpdateForcastResponse struct {
	UpdatedForcast       string   `protobuf:"bytes,1,opt,name=updatedForcast,proto3" json:"updatedForcast,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateForcastResponse) Reset()         { *m = UpdateForcastResponse{} }
func (m *UpdateForcastResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateForcastResponse) ProtoMessage()    {}
func (*UpdateForcastResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_flood_app_api_4106f79d85f73408, []int{9}
}
func (m *UpdateForcastResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateForcastResponse.Unmarshal(m, b)
}
func (m *UpdateForcastResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateForcastResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateForcastResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateForcastResponse.Merge(dst, src)
}
func (m *UpdateForcastResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateForcastResponse.Size(m)
}
func (m *UpdateForcastResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateForcastResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateForcastResponse proto.InternalMessageInfo

func (m *UpdateForcastResponse) GetUpdatedForcast() string {
	if m != nil {
		return m.UpdatedForcast
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckAlertsRequest)(nil), "pb.checkAlertsRequest")
	proto.RegisterType((*CheckAlertsResponse)(nil), "pb.checkAlertsResponse")
	proto.RegisterType((*TestAlertsRequest)(nil), "pb.testAlertsRequest")
	proto.RegisterType((*TestAlertsResponse)(nil), "pb.testAlertsResponse")
	proto.RegisterType((*MuteAlertRequest)(nil), "pb.muteAlertRequest")
	proto.RegisterType((*MuteAlertResponse)(nil), "pb.muteAlertResponse")
	proto.RegisterType((*GetForcastRequest)(nil), "pb.getForcastRequest")
	proto.RegisterType((*GetForcastResponse)(nil), "pb.getForcastResponse")
	proto.RegisterType((*UpdateForcastRequest)(nil), "pb.updateForcastRequest")
	proto.RegisterType((*UpdateForcastResponse)(nil), "pb.updateForcastResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FloodAlertAppServiceClient is the client API for FloodAlertAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FloodAlertAppServiceClient interface {
	CheckAlerts(ctx context.Context, in *CheckAlertsRequest, opts ...grpc.CallOption) (*CheckAlertsResponse, error)
	TestAlerts(ctx context.Context, in *TestAlertsRequest, opts ...grpc.CallOption) (*TestAlertsResponse, error)
	MuteAlerts(ctx context.Context, in *MuteAlertRequest, opts ...grpc.CallOption) (*MuteAlertResponse, error)
	GetForcast(ctx context.Context, in *GetForcastRequest, opts ...grpc.CallOption) (*GetForcastResponse, error)
	UpdateForcast(ctx context.Context, in *UpdateForcastRequest, opts ...grpc.CallOption) (*UpdateForcastResponse, error)
}

type floodAlertAppServiceClient struct {
	cc *grpc.ClientConn
}

func NewFloodAlertAppServiceClient(cc *grpc.ClientConn) FloodAlertAppServiceClient {
	return &floodAlertAppServiceClient{cc}
}

func (c *floodAlertAppServiceClient) CheckAlerts(ctx context.Context, in *CheckAlertsRequest, opts ...grpc.CallOption) (*CheckAlertsResponse, error) {
	out := new(CheckAlertsResponse)
	err := c.cc.Invoke(ctx, "/pb.FloodAlertAppService/checkAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floodAlertAppServiceClient) TestAlerts(ctx context.Context, in *TestAlertsRequest, opts ...grpc.CallOption) (*TestAlertsResponse, error) {
	out := new(TestAlertsResponse)
	err := c.cc.Invoke(ctx, "/pb.FloodAlertAppService/testAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floodAlertAppServiceClient) MuteAlerts(ctx context.Context, in *MuteAlertRequest, opts ...grpc.CallOption) (*MuteAlertResponse, error) {
	out := new(MuteAlertResponse)
	err := c.cc.Invoke(ctx, "/pb.FloodAlertAppService/muteAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floodAlertAppServiceClient) GetForcast(ctx context.Context, in *GetForcastRequest, opts ...grpc.CallOption) (*GetForcastResponse, error) {
	out := new(GetForcastResponse)
	err := c.cc.Invoke(ctx, "/pb.FloodAlertAppService/getForcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floodAlertAppServiceClient) UpdateForcast(ctx context.Context, in *UpdateForcastRequest, opts ...grpc.CallOption) (*UpdateForcastResponse, error) {
	out := new(UpdateForcastResponse)
	err := c.cc.Invoke(ctx, "/pb.FloodAlertAppService/updateForcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FloodAlertAppServiceServer is the server API for FloodAlertAppService service.
type FloodAlertAppServiceServer interface {
	CheckAlerts(context.Context, *CheckAlertsRequest) (*CheckAlertsResponse, error)
	TestAlerts(context.Context, *TestAlertsRequest) (*TestAlertsResponse, error)
	MuteAlerts(context.Context, *MuteAlertRequest) (*MuteAlertResponse, error)
	GetForcast(context.Context, *GetForcastRequest) (*GetForcastResponse, error)
	UpdateForcast(context.Context, *UpdateForcastRequest) (*UpdateForcastResponse, error)
}

func RegisterFloodAlertAppServiceServer(s *grpc.Server, srv FloodAlertAppServiceServer) {
	s.RegisterService(&_FloodAlertAppService_serviceDesc, srv)
}

func _FloodAlertAppService_CheckAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloodAlertAppServiceServer).CheckAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FloodAlertAppService/CheckAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloodAlertAppServiceServer).CheckAlerts(ctx, req.(*CheckAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloodAlertAppService_TestAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloodAlertAppServiceServer).TestAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FloodAlertAppService/TestAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloodAlertAppServiceServer).TestAlerts(ctx, req.(*TestAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloodAlertAppService_MuteAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MuteAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloodAlertAppServiceServer).MuteAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FloodAlertAppService/MuteAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloodAlertAppServiceServer).MuteAlerts(ctx, req.(*MuteAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloodAlertAppService_GetForcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloodAlertAppServiceServer).GetForcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FloodAlertAppService/GetForcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloodAlertAppServiceServer).GetForcast(ctx, req.(*GetForcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloodAlertAppService_UpdateForcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateForcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloodAlertAppServiceServer).UpdateForcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FloodAlertAppService/UpdateForcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloodAlertAppServiceServer).UpdateForcast(ctx, req.(*UpdateForcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FloodAlertAppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FloodAlertAppService",
	HandlerType: (*FloodAlertAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "checkAlerts",
			Handler:    _FloodAlertAppService_CheckAlerts_Handler,
		},
		{
			MethodName: "testAlerts",
			Handler:    _FloodAlertAppService_TestAlerts_Handler,
		},
		{
			MethodName: "muteAlerts",
			Handler:    _FloodAlertAppService_MuteAlerts_Handler,
		},
		{
			MethodName: "getForcast",
			Handler:    _FloodAlertAppService_GetForcast_Handler,
		},
		{
			MethodName: "updateForcast",
			Handler:    _FloodAlertAppService_UpdateForcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flood-app-api.proto",
}

func init() { proto.RegisterFile("flood-app-api.proto", fileDescriptor_flood_app_api_4106f79d85f73408) }

var fileDescriptor_flood_app_api_4106f79d85f73408 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x81, 0xdc, 0x70, 0x2f, 0x87, 0x70, 0xef, 0xe5, 0x00, 0xb5, 0x36, 0x2e, 0xc8, 0xc4,
	0x18, 0x12, 0xa5, 0x0b, 0x8c, 0x2b, 0x63, 0x94, 0x4d, 0x77, 0x6c, 0x8a, 0x2f, 0xd0, 0x3f, 0x07,
	0x24, 0xa2, 0x33, 0x76, 0xa6, 0x26, 0x3e, 0xab, 0x2f, 0x63, 0xfa, 0x97, 0x76, 0x4a, 0xc2, 0x72,
	0x7e, 0x73, 0xce, 0xf7, 0x9d, 0x9e, 0x6f, 0x0a, 0xa3, 0xcd, 0x9e, 0xf3, 0x70, 0xee, 0x09, 0x31,
	0xf7, 0xc4, 0xce, 0x16, 0x11, 0x57, 0x1c, 0x3b, 0xc2, 0x67, 0x0b, 0xc0, 0xe0, 0x85, 0x82, 0xd7,
	0xe5, 0x9e, 0x22, 0x25, 0x5d, 0xfa, 0x88, 0x49, 0x2a, 0xbc, 0x80, 0x9e, 0x97, 0x80, 0x15, 0x0f,
	0xc9, 0x6c, 0x4f, 0xdb, 0xb3, 0x5f, 0xee, 0x01, 0xb0, 0x39, 0x8c, 0x6a, 0x3d, 0x52, 0xf0, 0x77,
	0x49, 0x68, 0x40, 0x97, 0xc7, 0x4a, 0xc4, 0x2a, 0xed, 0xe8, 0xb9, 0xf9, 0x89, 0xdd, 0xc1, 0x50,
	0x91, 0x54, 0x75, 0x87, 0x29, 0xf4, 0x13, 0xb8, 0x22, 0x29, 0xbd, 0x2d, 0xe5, 0x1d, 0x55, 0xc4,
	0x6e, 0x00, 0xab, 0x6d, 0x27, 0x4c, 0x7c, 0xf8, 0xff, 0x16, 0x2b, 0x4a, 0xab, 0x0b, 0x8f, 0x4b,
	0x18, 0x94, 0xec, 0xf9, 0x4b, 0x14, 0x2e, 0x75, 0x88, 0x33, 0xf8, 0x57, 0x82, 0xb5, 0xf2, 0x54,
	0x2c, 0xcd, 0xce, 0xb4, 0x3d, 0xfb, 0xe3, 0xea, 0x98, 0x5d, 0xc3, 0xb0, 0xe2, 0x71, 0x62, 0xa0,
	0x11, 0x0c, 0xb7, 0xa4, 0x1c, 0x1e, 0x05, 0x9e, 0x2c, 0x26, 0x62, 0x36, 0x60, 0x15, 0xe6, 0x12,
	0x26, 0xfc, 0xde, 0x64, 0x28, 0xd7, 0x28, 0x8e, 0xcc, 0x80, 0x71, 0x2c, 0x42, 0x4f, 0x91, 0xa6,
	0xf3, 0x08, 0x13, 0x8d, 0xe7, 0x52, 0x57, 0xf0, 0x37, 0xbb, 0x08, 0x9d, 0x9a, 0xa2, 0x46, 0x17,
	0xdf, 0x1d, 0x18, 0x3b, 0xc9, 0x93, 0x48, 0x3f, 0x66, 0x29, 0xc4, 0x9a, 0xa2, 0xcf, 0x5d, 0x40,
	0xf8, 0x04, 0xfd, 0x4a, 0xb6, 0x68, 0xd8, 0xc2, 0xb7, 0x9b, 0x0f, 0xc4, 0x3a, 0x6b, 0xf0, 0x6c,
	0x00, 0xd6, 0xc2, 0x07, 0x80, 0x43, 0x6e, 0x38, 0x49, 0x0a, 0x1b, 0xf1, 0x5b, 0x86, 0x8e, 0xcb,
	0xf6, 0x7b, 0x80, 0x72, 0xc9, 0x12, 0xc7, 0x49, 0x9d, 0x1e, 0xac, 0x35, 0xd1, 0x68, 0xd5, 0xfb,
	0xb0, 0xdf, 0xcc, 0xbb, 0x11, 0x42, 0xe6, 0xdd, 0x8c, 0x81, 0xb5, 0xd0, 0x81, 0x41, 0x6d, 0xad,
	0x68, 0x26, 0xa5, 0xc7, 0x12, 0xb0, 0xce, 0x8f, 0xdc, 0x14, 0x3a, 0x7e, 0x37, 0xfd, 0xbf, 0x6e,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x27, 0xfc, 0x3b, 0xa9, 0x76, 0x03, 0x00, 0x00,
}
