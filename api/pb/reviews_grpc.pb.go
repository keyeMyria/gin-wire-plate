// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: reviews.proto

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

// ReviewsClient is the client API for Reviews service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewsClient interface {
	Query(ctx context.Context, in *QueryReviewsRequest, opts ...grpc.CallOption) (*QueryReviewsResponse, error)
}

type reviewsClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewsClient(cc grpc.ClientConnInterface) ReviewsClient {
	return &reviewsClient{cc}
}

func (c *reviewsClient) Query(ctx context.Context, in *QueryReviewsRequest, opts ...grpc.CallOption) (*QueryReviewsResponse, error) {
	out := new(QueryReviewsResponse)
	err := c.cc.Invoke(ctx, "/Reviews/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewsServer is the server API for Reviews service.
// All implementations must embed UnimplementedReviewsServer
// for forward compatibility
type ReviewsServer interface {
	Query(context.Context, *QueryReviewsRequest) (*QueryReviewsResponse, error)
	mustEmbedUnimplementedReviewsServer()
}

// UnimplementedReviewsServer must be embedded to have forward compatible implementations.
type UnimplementedReviewsServer struct {
}

func (UnimplementedReviewsServer) Query(context.Context, *QueryReviewsRequest) (*QueryReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedReviewsServer) mustEmbedUnimplementedReviewsServer() {}

// UnsafeReviewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewsServer will
// result in compilation errors.
type UnsafeReviewsServer interface {
	mustEmbedUnimplementedReviewsServer()
}

func RegisterReviewsServer(s grpc.ServiceRegistrar, srv ReviewsServer) {
	s.RegisterService(&Reviews_ServiceDesc, srv)
}

func _Reviews_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Reviews/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).Query(ctx, req.(*QueryReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reviews_ServiceDesc is the grpc.ServiceDesc for Reviews service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reviews_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Reviews",
	HandlerType: (*ReviewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Reviews_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reviews.proto",
}
