// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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

// The request message containing the user's name.
type MessageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_22596f7009dd4efa, []int{0}
}
func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (dst *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(dst, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type MessageReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReply) Reset()         { *m = MessageReply{} }
func (m *MessageReply) String() string { return proto.CompactTextString(m) }
func (*MessageReply) ProtoMessage()    {}
func (*MessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_22596f7009dd4efa, []int{1}
}
func (m *MessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageReply.Unmarshal(m, b)
}
func (m *MessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageReply.Marshal(b, m, deterministic)
}
func (dst *MessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReply.Merge(dst, src)
}
func (m *MessageReply) XXX_Size() int {
	return xxx_messageInfo_MessageReply.Size(m)
}
func (m *MessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReply proto.InternalMessageInfo

func (m *MessageReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "message.MessageRequest")
	proto.RegisterType((*MessageReply)(nil), "message.MessageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	Ping(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Ping(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/message.Greeter/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	Ping(context.Context, *MessageRequest) (*MessageReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Greeter/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Ping(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Greeter_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_22596f7009dd4efa) }

var fileDescriptor_message_22596f7009dd4efa = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x54, 0xb8,
	0xf8, 0x7c, 0x21, 0xcc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x83, 0x8b, 0x07,
	0xae, 0xaa, 0x20, 0xa7, 0x52, 0x48, 0x82, 0x0b, 0x66, 0x00, 0x54, 0x19, 0x8c, 0x6b, 0xe4, 0xcc,
	0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0x64, 0xc1, 0xc5, 0x12, 0x90, 0x99, 0x97,
	0x2e, 0x24, 0xae, 0x07, 0xb3, 0x1b, 0xd5, 0x26, 0x29, 0x51, 0x4c, 0x89, 0x82, 0x9c, 0x4a, 0x25,
	0x06, 0x27, 0x2d, 0x2e, 0x89, 0xcc, 0x7c, 0xbd, 0xf4, 0xa2, 0x82, 0x64, 0xbd, 0xd4, 0x8a, 0xc4,
	0xdc, 0x82, 0x9c, 0xd4, 0x62, 0x98, 0x52, 0x27, 0x98, 0x43, 0x02, 0x40, 0xfe, 0x08, 0x60, 0x4c,
	0x62, 0x03, 0x7b, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xc6, 0x39, 0x07, 0xe1, 0x00,
	0x00, 0x00,
}
