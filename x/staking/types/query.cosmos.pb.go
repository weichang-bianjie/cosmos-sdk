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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Validators queries all validators that match the given status.
	Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error)
	// Validator queries validator info for given validator address.
	Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error)
	// ValidatorDelegations queries delegate info for given validator.
	ValidatorDelegations(ctx context.Context, in *QueryValidatorDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorDelegationsResponse, error)
	// ValidatorUnbondingDelegations queries unbonding delegations of a validator.
	ValidatorUnbondingDelegations(ctx context.Context, in *QueryValidatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorUnbondingDelegationsResponse, error)
	// Delegation queries delegate info for given validator delegator pair.
	Delegation(ctx context.Context, in *QueryDelegationRequest, opts ...grpc.CallOption) (*QueryDelegationResponse, error)
	// UnbondingDelegation queries unbonding info for given validator delegator
	// pair.
	UnbondingDelegation(ctx context.Context, in *QueryUnbondingDelegationRequest, opts ...grpc.CallOption) (*QueryUnbondingDelegationResponse, error)
	// DelegatorDelegations queries all delegations of a given delegator address.
	DelegatorDelegations(ctx context.Context, in *QueryDelegatorDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorDelegationsResponse, error)
	// DelegatorUnbondingDelegations queries all unbonding delegations of a given
	// delegator address.
	DelegatorUnbondingDelegations(ctx context.Context, in *QueryDelegatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorUnbondingDelegationsResponse, error)
	// Redelegations queries redelegations of given address.
	Redelegations(ctx context.Context, in *QueryRedelegationsRequest, opts ...grpc.CallOption) (*QueryRedelegationsResponse, error)
	// DelegatorValidators queries all validators info for given delegator
	// address.
	DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorValidator queries validator info for given delegator validator
	// pair.
	DelegatorValidator(ctx context.Context, in *QueryDelegatorValidatorRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorResponse, error)
	// HistoricalInfo queries the historical info for given height.
	HistoricalInfo(ctx context.Context, in *QueryHistoricalInfoRequest, opts ...grpc.CallOption) (*QueryHistoricalInfoResponse, error)
	// Pool queries the pool info.
	Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// Parameters queries the staking parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc                             grpc.ClientConnInterface
	_Validators                    types.Invoker
	_Validator                     types.Invoker
	_ValidatorDelegations          types.Invoker
	_ValidatorUnbondingDelegations types.Invoker
	_Delegation                    types.Invoker
	_UnbondingDelegation           types.Invoker
	_DelegatorDelegations          types.Invoker
	_DelegatorUnbondingDelegations types.Invoker
	_Redelegations                 types.Invoker
	_DelegatorValidators           types.Invoker
	_DelegatorValidator            types.Invoker
	_HistoricalInfo                types.Invoker
	_Pool                          types.Invoker
	_Params                        types.Invoker
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc: cc}
}

