// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: chicken_farmer/v1/farmer.proto

package grpc

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

// FarmerServiceClient is the client API for FarmerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FarmerServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	SpendGoldEggs(ctx context.Context, in *SpendGoldEggsRequest, opts ...grpc.CallOption) (*SpendGoldEggsResponse, error)
	GetGoldEggs(ctx context.Context, in *GetGoldEggsRequest, opts ...grpc.CallOption) (*GetGoldEggsResponse, error)
}

type farmerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFarmerServiceClient(cc grpc.ClientConnInterface) FarmerServiceClient {
	return &farmerServiceClient{cc}
}

func (c *farmerServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/chicken_farmer.v1.FarmerService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmerServiceClient) SpendGoldEggs(ctx context.Context, in *SpendGoldEggsRequest, opts ...grpc.CallOption) (*SpendGoldEggsResponse, error) {
	out := new(SpendGoldEggsResponse)
	err := c.cc.Invoke(ctx, "/chicken_farmer.v1.FarmerService/SpendGoldEggs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmerServiceClient) GetGoldEggs(ctx context.Context, in *GetGoldEggsRequest, opts ...grpc.CallOption) (*GetGoldEggsResponse, error) {
	out := new(GetGoldEggsResponse)
	err := c.cc.Invoke(ctx, "/chicken_farmer.v1.FarmerService/GetGoldEggs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FarmerServiceServer is the server API for FarmerService service.
// All implementations must embed UnimplementedFarmerServiceServer
// for forward compatibility
type FarmerServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	SpendGoldEggs(context.Context, *SpendGoldEggsRequest) (*SpendGoldEggsResponse, error)
	GetGoldEggs(context.Context, *GetGoldEggsRequest) (*GetGoldEggsResponse, error)
	mustEmbedUnimplementedFarmerServiceServer()
}

// UnimplementedFarmerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFarmerServiceServer struct {
}

func (UnimplementedFarmerServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedFarmerServiceServer) SpendGoldEggs(context.Context, *SpendGoldEggsRequest) (*SpendGoldEggsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpendGoldEggs not implemented")
}
func (UnimplementedFarmerServiceServer) GetGoldEggs(context.Context, *GetGoldEggsRequest) (*GetGoldEggsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoldEggs not implemented")
}
func (UnimplementedFarmerServiceServer) mustEmbedUnimplementedFarmerServiceServer() {}

// UnsafeFarmerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FarmerServiceServer will
// result in compilation errors.
type UnsafeFarmerServiceServer interface {
	mustEmbedUnimplementedFarmerServiceServer()
}

func RegisterFarmerServiceServer(s grpc.ServiceRegistrar, srv FarmerServiceServer) {
	s.RegisterService(&FarmerService_ServiceDesc, srv)
}

func _FarmerService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmerServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chicken_farmer.v1.FarmerService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmerServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FarmerService_SpendGoldEggs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpendGoldEggsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmerServiceServer).SpendGoldEggs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chicken_farmer.v1.FarmerService/SpendGoldEggs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmerServiceServer).SpendGoldEggs(ctx, req.(*SpendGoldEggsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FarmerService_GetGoldEggs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoldEggsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmerServiceServer).GetGoldEggs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chicken_farmer.v1.FarmerService/GetGoldEggs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmerServiceServer).GetGoldEggs(ctx, req.(*GetGoldEggsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FarmerService_ServiceDesc is the grpc.ServiceDesc for FarmerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FarmerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chicken_farmer.v1.FarmerService",
	HandlerType: (*FarmerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _FarmerService_Register_Handler,
		},
		{
			MethodName: "SpendGoldEggs",
			Handler:    _FarmerService_SpendGoldEggs_Handler,
		},
		{
			MethodName: "GetGoldEggs",
			Handler:    _FarmerService_GetGoldEggs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chicken_farmer/v1/farmer.proto",
}
