// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/kritmet/go-gin-todo/domain"
	mock "github.com/stretchr/testify/mock"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// ReadTodoJSON provides a mock function with given fields:
func (_m *TodoRepository) ReadTodoJSON() ([]*domain.Todo, error) {
	ret := _m.Called()

	var r0 []*domain.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteTodoJSON provides a mock function with given fields: entities
func (_m *TodoRepository) WriteTodoJSON(entities []*domain.Todo) error {
	ret := _m.Called(entities)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*domain.Todo) error); ok {
		r0 = rf(entities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTodoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoRepository creates a new instance of TodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoRepository(t mockConstructorTestingTNewTodoRepository) *TodoRepository {
	mock := &TodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}