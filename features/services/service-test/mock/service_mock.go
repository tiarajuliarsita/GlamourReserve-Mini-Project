package mock

import (
	"glamour_reserve/entity/core"

	"github.com/stretchr/testify/mock"
)

type MockServiceRepo struct {
	mock.Mock
}

func (m *MockServiceRepo) FindById(id string) (core.ServiceCore, error) {
	args := m.Called(id)
	return args.Get(0).(core.ServiceCore), args.Error(1)
}

func (m *MockServiceRepo) FindAll(name string, offset, limit int) ([]core.ServiceCore, error) {
	args := m.Called(name, offset, limit)
	return args.Get(0).([]core.ServiceCore), args.Error(1)
}

func (m *MockServiceRepo) Create(service core.ServiceCore) (core.ServiceCore, error) {
	args := m.Called(service)
	return args.Get(0).(core.ServiceCore), args.Error(1)
}

func (m *MockServiceRepo) Delete(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MockServiceRepo) Update(id string, updateSvc core.ServiceCore) (core.ServiceCore, error) {
	args := m.Called(id, updateSvc)
	return args.Get(0).(core.ServiceCore), args.Error(1)
}
