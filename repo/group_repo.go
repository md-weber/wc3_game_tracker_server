package repo

import (
	"github.com/google/uuid"
	"log"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo/entities"
)

type GroupRepoI interface {
	GetAllGroups() ([]models.Group, error)
	SaveGroup(*models.Group) (*models.Group, error)
	GetGroup(uuid.UUID) (*models.Group, error)
}

type GroupRepository struct {
}

func (g GroupRepository) GetGroup(id uuid.UUID) (*models.Group, error) {
	db := GetOpenConnection()
	var groupEntity entities.Group

	err := db.Get(&groupEntity, "SELECT * FROM public.group WHERE id = $1", id.String())
	if err != nil {
		return nil, err
	}

	groupModel := new(models.Group)
	groupModel.Id = groupEntity.Id
	groupModel.Name = groupEntity.Name
	groupModel.Admin = groupEntity.Admin
	groupModel.Vetos = groupEntity.Vetos
	groupModel.MapsPerMatch = groupEntity.MapsPerMatch
	groupModel.Status = groupEntity.Status
	groupModel.LeagueId = groupEntity.LeagueId

	return groupModel, nil
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

func (g GroupRepository) SaveGroup(group *models.Group) (*models.Group, error) {
	db := GetOpenConnection()
	tx := db.MustBegin()

	if group.Id == uuid.Nil {
		group.Id = uuid.New()
	}

	err := tx.QueryRow("INSERT INTO public.group (id, name, league_id, admin, vetos, maps_per_match, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id ",
		group.Id,
		group.Name,
		group.LeagueId,
		group.Admin,
		group.Vetos,
		group.MapsPerMatch,
		group.Status,
	).Scan(&group.Id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return group, nil
}
