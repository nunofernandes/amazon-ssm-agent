// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Expirer is an autogenerated mock type for the Expirer type
type Expirer struct {
	mock.Mock
}

// ExpiresAt provides a mock function with given fields:
func (_m *Expirer) ExpiresAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}