// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	context "context"
	service "projects/LDmitryLD/repository/app/internal/modules/user/service"

	mock "github.com/stretchr/testify/mock"
)

// Userer is an autogenerated mock type for the Userer type
type Userer struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in
func (_m *Userer) Create(ctx context.Context, in service.CreateIn) service.CreateOut {
	ret := _m.Called(ctx, in)

	var r0 service.CreateOut
	if rf, ok := ret.Get(0).(func(context.Context, service.CreateIn) service.CreateOut); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(service.CreateOut)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, in
func (_m *Userer) Delete(ctx context.Context, in service.DeleteIn) service.DeleteOut {
	ret := _m.Called(ctx, in)

	var r0 service.DeleteOut
	if rf, ok := ret.Get(0).(func(context.Context, service.DeleteIn) service.DeleteOut); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(service.DeleteOut)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, in
func (_m *Userer) GetByID(ctx context.Context, in service.GetByIDIn) service.GetByIDOut {
	ret := _m.Called(ctx, in)

	var r0 service.GetByIDOut
	if rf, ok := ret.Get(0).(func(context.Context, service.GetByIDIn) service.GetByIDOut); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(service.GetByIDOut)
	}

	return r0
}

// List provides a mock function with given fields: ctx, in
func (_m *Userer) List(ctx context.Context, in service.ListIn) service.ListOut {
	ret := _m.Called(ctx, in)

	var r0 service.ListOut
	if rf, ok := ret.Get(0).(func(context.Context, service.ListIn) service.ListOut); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(service.ListOut)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, in
func (_m *Userer) Update(ctx context.Context, in service.UpdateIn) service.UpdateOut {
	ret := _m.Called(ctx, in)

	var r0 service.UpdateOut
	if rf, ok := ret.Get(0).(func(context.Context, service.UpdateIn) service.UpdateOut); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(service.UpdateOut)
	}

	return r0
}

// NewUserer creates a new instance of Userer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Userer {
	mock := &Userer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
