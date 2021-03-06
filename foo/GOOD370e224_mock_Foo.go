package foo

import mock "github.com/stretchr/testify/mock"

// MockFoo is an autogenerated mock type for the Foo type
type MockFoo struct {
	mock.Mock
}

// Something provides a mock function with given fields:
func (_m *MockFoo) Something() Data {
	ret := _m.Called()

	var r0 Data
	if rf, ok := ret.Get(0).(func() Data); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Data)
	}

	return r0
}
