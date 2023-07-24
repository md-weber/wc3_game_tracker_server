package mocks

import (
	"github.com/stretchr/testify/mock"
	"wc3_game_tracker/api/models"
)

type MockGroupRepo struct {
	mock.Mock
}

func (m *MockGroupRepo) GetAllGroups() ([]models.Group, error) {
	args := m.Called()
	return args.Get(0).([]models.Group), args.Error(1)
}

func (m *MockGroupRepo) SaveGroup(*models.Group) *models.Group {
	args := m.Called()
	return args.Get(0).(*models.Group)
}
