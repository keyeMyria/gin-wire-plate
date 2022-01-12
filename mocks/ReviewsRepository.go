// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "gin-wire-plate/internal/pkg/models"

// ReviewsRepository is an autogenerated mock type for the ReviewsRepository type
type ReviewsRepository struct {
	mock.Mock
}

// Query provides a mock function with given fields: productID
func (_m *ReviewsRepository) Query(productID uint64) ([]*models.Review, error) {
	ret := _m.Called(productID)

	var r0 []*models.Review
	if rf, ok := ret.Get(0).(func(uint64) []*models.Review); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Review)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
