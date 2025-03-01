// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/sachatarba/course-db/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IMembershipTypeRepository is an autogenerated mock type for the IMembershipTypeRepository type
type IMembershipTypeRepository struct {
	mock.Mock
}

// ChangeMembershipType provides a mock function with given fields: ctx, membershipType
func (_m *IMembershipTypeRepository) ChangeMembershipType(ctx context.Context, membershipType entity.MembershipType) error {
	ret := _m.Called(ctx, membershipType)

	if len(ret) == 0 {
		panic("no return value specified for ChangeMembershipType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.MembershipType) error); ok {
		r0 = rf(ctx, membershipType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMembershipType provides a mock function with given fields: ctx, membershipTypeID
func (_m *IMembershipTypeRepository) DeleteMembershipType(ctx context.Context, membershipTypeID uuid.UUID) error {
	ret := _m.Called(ctx, membershipTypeID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMembershipType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, membershipTypeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMembershipTypeByID provides a mock function with given fields: ctx, membershipTypeID
func (_m *IMembershipTypeRepository) GetMembershipTypeByID(ctx context.Context, membershipTypeID uuid.UUID) (entity.MembershipType, error) {
	ret := _m.Called(ctx, membershipTypeID)

	if len(ret) == 0 {
		panic("no return value specified for GetMembershipTypeByID")
	}

	var r0 entity.MembershipType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entity.MembershipType, error)); ok {
		return rf(ctx, membershipTypeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.MembershipType); ok {
		r0 = rf(ctx, membershipTypeID)
	} else {
		r0 = ret.Get(0).(entity.MembershipType)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, membershipTypeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMembershipTypesByGymID provides a mock function with given fields: ctx, gymID
func (_m *IMembershipTypeRepository) ListMembershipTypesByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.MembershipType, error) {
	ret := _m.Called(ctx, gymID)

	if len(ret) == 0 {
		panic("no return value specified for ListMembershipTypesByGymID")
	}

	var r0 []entity.MembershipType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]entity.MembershipType, error)); ok {
		return rf(ctx, gymID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []entity.MembershipType); ok {
		r0 = rf(ctx, gymID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.MembershipType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, gymID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterNewMembershipType provides a mock function with given fields: ctx, membershipType
func (_m *IMembershipTypeRepository) RegisterNewMembershipType(ctx context.Context, membershipType entity.MembershipType) error {
	ret := _m.Called(ctx, membershipType)

	if len(ret) == 0 {
		panic("no return value specified for RegisterNewMembershipType")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.MembershipType) error); ok {
		r0 = rf(ctx, membershipType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIMembershipTypeRepository creates a new instance of IMembershipTypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMembershipTypeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMembershipTypeRepository {
	mock := &IMembershipTypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