func (c *queryClient) Validators(ctx context.Context, in *QueryValidatorsRequest, opts ...grpc.CallOption) (*QueryValidatorsResponse, error) {
	if invoker := c._Validators; invoker != nil {
		var out QueryValidatorsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Validators, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/Validators")
		if err != nil {
			var out QueryValidatorsResponse
			err = c._Validators(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryValidatorsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Validators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Validator(ctx context.Context, in *QueryValidatorRequest, opts ...grpc.CallOption) (*QueryValidatorResponse, error) {
	if invoker := c._Validator; invoker != nil {
		var out QueryValidatorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Validator, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/Validator")
		if err != nil {
			var out QueryValidatorResponse
			err = c._Validator(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryValidatorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Validator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorDelegations(ctx context.Context, in *QueryValidatorDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorDelegationsResponse, error) {
	if invoker := c._ValidatorDelegations; invoker != nil {
		var out QueryValidatorDelegationsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._ValidatorDelegations, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/ValidatorDelegations")
		if err != nil {
			var out QueryValidatorDelegationsResponse
			err = c._ValidatorDelegations(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryValidatorDelegationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/ValidatorDelegations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorUnbondingDelegations(ctx context.Context, in *QueryValidatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryValidatorUnbondingDelegationsResponse, error) {
	if invoker := c._ValidatorUnbondingDelegations; invoker != nil {
		var out QueryValidatorUnbondingDelegationsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._ValidatorUnbondingDelegations, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/ValidatorUnbondingDelegations")
		if err != nil {
			var out QueryValidatorUnbondingDelegationsResponse
			err = c._ValidatorUnbondingDelegations(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryValidatorUnbondingDelegationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/ValidatorUnbondingDelegations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Delegation(ctx context.Context, in *QueryDelegationRequest, opts ...grpc.CallOption) (*QueryDelegationResponse, error) {
	if invoker := c._Delegation; invoker != nil {
		var out QueryDelegationResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Delegation, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/Delegation")
		if err != nil {
			var out QueryDelegationResponse
			err = c._Delegation(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryDelegationResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Delegation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UnbondingDelegation(ctx context.Context, in *QueryUnbondingDelegationRequest, opts ...grpc.CallOption) (*QueryUnbondingDelegationResponse, error) {
	if invoker := c._UnbondingDelegation; invoker != nil {
		var out QueryUnbondingDelegationResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._UnbondingDelegation, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/UnbondingDelegation")
		if err != nil {
			var out QueryUnbondingDelegationResponse
			err = c._UnbondingDelegation(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryUnbondingDelegationResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/UnbondingDelegation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorDelegations(ctx context.Context, in *QueryDelegatorDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorDelegationsResponse, error) {
	if invoker := c._DelegatorDelegations; invoker != nil {
		var out QueryDelegatorDelegationsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._DelegatorDelegations, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/DelegatorDelegations")
		if err != nil {
			var out QueryDelegatorDelegationsResponse
			err = c._DelegatorDelegations(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryDelegatorDelegationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/DelegatorDelegations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorUnbondingDelegations(ctx context.Context, in *QueryDelegatorUnbondingDelegationsRequest, opts ...grpc.CallOption) (*QueryDelegatorUnbondingDelegationsResponse, error) {
	if invoker := c._DelegatorUnbondingDelegations; invoker != nil {
		var out QueryDelegatorUnbondingDelegationsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._DelegatorUnbondingDelegations, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/DelegatorUnbondingDelegations")
		if err != nil {
			var out QueryDelegatorUnbondingDelegationsResponse
			err = c._DelegatorUnbondingDelegations(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryDelegatorUnbondingDelegationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/DelegatorUnbondingDelegations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Redelegations(ctx context.Context, in *QueryRedelegationsRequest, opts ...grpc.CallOption) (*QueryRedelegationsResponse, error) {
	if invoker := c._Redelegations; invoker != nil {
		var out QueryRedelegationsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Redelegations, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/Redelegations")
		if err != nil {
			var out QueryRedelegationsResponse
			err = c._Redelegations(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryRedelegationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Redelegations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorValidators(ctx context.Context, in *QueryDelegatorValidatorsRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorsResponse, error) {
	if invoker := c._DelegatorValidators; invoker != nil {
		var out QueryDelegatorValidatorsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._DelegatorValidators, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/DelegatorValidators")
		if err != nil {
			var out QueryDelegatorValidatorsResponse
			err = c._DelegatorValidators(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryDelegatorValidatorsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/DelegatorValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegatorValidator(ctx context.Context, in *QueryDelegatorValidatorRequest, opts ...grpc.CallOption) (*QueryDelegatorValidatorResponse, error) {
	if invoker := c._DelegatorValidator; invoker != nil {
		var out QueryDelegatorValidatorResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._DelegatorValidator, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/DelegatorValidator")
		if err != nil {
			var out QueryDelegatorValidatorResponse
			err = c._DelegatorValidator(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryDelegatorValidatorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/DelegatorValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HistoricalInfo(ctx context.Context, in *QueryHistoricalInfoRequest, opts ...grpc.CallOption) (*QueryHistoricalInfoResponse, error) {
	if invoker := c._HistoricalInfo; invoker != nil {
		var out QueryHistoricalInfoResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._HistoricalInfo, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/HistoricalInfo")
		if err != nil {
			var out QueryHistoricalInfoResponse
			err = c._HistoricalInfo(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryHistoricalInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/HistoricalInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	if invoker := c._Pool; invoker != nil {
		var out QueryPoolResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Pool, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/Pool")
		if err != nil {
			var out QueryPoolResponse
			err = c._Pool(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Pool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	if invoker := c._Params; invoker != nil {
		var out QueryParamsResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Params, err = invokerConn.Invoker("/cosmos.staking.v1beta1.Query/Params")
		if err != nil {
			var out QueryParamsResponse
			err = c._Params(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.staking.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Validators queries all validators that match the given status.
	Validators(context.Context, *QueryValidatorsRequest) (*QueryValidatorsResponse, error)
	// Validator queries validator info for given validator address.
	Validator(context.Context, *QueryValidatorRequest) (*QueryValidatorResponse, error)
	// ValidatorDelegations queries delegate info for given validator.
	ValidatorDelegations(context.Context, *QueryValidatorDelegationsRequest) (*QueryValidatorDelegationsResponse, error)
	// ValidatorUnbondingDelegations queries unbonding delegations of a validator.
	ValidatorUnbondingDelegations(context.Context, *QueryValidatorUnbondingDelegationsRequest) (*QueryValidatorUnbondingDelegationsResponse, error)
	// Delegation queries delegate info for given validator delegator pair.
	Delegation(context.Context, *QueryDelegationRequest) (*QueryDelegationResponse, error)
	// UnbondingDelegation queries unbonding info for given validator delegator
	// pair.
	UnbondingDelegation(context.Context, *QueryUnbondingDelegationRequest) (*QueryUnbondingDelegationResponse, error)
	// DelegatorDelegations queries all delegations of a given delegator address.
	DelegatorDelegations(context.Context, *QueryDelegatorDelegationsRequest) (*QueryDelegatorDelegationsResponse, error)
	// DelegatorUnbondingDelegations queries all unbonding delegations of a given
	// delegator address.
	DelegatorUnbondingDelegations(context.Context, *QueryDelegatorUnbondingDelegationsRequest) (*QueryDelegatorUnbondingDelegationsResponse, error)
	// Redelegations queries redelegations of given address.
	Redelegations(context.Context, *QueryRedelegationsRequest) (*QueryRedelegationsResponse, error)
	// DelegatorValidators queries all validators info for given delegator
	// address.
	DelegatorValidators(context.Context, *QueryDelegatorValidatorsRequest) (*QueryDelegatorValidatorsResponse, error)
	// DelegatorValidator queries validator info for given delegator validator
	// pair.
	DelegatorValidator(context.Context, *QueryDelegatorValidatorRequest) (*QueryDelegatorValidatorResponse, error)
	// HistoricalInfo queries the historical info for given height.
	HistoricalInfo(context.Context, *QueryHistoricalInfoRequest) (*QueryHistoricalInfoResponse, error)
	// Pool queries the pool info.
	Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error)
	// Parameters queries the staking parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Validators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validators(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validators(types.UnwrapSDKContext(ctx), req.(*QueryValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Validator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validator(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validator(types.UnwrapSDKContext(ctx), req.(*QueryValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorDelegations(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryValidatorDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorDelegations(types.UnwrapSDKContext(ctx), req.(*QueryValidatorDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorUnbondingDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorUnbondingDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorUnbondingDelegations(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryValidatorUnbondingDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorUnbondingDelegations(types.UnwrapSDKContext(ctx), req.(*QueryValidatorUnbondingDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Delegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Delegation(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryDelegation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Delegation(types.UnwrapSDKContext(ctx), req.(*QueryDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UnbondingDelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUnbondingDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UnbondingDelegation(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryUnbondingDelegation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UnbondingDelegation(types.UnwrapSDKContext(ctx), req.(*QueryUnbondingDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorDelegations(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryDelegatorDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorDelegations(types.UnwrapSDKContext(ctx), req.(*QueryDelegatorDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorUnbondingDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorUnbondingDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorUnbondingDelegations(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorUnbondingDelegations(types.UnwrapSDKContext(ctx), req.(*QueryDelegatorUnbondingDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Redelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRedelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Redelegations(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryRedelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Redelegations(types.UnwrapSDKContext(ctx), req.(*QueryRedelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidators(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryDelegatorValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidators(types.UnwrapSDKContext(ctx), req.(*QueryDelegatorValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegatorValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidator(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryDelegatorValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidator(types.UnwrapSDKContext(ctx), req.(*QueryDelegatorValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HistoricalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHistoricalInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HistoricalInfo(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryHistoricalInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HistoricalInfo(types.UnwrapSDKContext(ctx), req.(*QueryHistoricalInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(types.UnwrapSDKContext(ctx), req.(*QueryPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.QueryParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(types.UnwrapSDKContext(ctx), req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.staking.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validators",
			Handler:    _Query_Validators_Handler,
		},
		{
			MethodName: "Validator",
			Handler:    _Query_Validator_Handler,
		},
		{
			MethodName: "ValidatorDelegations",
			Handler:    _Query_ValidatorDelegations_Handler,
		},
		{
			MethodName: "ValidatorUnbondingDelegations",
			Handler:    _Query_ValidatorUnbondingDelegations_Handler,
		},
		{
			MethodName: "Delegation",
			Handler:    _Query_Delegation_Handler,
		},
		{
			MethodName: "UnbondingDelegation",
			Handler:    _Query_UnbondingDelegation_Handler,
		},
		{
			MethodName: "DelegatorDelegations",
			Handler:    _Query_DelegatorDelegations_Handler,
		},
		{
			MethodName: "DelegatorUnbondingDelegations",
			Handler:    _Query_DelegatorUnbondingDelegations_Handler,
		},
		{
			MethodName: "Redelegations",
			Handler:    _Query_Redelegations_Handler,
		},
		{
			MethodName: "DelegatorValidators",
			Handler:    _Query_DelegatorValidators_Handler,
		},
		{
			MethodName: "DelegatorValidator",
			Handler:    _Query_DelegatorValidator_Handler,
		},
		{
			MethodName: "HistoricalInfo",
			Handler:    _Query_HistoricalInfo_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Metadata: "cosmos/staking/v1beta1/query.proto",
}

const (
	QueryValidatorsMethod                    = "/cosmos.staking.v1beta1.Query/Validators"
	QueryValidatorMethod                     = "/cosmos.staking.v1beta1.Query/Validator"
	QueryValidatorDelegationsMethod          = "/cosmos.staking.v1beta1.Query/ValidatorDelegations"
	QueryValidatorUnbondingDelegationsMethod = "/cosmos.staking.v1beta1.Query/ValidatorUnbondingDelegations"
	QueryDelegationMethod                    = "/cosmos.staking.v1beta1.Query/Delegation"
	QueryUnbondingDelegationMethod           = "/cosmos.staking.v1beta1.Query/UnbondingDelegation"
	QueryDelegatorDelegationsMethod          = "/cosmos.staking.v1beta1.Query/DelegatorDelegations"
	QueryDelegatorUnbondingDelegationsMethod = "/cosmos.staking.v1beta1.Query/DelegatorUnbondingDelegations"
	QueryRedelegationsMethod                 = "/cosmos.staking.v1beta1.Query/Redelegations"
	QueryDelegatorValidatorsMethod           = "/cosmos.staking.v1beta1.Query/DelegatorValidators"
	QueryDelegatorValidatorMethod            = "/cosmos.staking.v1beta1.Query/DelegatorValidator"
	QueryHistoricalInfoMethod                = "/cosmos.staking.v1beta1.Query/HistoricalInfo"
	QueryPoolMethod                          = "/cosmos.staking.v1beta1.Query/Pool"
	QueryParamsMethod                        = "/cosmos.staking.v1beta1.Query/Params"
)
