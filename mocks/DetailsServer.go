// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import pb "gin-wire-plate/api/pb"

// DetailsServer is an autogenerated mock type for the DetailsServer type
type DetailsServer struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *DetailsServer) Get(_a0 context.Context, _a1 *pb.GetDetailRequest) (*pb.Detail, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.Detail
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetDetailRequest) *pb.Detail); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Detail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetDetailRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
