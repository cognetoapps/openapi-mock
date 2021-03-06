// Code generated by mockery v1.0.0. DO NOT EDIT.

package data

import mock "github.com/stretchr/testify/mock"

// mockArrayLengthGenerator is an autogenerated mock type for the arrayLengthGenerator type
type mockArrayLengthGenerator struct {
	mock.Mock
}

// GenerateLength provides a mock function with given fields: min, max
func (_m *mockArrayLengthGenerator) GenerateLength(min uint64, max uint64) (uint64, uint64) {
	ret := _m.Called(min, max)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64, uint64) uint64); ok {
		r0 = rf(min, max)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(uint64, uint64) uint64); ok {
		r1 = rf(min, max)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}
