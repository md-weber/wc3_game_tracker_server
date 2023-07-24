package repo

import (
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo/entities"
)

type GroupRepoI interface {
	GetAllGroups() ([]models.Group, error)
}

type GroupRepository struct {
}

func (g GroupRepository) GetAllGroups() ([]models.Group, error) {
	db := GetOpenConnection()

	var groupEntities []entities.Group

	err := db.Select(&groupEntities, `SELECT * FROM "group"`)

	if err != nil {
		return nil, err
	}

	groupModels := make([]models.Group, len(groupEntities))

	for i, e := range groupEntities {
		groupModels[i] = models.Group{
			Id:           e.Id,
			Name:         e.Name,
			LeagueId:     e.LeagueId,
			Admin:        e.Admin,
			Vetos:        e.Vetos,
			MapsPerMatch: e.MapsPerMatch,
			Status:       e.Status,
		}
	}
	return groupModels, nil
}
