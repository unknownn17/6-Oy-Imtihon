// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: budget.proto

package budgetproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BudgetService_CreateBudget_FullMethodName = "/BudgetService/CreateBudget"
	BudgetService_GetBudgets_FullMethodName   = "/BudgetService/GetBudgets"
	BudgetService_UpdateBudget_FullMethodName = "/BudgetService/UpdateBudget"
)

// BudgetServiceClient is the client API for BudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BudgetServiceClient interface {
	CreateBudget(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*CreateBudgetResponse, error)
	GetBudgets(ctx context.Context, in *GetBudgetsRequest, opts ...grpc.CallOption) (*GetBudgetsResponse, error)
	UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*UpdateBudgetResponse, error)
}

type budgetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBudgetServiceClient(cc grpc.ClientConnInterface) BudgetServiceClient {
	return &budgetServiceClient{cc}
}

func (c *budgetServiceClient) CreateBudget(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*CreateBudgetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBudgetResponse)
	err := c.cc.Invoke(ctx, BudgetService_CreateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) GetBudgets(ctx context.Context, in *GetBudgetsRequest, opts ...grpc.CallOption) (*GetBudgetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBudgetsResponse)
	err := c.cc.Invoke(ctx, BudgetService_GetBudgets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*UpdateBudgetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBudgetResponse)
	err := c.cc.Invoke(ctx, BudgetService_UpdateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetServiceServer is the server API for BudgetService service.
// All implementations must embed UnimplementedBudgetServiceServer
// for forward compatibility.
type BudgetServiceServer interface {
	CreateBudget(context.Context, *CreateBudgetRequest) (*CreateBudgetResponse, error)
	GetBudgets(context.Context, *GetBudgetsRequest) (*GetBudgetsResponse, error)
	UpdateBudget(context.Context, *UpdateBudgetRequest) (*UpdateBudgetResponse, error)
	mustEmbedUnimplementedBudgetServiceServer()
}

// UnimplementedBudgetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBudgetServiceServer struct{}

func (UnimplementedBudgetServiceServer) CreateBudget(context.Context, *CreateBudgetRequest) (*CreateBudgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBudget not implemented")
}
func (UnimplementedBudgetServiceServer) GetBudgets(context.Context, *GetBudgetsRequest) (*GetBudgetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBudgets not implemented")
}
func (UnimplementedBudgetServiceServer) UpdateBudget(context.Context, *UpdateBudgetRequest) (*UpdateBudgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBudget not implemented")
}
func (UnimplementedBudgetServiceServer) mustEmbedUnimplementedBudgetServiceServer() {}
func (UnimplementedBudgetServiceServer) testEmbeddedByValue()                       {}

// UnsafeBudgetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BudgetServiceServer will
// result in compilation errors.
type UnsafeBudgetServiceServer interface {
	mustEmbedUnimplementedBudgetServiceServer()
}

func RegisterBudgetServiceServer(s grpc.ServiceRegistrar, srv BudgetServiceServer) {
	// If the following call pancis, it indicates UnimplementedBudgetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BudgetService_ServiceDesc, srv)
}

func _BudgetService_CreateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).CreateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_CreateBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).CreateBudget(ctx, req.(*CreateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_GetBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBudgetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_GetBudgets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetBudgets(ctx, req.(*GetBudgetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_UpdateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).UpdateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_UpdateBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).UpdateBudget(ctx, req.(*UpdateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BudgetService_ServiceDesc is the grpc.ServiceDesc for BudgetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BudgetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BudgetService",
	HandlerType: (*BudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBudget",
			Handler:    _BudgetService_CreateBudget_Handler,
		},
		{
			MethodName: "GetBudgets",
			Handler:    _BudgetService_GetBudgets_Handler,
		},
		{
			MethodName: "UpdateBudget",
			Handler:    _BudgetService_UpdateBudget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "budget.proto",
}