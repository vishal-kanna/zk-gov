// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sdk/zkgov/v1beta1/tx.proto

package zkgovv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_RegisterUser_FullMethodName    = "/sdk.zkgov.v1beta1.Msg/RegisterUser"
	Msg_VoteProposal_FullMethodName    = "/sdk.zkgov.v1beta1.Msg/VoteProposal"
	Msg_CreateProposal_FullMethodName  = "/sdk.zkgov.v1beta1.Msg/CreateProposal"
	Msg_ProcessProposal_FullMethodName = "/sdk.zkgov.v1beta1.Msg/ProcessProposal"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterUser
	RegisterUser(ctx context.Context, in *MsgRegisterUser, opts ...grpc.CallOption) (*MsgRegisterUserResponse, error)
	// VoteProposal
	VoteProposal(ctx context.Context, in *MsgVoteProposal, opts ...grpc.CallOption) (*MsgVoteProposalResponse, error)
	// CreateProposal
	CreateProposal(ctx context.Context, in *MsgCreateProposal, opts ...grpc.CallOption) (*MsgCreateProposalResponse, error)
	// ProcessProposal
	ProcessProposal(ctx context.Context, in *MsgProcessProposal, opts ...grpc.CallOption) (*MsgProcessProposalResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterUser(ctx context.Context, in *MsgRegisterUser, opts ...grpc.CallOption) (*MsgRegisterUserResponse, error) {
	out := new(MsgRegisterUserResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VoteProposal(ctx context.Context, in *MsgVoteProposal, opts ...grpc.CallOption) (*MsgVoteProposalResponse, error) {
	out := new(MsgVoteProposalResponse)
	err := c.cc.Invoke(ctx, Msg_VoteProposal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateProposal(ctx context.Context, in *MsgCreateProposal, opts ...grpc.CallOption) (*MsgCreateProposalResponse, error) {
	out := new(MsgCreateProposalResponse)
	err := c.cc.Invoke(ctx, Msg_CreateProposal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ProcessProposal(ctx context.Context, in *MsgProcessProposal, opts ...grpc.CallOption) (*MsgProcessProposalResponse, error) {
	out := new(MsgProcessProposalResponse)
	err := c.cc.Invoke(ctx, Msg_ProcessProposal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// RegisterUser
	RegisterUser(context.Context, *MsgRegisterUser) (*MsgRegisterUserResponse, error)
	// VoteProposal
	VoteProposal(context.Context, *MsgVoteProposal) (*MsgVoteProposalResponse, error)
	// CreateProposal
	CreateProposal(context.Context, *MsgCreateProposal) (*MsgCreateProposalResponse, error)
	// ProcessProposal
	ProcessProposal(context.Context, *MsgProcessProposal) (*MsgProcessProposalResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RegisterUser(context.Context, *MsgRegisterUser) (*MsgRegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedMsgServer) VoteProposal(context.Context, *MsgVoteProposal) (*MsgVoteProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteProposal not implemented")
}
func (UnimplementedMsgServer) CreateProposal(context.Context, *MsgCreateProposal) (*MsgCreateProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProposal not implemented")
}
func (UnimplementedMsgServer) ProcessProposal(context.Context, *MsgProcessProposal) (*MsgProcessProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessProposal not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterUser(ctx, req.(*MsgRegisterUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VoteProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_VoteProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteProposal(ctx, req.(*MsgVoteProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProposal(ctx, req.(*MsgCreateProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProcessProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ProcessProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ProcessProposal(ctx, req.(*MsgProcessProposal))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sdk.zkgov.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _Msg_RegisterUser_Handler,
		},
		{
			MethodName: "VoteProposal",
			Handler:    _Msg_VoteProposal_Handler,
		},
		{
			MethodName: "CreateProposal",
			Handler:    _Msg_CreateProposal_Handler,
		},
		{
			MethodName: "ProcessProposal",
			Handler:    _Msg_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdk/zkgov/v1beta1/tx.proto",
}
