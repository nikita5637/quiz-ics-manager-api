// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	place "github.com/nikita5637/quiz-registrator-api/pkg/pb/place"
)

// PlaceServiceClient is an autogenerated mock type for the PlaceServiceClient type
type PlaceServiceClient struct {
	mock.Mock
}

type PlaceServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *PlaceServiceClient) EXPECT() *PlaceServiceClient_Expecter {
	return &PlaceServiceClient_Expecter{mock: &_m.Mock}
}

// GetPlace provides a mock function with given fields: ctx, in, opts
func (_m *PlaceServiceClient) GetPlace(ctx context.Context, in *place.GetPlaceRequest, opts ...grpc.CallOption) (*place.Place, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *place.Place
	if rf, ok := ret.Get(0).(func(context.Context, *place.GetPlaceRequest, ...grpc.CallOption) *place.Place); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*place.Place)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *place.GetPlaceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlaceServiceClient_GetPlace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlace'
type PlaceServiceClient_GetPlace_Call struct {
	*mock.Call
}

// GetPlace is a helper method to define mock.On call
//  - ctx context.Context
//  - in *place.GetPlaceRequest
//  - opts ...grpc.CallOption
func (_e *PlaceServiceClient_Expecter) GetPlace(ctx interface{}, in interface{}, opts ...interface{}) *PlaceServiceClient_GetPlace_Call {
	return &PlaceServiceClient_GetPlace_Call{Call: _e.mock.On("GetPlace",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *PlaceServiceClient_GetPlace_Call) Run(run func(ctx context.Context, in *place.GetPlaceRequest, opts ...grpc.CallOption)) *PlaceServiceClient_GetPlace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*place.GetPlaceRequest), variadicArgs...)
	})
	return _c
}

func (_c *PlaceServiceClient_GetPlace_Call) Return(_a0 *place.Place, _a1 error) *PlaceServiceClient_GetPlace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewPlaceServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewPlaceServiceClient creates a new instance of PlaceServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPlaceServiceClient(t mockConstructorTestingTNewPlaceServiceClient) *PlaceServiceClient {
	mock := &PlaceServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
