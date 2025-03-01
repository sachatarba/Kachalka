// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/sachatarba/course-db/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ITrainerRepository is an autogenerated mock type for the ITrainerRepository type
type ITrainerRepository struct {
	mock.Mock
}

// ChangeTrainer provides a mock function with given fields: ctx, trainer
func (_m *ITrainerRepository) ChangeTrainer(ctx context.Context, trainer entity.Trainer) error {
	ret := _m.Called(ctx, trainer)

	if len(ret) == 0 {
		panic("no return value specified for ChangeTrainer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Trainer) error); ok {
		r0 = rf(ctx, trainer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTrainer provides a mock function with given fields: ctx, trainerID
func (_m *ITrainerRepository) DeleteTrainer(ctx context.Context, trainerID uuid.UUID) error {
	ret := _m.Called(ctx, trainerID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTrainer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, trainerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTrainerByID provides a mock function with given fields: ctx, trainerID
func (_m *ITrainerRepository) GetTrainerByID(ctx context.Context, trainerID uuid.UUID) (entity.Trainer, error) {
	ret := _m.Called(ctx, trainerID)

	if len(ret) == 0 {
		panic("no return value specified for GetTrainerByID")
	}

	var r0 entity.Trainer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entity.Trainer, error)); ok {
		return rf(ctx, trainerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.Trainer); ok {
		r0 = rf(ctx, trainerID)
	} else {
		r0 = ret.Get(0).(entity.Trainer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, trainerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTrainers provides a mock function with given fields: ctx
func (_m *ITrainerRepository) ListTrainers(ctx context.Context) ([]entity.Trainer, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListTrainers")
	}

	var r0 []entity.Trainer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Trainer, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Trainer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Trainer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTrainersByGymID provides a mock function with given fields: ctx, gymID
func (_m *ITrainerRepository) ListTrainersByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.Trainer, error) {
	ret := _m.Called(ctx, gymID)

	if len(ret) == 0 {
		panic("no return value specified for ListTrainersByGymID")
	}

	var r0 []entity.Trainer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]entity.Trainer, error)); ok {
		return rf(ctx, gymID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []entity.Trainer); ok {
		r0 = rf(ctx, gymID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Trainer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, gymID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterNewTrainer provides a mock function with given fields: ctx, trainer
func (_m *ITrainerRepository) RegisterNewTrainer(ctx context.Context, trainer entity.Trainer) error {
	ret := _m.Called(ctx, trainer)

	if len(ret) == 0 {
		panic("no return value specified for RegisterNewTrainer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Trainer) error); ok {
		r0 = rf(ctx, trainer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewITrainerRepository creates a new instance of ITrainerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITrainerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITrainerRepository {
	mock := &ITrainerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
