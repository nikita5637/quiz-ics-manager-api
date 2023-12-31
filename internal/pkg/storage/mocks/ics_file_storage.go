// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mysql "github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage/mysql"
	mock "github.com/stretchr/testify/mock"
)

// ICSFileStorage is an autogenerated mock type for the ICSFileStorage type
type ICSFileStorage struct {
	mock.Mock
}

type ICSFileStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *ICSFileStorage) EXPECT() *ICSFileStorage_Expecter {
	return &ICSFileStorage_Expecter{mock: &_m.Mock}
}

// CreateICSFile provides a mock function with given fields: ctx, dbICSFile
func (_m *ICSFileStorage) CreateICSFile(ctx context.Context, dbICSFile mysql.IcsFile) (int, error) {
	ret := _m.Called(ctx, dbICSFile)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, mysql.IcsFile) int); ok {
		r0 = rf(ctx, dbICSFile)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, mysql.IcsFile) error); ok {
		r1 = rf(ctx, dbICSFile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICSFileStorage_CreateICSFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateICSFile'
type ICSFileStorage_CreateICSFile_Call struct {
	*mock.Call
}

// CreateICSFile is a helper method to define mock.On call
//  - ctx context.Context
//  - dbICSFile mysql.IcsFile
func (_e *ICSFileStorage_Expecter) CreateICSFile(ctx interface{}, dbICSFile interface{}) *ICSFileStorage_CreateICSFile_Call {
	return &ICSFileStorage_CreateICSFile_Call{Call: _e.mock.On("CreateICSFile", ctx, dbICSFile)}
}

func (_c *ICSFileStorage_CreateICSFile_Call) Run(run func(ctx context.Context, dbICSFile mysql.IcsFile)) *ICSFileStorage_CreateICSFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(mysql.IcsFile))
	})
	return _c
}

func (_c *ICSFileStorage_CreateICSFile_Call) Return(_a0 int, _a1 error) *ICSFileStorage_CreateICSFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteICSFile provides a mock function with given fields: ctx, id
func (_m *ICSFileStorage) DeleteICSFile(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICSFileStorage_DeleteICSFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteICSFile'
type ICSFileStorage_DeleteICSFile_Call struct {
	*mock.Call
}

// DeleteICSFile is a helper method to define mock.On call
//  - ctx context.Context
//  - id int
func (_e *ICSFileStorage_Expecter) DeleteICSFile(ctx interface{}, id interface{}) *ICSFileStorage_DeleteICSFile_Call {
	return &ICSFileStorage_DeleteICSFile_Call{Call: _e.mock.On("DeleteICSFile", ctx, id)}
}

func (_c *ICSFileStorage_DeleteICSFile_Call) Run(run func(ctx context.Context, id int)) *ICSFileStorage_DeleteICSFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *ICSFileStorage_DeleteICSFile_Call) Return(_a0 error) *ICSFileStorage_DeleteICSFile_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetICSFileByExternalGameID provides a mock function with given fields: ctx, externalGameID
func (_m *ICSFileStorage) GetICSFileByExternalGameID(ctx context.Context, externalGameID int) (*mysql.IcsFile, error) {
	ret := _m.Called(ctx, externalGameID)

	var r0 *mysql.IcsFile
	if rf, ok := ret.Get(0).(func(context.Context, int) *mysql.IcsFile); ok {
		r0 = rf(ctx, externalGameID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.IcsFile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, externalGameID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICSFileStorage_GetICSFileByExternalGameID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetICSFileByExternalGameID'
type ICSFileStorage_GetICSFileByExternalGameID_Call struct {
	*mock.Call
}

// GetICSFileByExternalGameID is a helper method to define mock.On call
//  - ctx context.Context
//  - externalGameID int
func (_e *ICSFileStorage_Expecter) GetICSFileByExternalGameID(ctx interface{}, externalGameID interface{}) *ICSFileStorage_GetICSFileByExternalGameID_Call {
	return &ICSFileStorage_GetICSFileByExternalGameID_Call{Call: _e.mock.On("GetICSFileByExternalGameID", ctx, externalGameID)}
}

func (_c *ICSFileStorage_GetICSFileByExternalGameID_Call) Run(run func(ctx context.Context, externalGameID int)) *ICSFileStorage_GetICSFileByExternalGameID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *ICSFileStorage_GetICSFileByExternalGameID_Call) Return(_a0 *mysql.IcsFile, _a1 error) *ICSFileStorage_GetICSFileByExternalGameID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetICSFileByID provides a mock function with given fields: ctx, id
func (_m *ICSFileStorage) GetICSFileByID(ctx context.Context, id int) (*mysql.IcsFile, error) {
	ret := _m.Called(ctx, id)

	var r0 *mysql.IcsFile
	if rf, ok := ret.Get(0).(func(context.Context, int) *mysql.IcsFile); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.IcsFile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICSFileStorage_GetICSFileByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetICSFileByID'
type ICSFileStorage_GetICSFileByID_Call struct {
	*mock.Call
}

// GetICSFileByID is a helper method to define mock.On call
//  - ctx context.Context
//  - id int
func (_e *ICSFileStorage_Expecter) GetICSFileByID(ctx interface{}, id interface{}) *ICSFileStorage_GetICSFileByID_Call {
	return &ICSFileStorage_GetICSFileByID_Call{Call: _e.mock.On("GetICSFileByID", ctx, id)}
}

func (_c *ICSFileStorage_GetICSFileByID_Call) Run(run func(ctx context.Context, id int)) *ICSFileStorage_GetICSFileByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *ICSFileStorage_GetICSFileByID_Call) Return(_a0 *mysql.IcsFile, _a1 error) *ICSFileStorage_GetICSFileByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetICSFiles provides a mock function with given fields: ctx
func (_m *ICSFileStorage) GetICSFiles(ctx context.Context) ([]mysql.IcsFile, error) {
	ret := _m.Called(ctx)

	var r0 []mysql.IcsFile
	if rf, ok := ret.Get(0).(func(context.Context) []mysql.IcsFile); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mysql.IcsFile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICSFileStorage_GetICSFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetICSFiles'
type ICSFileStorage_GetICSFiles_Call struct {
	*mock.Call
}

// GetICSFiles is a helper method to define mock.On call
//  - ctx context.Context
func (_e *ICSFileStorage_Expecter) GetICSFiles(ctx interface{}) *ICSFileStorage_GetICSFiles_Call {
	return &ICSFileStorage_GetICSFiles_Call{Call: _e.mock.On("GetICSFiles", ctx)}
}

func (_c *ICSFileStorage_GetICSFiles_Call) Run(run func(ctx context.Context)) *ICSFileStorage_GetICSFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ICSFileStorage_GetICSFiles_Call) Return(_a0 []mysql.IcsFile, _a1 error) *ICSFileStorage_GetICSFiles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewICSFileStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewICSFileStorage creates a new instance of ICSFileStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICSFileStorage(t mockConstructorTestingTNewICSFileStorage) *ICSFileStorage {
	mock := &ICSFileStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
