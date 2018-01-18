// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_pdk.proto

package ddproto

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

// Ignoring public import of pdk_req_createDesk from pdk_hall.proto

// Ignoring public import of pdk_ack_createDesk from pdk_hall.proto

// Ignoring public import of pdk_req_gameRecord from pdk_hall.proto

// Ignoring public import of pdk_base_userRecord from pdk_hall.proto

// Ignoring public import of pdk_base_gameRecord from pdk_hall.proto

// Ignoring public import of pdk_ack_gameRecord from pdk_hall.proto

// Ignoring public import of pdk_req_enterDesk from pdk_hall.proto

// Ignoring public import of pdk_ack_enterDesk from pdk_hall.proto

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PdkRpc service

type PdkRpcClient interface {
	// 创建房间
	CreateRoom(ctx context.Context, in *PdkReqCreateDesk, opts ...grpc.CallOption) (*PdkAckCreateDesk, error)
}

type pdkRpcClient struct {
	cc *grpc.ClientConn
}

func NewPdkRpcClient(cc *grpc.ClientConn) PdkRpcClient {
	return &pdkRpcClient{cc}
}

func (c *pdkRpcClient) CreateRoom(ctx context.Context, in *PdkReqCreateDesk, opts ...grpc.CallOption) (*PdkAckCreateDesk, error) {
	out := new(PdkAckCreateDesk)
	err := grpc.Invoke(ctx, "/ddproto.PdkRpc/CreateRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PdkRpc service

type PdkRpcServer interface {
	// 创建房间
	CreateRoom(context.Context, *PdkReqCreateDesk) (*PdkAckCreateDesk, error)
}

func RegisterPdkRpcServer(s *grpc.Server, srv PdkRpcServer) {
	s.RegisterService(&_PdkRpc_serviceDesc, srv)
}

func _PdkRpc_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PdkReqCreateDesk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PdkRpcServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ddproto.PdkRpc/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PdkRpcServer).CreateRoom(ctx, req.(*PdkReqCreateDesk))
	}
	return interceptor(ctx, in, info, handler)
}

var _PdkRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ddproto.PdkRpc",
	HandlerType: (*PdkRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _PdkRpc_CreateRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_pdk.proto",
}

func init() { proto.RegisterFile("rpc_pdk.proto", fileDescriptor55) }

var fileDescriptor55 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2a, 0x48, 0x8e,
	0x2f, 0x48, 0xc9, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0xa4,
	0xf8, 0x0a, 0x52, 0xb2, 0xe3, 0x33, 0x12, 0x73, 0x72, 0x20, 0x12, 0x46, 0x41, 0x5c, 0x6c, 0x01,
	0x29, 0xd9, 0x41, 0x05, 0xc9, 0x42, 0x1e, 0x5c, 0x5c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x41,
	0xf9, 0xf9, 0xb9, 0x42, 0xd2, 0x7a, 0x50, 0x1d, 0x7a, 0x20, 0x0d, 0x45, 0xa9, 0x85, 0xf1, 0xc9,
	0x60, 0x49, 0x97, 0xd4, 0xe2, 0x6c, 0x29, 0x54, 0xc9, 0xc4, 0xe4, 0x6c, 0x24, 0x49, 0x25, 0x06,
	0x27, 0x26, 0x0f, 0xe6, 0x00, 0x06, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x87, 0xda, 0x95,
	0x82, 0x00, 0x00, 0x00,
}
