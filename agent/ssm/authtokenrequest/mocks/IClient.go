// Code generated by mockery v2.23.2. DO NOT EDIT.

package mocks

import (
	context "context"

	ssm "github.com/aws/aws-sdk-go/service/ssm"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// RequestManagedInstanceRoleToken provides a mock function with given fields: fingerprint
func (_m *IClient) RequestManagedInstanceRoleToken(fingerprint string) (*ssm.RequestManagedInstanceRoleTokenOutput, error) {
	ret := _m.Called(fingerprint)

	var r0 *ssm.RequestManagedInstanceRoleTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*ssm.RequestManagedInstanceRoleTokenOutput, error)); ok {
		return rf(fingerprint)
	}
	if rf, ok := ret.Get(0).(func(string) *ssm.RequestManagedInstanceRoleTokenOutput); ok {
		r0 = rf(fingerprint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.RequestManagedInstanceRoleTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fingerprint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestManagedInstanceRoleTokenWithContext provides a mock function with given fields: ctx, fingerprint
func (_m *IClient) RequestManagedInstanceRoleTokenWithContext(ctx context.Context, fingerprint string) (*ssm.RequestManagedInstanceRoleTokenOutput, error) {
	ret := _m.Called(ctx, fingerprint)

	var r0 *ssm.RequestManagedInstanceRoleTokenOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ssm.RequestManagedInstanceRoleTokenOutput, error)); ok {
		return rf(ctx, fingerprint)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ssm.RequestManagedInstanceRoleTokenOutput); ok {
		r0 = rf(ctx, fingerprint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.RequestManagedInstanceRoleTokenOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fingerprint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateManagedInstancePublicKey provides a mock function with given fields: publicKey, publicKeyType
func (_m *IClient) UpdateManagedInstancePublicKey(publicKey string, publicKeyType string) (*ssm.UpdateManagedInstancePublicKeyOutput, error) {
	ret := _m.Called(publicKey, publicKeyType)

	var r0 *ssm.UpdateManagedInstancePublicKeyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*ssm.UpdateManagedInstancePublicKeyOutput, error)); ok {
		return rf(publicKey, publicKeyType)
	}
	if rf, ok := ret.Get(0).(func(string, string) *ssm.UpdateManagedInstancePublicKeyOutput); ok {
		r0 = rf(publicKey, publicKeyType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.UpdateManagedInstancePublicKeyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(publicKey, publicKeyType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateManagedInstancePublicKeyWithContext provides a mock function with given fields: ctx, publicKey, publicKeyType
func (_m *IClient) UpdateManagedInstancePublicKeyWithContext(ctx context.Context, publicKey string, publicKeyType string) (*ssm.UpdateManagedInstancePublicKeyOutput, error) {
	ret := _m.Called(ctx, publicKey, publicKeyType)

	var r0 *ssm.UpdateManagedInstancePublicKeyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*ssm.UpdateManagedInstancePublicKeyOutput, error)); ok {
		return rf(ctx, publicKey, publicKeyType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *ssm.UpdateManagedInstancePublicKeyOutput); ok {
		r0 = rf(ctx, publicKey, publicKeyType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssm.UpdateManagedInstancePublicKeyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, publicKey, publicKeyType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewIClient creates a new instance of IClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIClient(t mockConstructorTestingTNewIClient) *IClient {
	mock := &IClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
