// Code generated by mockery v2.23.1. DO NOT EDIT.

package lists

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockFileDownloader is an autogenerated mock type for the FileDownloader type
type MockFileDownloader struct {
	mock.Mock
}

type MockFileDownloader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFileDownloader) EXPECT() *MockFileDownloader_Expecter {
	return &MockFileDownloader_Expecter{mock: &_m.Mock}
}

// DownloadFile provides a mock function with given fields: link
func (_m *MockFileDownloader) DownloadFile(link string) (io.ReadCloser, error) {
	ret := _m.Called(link)

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (io.ReadCloser, error)); ok {
		return rf(link)
	}
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(link)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(link)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileDownloader_DownloadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadFile'
type MockFileDownloader_DownloadFile_Call struct {
	*mock.Call
}

// DownloadFile is a helper method to define mock.On call
//   - link string
func (_e *MockFileDownloader_Expecter) DownloadFile(link interface{}) *MockFileDownloader_DownloadFile_Call {
	return &MockFileDownloader_DownloadFile_Call{Call: _e.mock.On("DownloadFile", link)}
}

func (_c *MockFileDownloader_DownloadFile_Call) Run(run func(link string)) *MockFileDownloader_DownloadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockFileDownloader_DownloadFile_Call) Return(_a0 io.ReadCloser, _a1 error) *MockFileDownloader_DownloadFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileDownloader_DownloadFile_Call) RunAndReturn(run func(string) (io.ReadCloser, error)) *MockFileDownloader_DownloadFile_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockFileDownloader interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockFileDownloader creates a new instance of MockFileDownloader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFileDownloader(t mockConstructorTestingTNewMockFileDownloader) *MockFileDownloader {
	mock := &MockFileDownloader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
