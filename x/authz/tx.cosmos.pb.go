// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authz

import (
	context "context"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// Grant grants the provided authorization to the grantee on the granter's
	// account with the provided expiration time.
	Grant(ctx context.Context, in *MsgGrant, opts ...grpc.CallOption) (*MsgGrantResponse, error)
	// Exec attempts to execute the provided messages using
	// authorizations granted to the grantee. Each message should have only
	// one signer corresponding to the granter of the authorization.
	Exec(ctx context.Context, in *MsgExec, opts ...grpc.CallOption) (*MsgExecResponse, error)
	// Revoke revokes any authorization corresponding to the provided method name on the
	// granter's account that has been granted to the grantee.
	Revoke(ctx context.Context, in *MsgRevoke, opts ...grpc.CallOption) (*MsgRevokeResponse, error)
}

type msgClient struct {
	cc      grpc.ClientConnInterface
	_Grant  types.Invoker
	_Exec   types.Invoker
	_Revoke types.Invoker
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc: cc}
}

func (c *msgClient) Grant(ctx context.Context, in *MsgGrant, opts ...grpc.CallOption) (*MsgGrantResponse, error) {
	if invoker := c._Grant; invoker != nil {
		var out MsgGrantResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Grant, err = invokerConn.Invoker("/cosmos.authz.v1beta1.Msg/Grant")
		if err != nil {
			var out MsgGrantResponse
			err = c._Grant(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgGrantResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Msg/Grant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Exec(ctx context.Context, in *MsgExec, opts ...grpc.CallOption) (*MsgExecResponse, error) {
	if invoker := c._Exec; invoker != nil {
		var out MsgExecResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Exec, err = invokerConn.Invoker("/cosmos.authz.v1beta1.Msg/Exec")
		if err != nil {
			var out MsgExecResponse
			err = c._Exec(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgExecResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Msg/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Revoke(ctx context.Context, in *MsgRevoke, opts ...grpc.CallOption) (*MsgRevokeResponse, error) {
	if invoker := c._Revoke; invoker != nil {
		var out MsgRevokeResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Revoke, err = invokerConn.Invoker("/cosmos.authz.v1beta1.Msg/Revoke")
		if err != nil {
			var out MsgRevokeResponse
			err = c._Revoke(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgRevokeResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Msg/Revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Grant grants the provided authorization to the grantee on the granter's
	// account with the provided expiration time.
	Grant(context.Context, *MsgGrant) (*MsgGrantResponse, error)
	// Exec attempts to execute the provided messages using
	// authorizations granted to the grantee. Each message should have only
	// one signer corresponding to the granter of the authorization.
	Exec(context.Context, *MsgExec) (*MsgExecResponse, error)
	// Revoke revokes any authorization corresponding to the provided method name on the
	// granter's account that has been granted to the grantee.
	Revoke(context.Context, *MsgRevoke) (*MsgRevokeResponse, error)
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_Grant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGrant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Grant(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.MsgGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Grant(types.UnwrapSDKContext(ctx), req.(*MsgGrant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Exec(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.MsgExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Exec(types.UnwrapSDKContext(ctx), req.(*MsgExec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevoke)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Revoke(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.MsgRevoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Revoke(types.UnwrapSDKContext(ctx), req.(*MsgRevoke))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grant",
			Handler:    _Msg_Grant_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Msg_Exec_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _Msg_Revoke_Handler,
		},
	},
	Metadata: "cosmos/authz/v1beta1/tx.proto",
}

const (
	MsgGrantMethod  = "/cosmos.authz.v1beta1.Msg/Grant"
	MsgExecMethod   = "/cosmos.authz.v1beta1.Msg/Exec"
	MsgRevokeMethod = "/cosmos.authz.v1beta1.Msg/Revoke"
)
