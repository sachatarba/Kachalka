// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/sachatarba/course-db/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IAuthorizationService is an autogenerated mock type for the IAuthorizationService type
type IAuthorizationService struct {
	mock.Mock
}

// Authorize provides a mock function with given fields: ctx, login, password
func (_m *IAuthorizationService) Authorize(ctx context.Context, login string, password string) (entity.Session, error) {
	ret := _m.Called(ctx, login, password)

	if len(ret) == 0 {
		panic("no return value specified for Authorize")
	}

	var r0 entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (entity.Session, error)); ok {
		return rf(ctx, login, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entity.Session); ok {
		r0 = rf(ctx, login, password)
	} else {
		r0 = ret.Get(0).(entity.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteClient provides a mock function with given fields: ctx, clientID
func (_m *IAuthorizationService) DeleteClient(ctx context.Context, clientID uuid.UUID) (entity.Session, error) {
	ret := _m.Called(ctx, clientID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteClient")
	}

	var r0 entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entity.Session, error)); ok {
		return rf(ctx, clientID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.Session); ok {
		r0 = rf(ctx, clientID)
	} else {
		r0 = ret.Get(0).(entity.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthorize provides a mock function with given fields: ctx, sessionID
func (_m *IAuthorizationService) IsAuthorize(ctx context.Context, sessionID uuid.UUID) (*entity.Session, error) {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for IsAuthorize")
	}

	var r0 *entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entity.Session, error)); ok {
		return rf(ctx, sessionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, sessionID
func (_m *IAuthorizationService) Logout(ctx context.Context, sessionID uuid.UUID) (entity.Session, error) {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for Logout")
	}

	var r0 entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entity.Session, error)); ok {
		return rf(ctx, sessionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Get(0).(entity.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, client
func (_m *IAuthorizationService) Register(ctx context.Context, client entity.Client) (entity.Session, error) {
	ret := _m.Called(ctx, client)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Client) (entity.Session, error)); ok {
		return rf(ctx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Client) entity.Session); ok {
		r0 = rf(ctx, client)
	} else {
		r0 = ret.Get(0).(entity.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Client) error); ok {
		r1 = rf(ctx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIAuthorizationService creates a new instance of IAuthorizationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthorizationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthorizationService {
	mock := &IAuthorizationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
