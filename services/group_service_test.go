package services

import (
	"github.com/google/uuid"
	"testing"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAllGroups(t *testing.T) {
	// Create an instance of our test object.
	mockRepo := new(mocks.MockGroupRepo)

	// Set up expectation.
	mockRepo.On("GetAllGroups").Return([]models.Group{
		{
			Id:   uuid.MustParse("2e01589b-47cb-4d0e-b248-a2c837bcfa98"),
			Name: "Group1",
		},
		{
			Id:   uuid.MustParse("6c75ad7d-6dfa-4f14-8d71-df41fac6fead"),
			Name: "Group2",
		},
	}, nil)

	// Create an instance of our test object.
	service := GroupService{
		GroupRepo: mockRepo,
	}

	// Call the method we are testing.
	groups, err := service.GetAllGroups()

	// Assert expectations.
	assert.NoError(t, err)
	assert.Len(t, groups, 2)
	assert.Equal(t, "Group1", groups[0].Name)
	assert.Equal(t, "Group2", groups[1].Name)
	mockRepo.AssertExpectations(t)
}
