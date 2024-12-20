// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"
)

// OSProvider is an autogenerated mock type for the OSProvider type
type OSProvider struct {
	mock.Mock
}

type OSProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *OSProvider) EXPECT() *OSProvider_Expecter {
	return &OSProvider_Expecter{mock: &_m.Mock}
}

// ReadDir provides a mock function with given fields: _a0
func (_m *OSProvider) ReadDir(_a0 string) ([]fs.DirEntry, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ReadDir")
	}

	var r0 []fs.DirEntry
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]fs.DirEntry, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) []fs.DirEntry); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]fs.DirEntry)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OSProvider_ReadDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadDir'
type OSProvider_ReadDir_Call struct {
	*mock.Call
}

// ReadDir is a helper method to define mock.On call
//   - _a0 string
func (_e *OSProvider_Expecter) ReadDir(_a0 interface{}) *OSProvider_ReadDir_Call {
	return &OSProvider_ReadDir_Call{Call: _e.mock.On("ReadDir", _a0)}
}

func (_c *OSProvider_ReadDir_Call) Run(run func(_a0 string)) *OSProvider_ReadDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OSProvider_ReadDir_Call) Return(_a0 []fs.DirEntry, _a1 error) *OSProvider_ReadDir_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OSProvider_ReadDir_Call) RunAndReturn(run func(string) ([]fs.DirEntry, error)) *OSProvider_ReadDir_Call {
	_c.Call.Return(run)
	return _c
}

// NewOSProvider creates a new instance of OSProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOSProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *OSProvider {
	mock := &OSProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
