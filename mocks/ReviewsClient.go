// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import pb "gin-wire-plate/api/pb"

// ReviewsClient is an autogenerated mock type for the ReviewsClient type
type ReviewsClient struct {
	mock.Mock
}

// Query provides a mock function with given fields: ctx, in, opts
func (_m *ReviewsClient) Query(ctx context.Context, in *pb.QueryReviewsRequest, opts ...grpc.CallOption) (*pb.QueryReviewsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.QueryReviewsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.QueryReviewsRequest, ...grpc.CallOption) *pb.QueryReviewsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.QueryReviewsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.QueryReviewsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
