// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	pagination "charum/dto/pagination"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	users "charum/business/users"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UseCase) Delete(id primitive.ObjectID) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *UseCase) GetAll() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *UseCase) GetByID(id primitive.ObjectID) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManyWithPagination provides a mock function with given fields: _a0, domain
func (_m *UseCase) GetManyWithPagination(_a0 pagination.Request, domain *users.Domain) ([]users.Domain, int, int, error) {
	ret := _m.Called(_a0, domain)

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func(pagination.Request, *users.Domain) []users.Domain); ok {
		r0 = rf(_a0, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(pagination.Request, *users.Domain) int); ok {
		r1 = rf(_a0, domain)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(pagination.Request, *users.Domain) int); ok {
		r2 = rf(_a0, domain)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(pagination.Request, *users.Domain) error); ok {
		r3 = rf(_a0, domain)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Login provides a mock function with given fields: key, password
func (_m *UseCase) Login(key string, password string) (users.Domain, string, error) {
	ret := _m.Called(key, password)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(string, string) users.Domain); ok {
		r0 = rf(key, password)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(key, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(key, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: domain, profilePicture
func (_m *UseCase) Register(domain *users.Domain, profilePicture *multipart.FileHeader) (users.Domain, string, error) {
	ret := _m.Called(domain, profilePicture)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain, *multipart.FileHeader) users.Domain); ok {
		r0 = rf(domain, profilePicture)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*users.Domain, *multipart.FileHeader) string); ok {
		r1 = rf(domain, profilePicture)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*users.Domain, *multipart.FileHeader) error); ok {
		r2 = rf(domain, profilePicture)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Suspend provides a mock function with given fields: id
func (_m *UseCase) Suspend(id primitive.ObjectID) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unsuspend provides a mock function with given fields: id
func (_m *UseCase) Unsuspend(id primitive.ObjectID) (users.Domain, error) {
	ret := _m.Called(id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) users.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: domain, profilePicture
func (_m *UseCase) Update(domain *users.Domain, profilePicture *multipart.FileHeader) (users.Domain, error) {
	ret := _m.Called(domain, profilePicture)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain, *multipart.FileHeader) users.Domain); ok {
		r0 = rf(domain, profilePicture)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain, *multipart.FileHeader) error); ok {
		r1 = rf(domain, profilePicture)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: domain
func (_m *UseCase) UpdatePassword(domain *users.Domain) (users.Domain, error) {
	ret := _m.Called(domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*users.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}