// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

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
	// CreateValidator defines a method for creating a new validator.
	CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error)
	// EditValidator defines a method for editing an existing validator.
	EditValidator(ctx context.Context, in *MsgEditValidator, opts ...grpc.CallOption) (*MsgEditValidatorResponse, error)
	// Delegate defines a method for performing a delegation of coins
	// from a delegator to a validator.
	Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error)
	// BeginRedelegate defines a method for performing a redelegation
	// of coins from a delegator and source validator to a destination validator.
	BeginRedelegate(ctx context.Context, in *MsgBeginRedelegate, opts ...grpc.CallOption) (*MsgBeginRedelegateResponse, error)
	// Undelegate defines a method for performing an undelegation from a
	// delegate and a validator.
	Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error)
}

type msgClient struct {
	cc               grpc.ClientConnInterface
	_CreateValidator types.Invoker
	_EditValidator   types.Invoker
	_Delegate        types.Invoker
	_BeginRedelegate types.Invoker
	_Undelegate      types.Invoker
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc: cc}
}

func (c *msgClient) CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error) {
	if invoker := c._CreateValidator; invoker != nil {
		var out MsgCreateValidatorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._CreateValidator, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Msg/CreateValidator")
		if err != nil {
			var out MsgCreateValidatorResponse
			err = c._CreateValidator(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgCreateValidatorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Msg/CreateValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditValidator(ctx context.Context, in *MsgEditValidator, opts ...grpc.CallOption) (*MsgEditValidatorResponse, error) {
	if invoker := c._EditValidator; invoker != nil {
		var out MsgEditValidatorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._EditValidator, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Msg/EditValidator")
		if err != nil {
			var out MsgEditValidatorResponse
			err = c._EditValidator(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgEditValidatorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Msg/EditValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error) {
	if invoker := c._Delegate; invoker != nil {
		var out MsgDelegateResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Delegate, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Msg/Delegate")
		if err != nil {
			var out MsgDelegateResponse
			err = c._Delegate(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgDelegateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Msg/Delegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BeginRedelegate(ctx context.Context, in *MsgBeginRedelegate, opts ...grpc.CallOption) (*MsgBeginRedelegateResponse, error) {
	if invoker := c._BeginRedelegate; invoker != nil {
		var out MsgBeginRedelegateResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._BeginRedelegate, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Msg/BeginRedelegate")
		if err != nil {
			var out MsgBeginRedelegateResponse
			err = c._BeginRedelegate(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgBeginRedelegateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Msg/BeginRedelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error) {
	if invoker := c._Undelegate; invoker != nil {
		var out MsgUndelegateResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Undelegate, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Msg/Undelegate")
		if err != nil {
			var out MsgUndelegateResponse
			err = c._Undelegate(ctx, in, &out)
			return &out, err
		}
	}
	out := new(MsgUndelegateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Msg/Undelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateValidator defines a method for creating a new validator.
	CreateValidator(context.Context, *MsgCreateValidator) (*MsgCreateValidatorResponse, error)
	// EditValidator defines a method for editing an existing validator.
	EditValidator(context.Context, *MsgEditValidator) (*MsgEditValidatorResponse, error)
	// Delegate defines a method for performing a delegation of coins
	// from a delegator to a validator.
	Delegate(context.Context, *MsgDelegate) (*MsgDelegateResponse, error)
	// BeginRedelegate defines a method for performing a redelegation
	// of coins from a delegator and source validator to a destination validator.
	BeginRedelegate(context.Context, *MsgBeginRedelegate) (*MsgBeginRedelegateResponse, error)
	// Undelegate defines a method for performing an undelegation from a
	// delegate and a validator.
	Undelegate(context.Context, *MsgUndelegate) (*MsgUndelegateResponse, error)
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateValidator(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.MsgCreateValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateValidator(types.UnwrapSDKContext(ctx), req.(*MsgCreateValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditValidator(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.MsgEditValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditValidator(types.UnwrapSDKContext(ctx), req.(*MsgEditValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Delegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Delegate(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.MsgDelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Delegate(types.UnwrapSDKContext(ctx), req.(*MsgDelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BeginRedelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBeginRedelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BeginRedelegate(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.MsgBeginRedelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BeginRedelegate(types.UnwrapSDKContext(ctx), req.(*MsgBeginRedelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Undelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUndelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Undelegate(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.MsgUndelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Undelegate(types.UnwrapSDKContext(ctx), req.(*MsgUndelegate))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.staking.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateValidator",
			Handler:    _Msg_CreateValidator_Handler,
		},
		{
			MethodName: "EditValidator",
			Handler:    _Msg_EditValidator_Handler,
		},
		{
			MethodName: "Delegate",
			Handler:    _Msg_Delegate_Handler,
		},
		{
			MethodName: "BeginRedelegate",
			Handler:    _Msg_BeginRedelegate_Handler,
		},
		{
			MethodName: "Undelegate",
			Handler:    _Msg_Undelegate_Handler,
		},
	},
	Metadata: "cosmos/staking/v1beta1/tx.proto",
}

const (
	MsgCreateValidatorMethod = "/cosmos.staking.v1beta1.Msg/CreateValidator"
	MsgEditValidatorMethod   = "/cosmos.staking.v1beta1.Msg/EditValidator"
	MsgDelegateMethod        = "/cosmos.staking.v1beta1.Msg/Delegate"
	MsgBeginRedelegateMethod = "/cosmos.staking.v1beta1.Msg/BeginRedelegate"
	MsgUndelegateMethod      = "/cosmos.staking.v1beta1.Msg/Undelegate"
)
