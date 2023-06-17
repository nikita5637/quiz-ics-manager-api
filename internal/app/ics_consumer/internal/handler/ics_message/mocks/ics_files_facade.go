// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/nikita5637/quiz-ics-manager-api/internal/pkg/model"
)

// ICSFilesFacade is an autogenerated mock type for the ICSFilesFacade type
type ICSFilesFacade struct {
	mock.Mock
}

type ICSFilesFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *ICSFilesFacade) EXPECT() *ICSFilesFacade_Expecter {
	return &ICSFilesFacade_Expecter{mock: &_m.Mock}
}

// CreateICSFile provides a mock function with given fields: ctx, icsFile
func (_m *ICSFilesFacade) CreateICSFile(ctx context.Context, icsFile model.ICSFile) (model.ICSFile, error) {
	ret := _m.Called(ctx, icsFile)

	var r0 model.ICSFile
	if rf, ok := ret.Get(0).(func(context.Context, model.ICSFile) model.ICSFile); ok {
		r0 = rf(ctx, icsFile)
	} else {
		r0 = ret.Get(0).(model.ICSFile)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.ICSFile) error); ok {
		r1 = rf(ctx, icsFile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICSFilesFacade_CreateICSFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateICSFile'
type ICSFilesFacade_CreateICSFile_Call struct {
	*mock.Call
}

// CreateICSFile is a helper method to define mock.On call
//  - ctx context.Context
//  - icsFile model.ICSFile
func (_e *ICSFilesFacade_Expecter) CreateICSFile(ctx interface{}, icsFile interface{}) *ICSFilesFacade_CreateICSFile_Call {
	return &ICSFilesFacade_CreateICSFile_Call{Call: _e.mock.On("CreateICSFile", ctx, icsFile)}
}

func (_c *ICSFilesFacade_CreateICSFile_Call) Run(run func(ctx context.Context, icsFile model.ICSFile)) *ICSFilesFacade_CreateICSFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.ICSFile))
	})
	return _c
}

func (_c *ICSFilesFacade_CreateICSFile_Call) Return(_a0 model.ICSFile, _a1 error) *ICSFilesFacade_CreateICSFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteICSFile provides a mock function with given fields: ctx, id
func (_m *ICSFilesFacade) DeleteICSFile(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICSFilesFacade_DeleteICSFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteICSFile'
type ICSFilesFacade_DeleteICSFile_Call struct {
	*mock.Call
}

// DeleteICSFile is a helper method to define mock.On call
//  - ctx context.Context
//  - id int32
func (_e *ICSFilesFacade_Expecter) DeleteICSFile(ctx interface{}, id interface{}) *ICSFilesFacade_DeleteICSFile_Call {
	return &ICSFilesFacade_DeleteICSFile_Call{Call: _e.mock.On("DeleteICSFile", ctx, id)}
}

func (_c *ICSFilesFacade_DeleteICSFile_Call) Run(run func(ctx context.Context, id int32)) *ICSFilesFacade_DeleteICSFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *ICSFilesFacade_DeleteICSFile_Call) Return(_a0 error) *ICSFilesFacade_DeleteICSFile_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetICSFileByGameID provides a mock function with given fields: ctx, gameID
func (_m *ICSFilesFacade) GetICSFileByGameID(ctx context.Context, gameID int32) (model.ICSFile, error) {
	ret := _m.Called(ctx, gameID)

	var r0 model.ICSFile
	if rf, ok := ret.Get(0).(func(context.Context, int32) model.ICSFile); ok {
		r0 = rf(ctx, gameID)
	} else {
		r0 = ret.Get(0).(model.ICSFile)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, gameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICSFilesFacade_GetICSFileByGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetICSFileByGameID'
type ICSFilesFacade_GetICSFileByGameID_Call struct {
	*mock.Call
}

// GetICSFileByGameID is a helper method to define mock.On call
//  - ctx context.Context
//  - gameID int32
func (_e *ICSFilesFacade_Expecter) GetICSFileByGameID(ctx interface{}, gameID interface{}) *ICSFilesFacade_GetICSFileByGameID_Call {
	return &ICSFilesFacade_GetICSFileByGameID_Call{Call: _e.mock.On("GetICSFileByGameID", ctx, gameID)}
}

func (_c *ICSFilesFacade_GetICSFileByGameID_Call) Run(run func(ctx context.Context, gameID int32)) *ICSFilesFacade_GetICSFileByGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *ICSFilesFacade_GetICSFileByGameID_Call) Return(_a0 model.ICSFile, _a1 error) *ICSFilesFacade_GetICSFileByGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewICSFilesFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewICSFilesFacade creates a new instance of ICSFilesFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICSFilesFacade(t mockConstructorTestingTNewICSFilesFacade) *ICSFilesFacade {
	mock := &ICSFilesFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
