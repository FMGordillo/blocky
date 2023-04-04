// Code generated by mockery v2.23.1. DO NOT EDIT.

package resolver

import (
	logrus "github.com/sirupsen/logrus"
	mock "github.com/stretchr/testify/mock"

	model "github.com/0xERR0R/blocky/model"
)

// MockResolver is an autogenerated mock type for the Resolver type
type MockResolver struct {
	mock.Mock
}

type MockResolver_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResolver) EXPECT() *MockResolver_Expecter {
	return &MockResolver_Expecter{mock: &_m.Mock}
}

// IsEnabled provides a mock function with given fields:
func (_m *MockResolver) IsEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockResolver_IsEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsEnabled'
type MockResolver_IsEnabled_Call struct {
	*mock.Call
}

// IsEnabled is a helper method to define mock.On call
func (_e *MockResolver_Expecter) IsEnabled() *MockResolver_IsEnabled_Call {
	return &MockResolver_IsEnabled_Call{Call: _e.mock.On("IsEnabled")}
}

func (_c *MockResolver_IsEnabled_Call) Run(run func()) *MockResolver_IsEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResolver_IsEnabled_Call) Return(_a0 bool) *MockResolver_IsEnabled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockResolver_IsEnabled_Call) RunAndReturn(run func() bool) *MockResolver_IsEnabled_Call {
	_c.Call.Return(run)
	return _c
}

// LogConfig provides a mock function with given fields: _a0
func (_m *MockResolver) LogConfig(_a0 *logrus.Entry) {
	_m.Called(_a0)
}

// MockResolver_LogConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogConfig'
type MockResolver_LogConfig_Call struct {
	*mock.Call
}

// LogConfig is a helper method to define mock.On call
//   - _a0 *logrus.Entry
func (_e *MockResolver_Expecter) LogConfig(_a0 interface{}) *MockResolver_LogConfig_Call {
	return &MockResolver_LogConfig_Call{Call: _e.mock.On("LogConfig", _a0)}
}

func (_c *MockResolver_LogConfig_Call) Run(run func(_a0 *logrus.Entry)) *MockResolver_LogConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*logrus.Entry))
	})
	return _c
}

func (_c *MockResolver_LogConfig_Call) Return() *MockResolver_LogConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockResolver_LogConfig_Call) RunAndReturn(run func(*logrus.Entry)) *MockResolver_LogConfig_Call {
	_c.Call.Return(run)
	return _c
}

// Resolve provides a mock function with given fields: req
func (_m *MockResolver) Resolve(req *model.Request) (*model.Response, error) {
	ret := _m.Called(req)

	var r0 *model.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Request) (*model.Response, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.Request) *model.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResolver_Resolve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resolve'
type MockResolver_Resolve_Call struct {
	*mock.Call
}

// Resolve is a helper method to define mock.On call
//   - req *model.Request
func (_e *MockResolver_Expecter) Resolve(req interface{}) *MockResolver_Resolve_Call {
	return &MockResolver_Resolve_Call{Call: _e.mock.On("Resolve", req)}
}

func (_c *MockResolver_Resolve_Call) Run(run func(req *model.Request)) *MockResolver_Resolve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Request))
	})
	return _c
}

func (_c *MockResolver_Resolve_Call) Return(_a0 *model.Response, _a1 error) *MockResolver_Resolve_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResolver_Resolve_Call) RunAndReturn(run func(*model.Request) (*model.Response, error)) *MockResolver_Resolve_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *MockResolver) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockResolver_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type MockResolver_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *MockResolver_Expecter) Type() *MockResolver_Type_Call {
	return &MockResolver_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *MockResolver_Type_Call) Run(run func()) *MockResolver_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockResolver_Type_Call) Return(_a0 string) *MockResolver_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockResolver_Type_Call) RunAndReturn(run func() string) *MockResolver_Type_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockResolver creates a new instance of MockResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockResolver(t mockConstructorTestingTNewMockResolver) *MockResolver {
	mock := &MockResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
