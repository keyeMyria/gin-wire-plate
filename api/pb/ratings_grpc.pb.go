// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: ratings.proto

package pb

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

// RatingsClient is the client API for Ratings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingsClient interface {
	Get(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*Rating, error)
}

type ratingsClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingsClient(cc grpc.ClientConnInterface) RatingsClient {
	return &ratingsClient{cc}
}

func (c *ratingsClient) Get(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*Rating, error) {
	out := new(Rating)
	err := c.cc.Invoke(ctx, "/Ratings/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingsServer is the server API for Ratings service.
// All implementations should embed UnimplementedRatingsServer
// for forward compatibility
type RatingsServer interface {
	Get(context.Context, *GetRatingRequest) (*Rating, error)
}

// UnimplementedRatingsServer should be embedded to have forward compatible implementations.
type UnimplementedRatingsServer struct {
}

func (UnimplementedRatingsServer) Get(context.Context, *GetRatingRequest) (*Rating, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeRatingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingsServer will
// result in compilation errors.
type UnsafeRatingsServer interface {
	mustEmbedUnimplementedRatingsServer()
}

func RegisterRatingsServer(s grpc.ServiceRegistrar, srv RatingsServer) {
	s.RegisterService(&Ratings_ServiceDesc, srv)
}

func _Ratings_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ratings/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingsServer).Get(ctx, req.(*GetRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ratings_ServiceDesc is the grpc.ServiceDesc for Ratings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ratings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ratings",
	HandlerType: (*RatingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Ratings_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ratings.proto",
}
