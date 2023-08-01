package services

import (
	"github.com/google/uuid"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo"
)

type GroupServiceI interface {
	GetAllGroups() ([]models.Group, error)
	SaveGroup(*models.Group) (uuid.UUID, error)
	GetGroup(uuid.UUID) (*models.Group, error)
}

type GroupService struct {
	GroupRepo repo.GroupRepoI
}

func (g GroupService) GetAllGroups() ([]models.Group, error) {
	groups, err := g.GroupRepo.GetAllGroups()
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (g GroupService) SaveGroup(group *models.Group) (uuid.UUID, error) {
	result, err := g.GroupRepo.SaveGroup(group)
	return result.Id, err
}

func (g GroupService) GetGroup(id uuid.UUID) (*models.Group, error) {
	result, err := g.GroupRepo.GetGroup(id)
	return result, err
}
