// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	users "BE-Sosmed/features/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: UserID
func (_m *Repository) DeleteUser(UserID uint) error {
	ret := _m.Called(UserID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(UserID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertUser provides a mock function with given fields: newUser
func (_m *Repository) InsertUser(newUser users.User) (users.User, error) {
	ret := _m.Called(newUser)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(users.User) (users.User, error)); ok {
		return rf(newUser)
	}
	if rf, ok := ret.Get(0).(func(users.User) users.User); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(users.User) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email
func (_m *Repository) Login(email string) (users.User, error) {
	ret := _m.Called(email)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (users.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) users.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadUserById provides a mock function with given fields: UserID
func (_m *Repository) ReadUserById(UserID uint) (users.User, error) {
	ret := _m.Called(UserID)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (users.User, error)); ok {
		return rf(UserID)
	}
	if rf, ok := ret.Get(0).(func(uint) users.User); ok {
		r0 = rf(UserID)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(UserID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: UserID, updatedUser
func (_m *Repository) UpdateUser(UserID uint, updatedUser users.User) (users.User, error) {
	ret := _m.Called(UserID, updatedUser)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, users.User) (users.User, error)); ok {
		return rf(UserID, updatedUser)
	}
	if rf, ok := ret.Get(0).(func(uint, users.User) users.User); ok {
		r0 = rf(UserID, updatedUser)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(uint, users.User) error); ok {
		r1 = rf(UserID, updatedUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
