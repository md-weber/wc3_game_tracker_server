package repo

import (
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo/entities"
)

func GetAllGroups() (error, []models.Group) {
	db := GetOpenConnection()

	var groupEntities []entities.Group

	err := db.Select(&groupEntities, `SELECT * FROM "group"`)

	if err != nil {
		return err, nil
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
	return nil, groupModels
}
