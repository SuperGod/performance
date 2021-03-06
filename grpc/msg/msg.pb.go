// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package msg

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

type MsgReq struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgReq) Reset()         { *m = MsgReq{} }
func (m *MsgReq) String() string { return proto.CompactTextString(m) }
func (*MsgReq) ProtoMessage()    {}
func (*MsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e2bc963fafc46e94, []int{0}
}
func (m *MsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgReq.Unmarshal(m, b)
}
func (m *MsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgReq.Marshal(b, m, deterministic)
}
func (dst *MsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReq.Merge(dst, src)
}
func (m *MsgReq) XXX_Size() int {
	return xxx_messageInfo_MsgReq.Size(m)
}
func (m *MsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReq proto.InternalMessageInfo

func (m *MsgReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MsgReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgReply) Reset()         { *m = MsgReply{} }
func (m *MsgReply) String() string { return proto.CompactTextString(m) }
func (*MsgReply) ProtoMessage()    {}
func (*MsgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_e2bc963fafc46e94, []int{1}
}
func (m *MsgReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgReply.Unmarshal(m, b)
}
func (m *MsgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgReply.Marshal(b, m, deterministic)
}
func (dst *MsgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReply.Merge(dst, src)
}
func (m *MsgReply) XXX_Size() int {
	return xxx_messageInfo_MsgReply.Size(m)
}
func (m *MsgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReply.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgReq)(nil), "MsgReq")
	proto.RegisterType((*MsgReply)(nil), "MsgReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// Sends a greeting
	Send(ctx context.Context, in *MsgReq, opts ...grpc.CallOption) (*MsgReply, error)
}

type msgServiceClient struct {
	cc *grpc.ClientConn
}

func NewMsgServiceClient(cc *grpc.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Send(ctx context.Context, in *MsgReq, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/MsgService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// Sends a greeting
	Send(context.Context, *MsgReq) (*MsgReply, error)
}

func RegisterMsgServiceServer(s *grpc.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MsgService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Send(ctx, req.(*MsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _MsgService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_msg_e2bc963fafc46e94) }

var fileDescriptor_msg_e2bc963fafc46e94 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe1, 0x62, 0xf3, 0x2d, 0x4e, 0x0f, 0x4a, 0x2d, 0x14,
	0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3,
	0x95, 0xb8, 0xb8, 0x38, 0xc0, 0xb2, 0x05, 0x39, 0x95, 0x46, 0x5a, 0x5c, 0x5c, 0xbe, 0xc5, 0xe9,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x32, 0x5c, 0x2c, 0xc1, 0xa9, 0x79, 0x29, 0x42,
	0xec, 0x7a, 0x10, 0xed, 0x52, 0x9c, 0x7a, 0x30, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xc3, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x88, 0x4c, 0xc3, 0x69, 0x00, 0x00, 0x00,
}
