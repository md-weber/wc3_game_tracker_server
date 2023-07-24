package services

import (
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo"
)

type GroupServiceI interface {
	GetAllGroups() ([]models.Group, error)
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
