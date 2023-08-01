package repo

import (
	"github.com/google/uuid"
	"log"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo/entities"
)

func GetLeague(id uuid.UUID) (*models.League, error) {
	db := GetOpenConnection()
	var leagueEntity entities.LeagueEntity

	err := db.Get(&leagueEntity, "SELECT * FROM league WHERE id = $1", id.String())
	if err != nil {
		return nil, err
	}

	leagueModel := new(models.League)
	leagueModel.Id = leagueEntity.Id
	leagueModel.Name = leagueEntity.Name
	leagueModel.StartDate = models.CustomTime{Time: leagueEntity.StartDate}
	leagueModel.EndDate = models.CustomTime{Time: leagueEntity.EndDate}
	leagueModel.Website = leagueEntity.Website

	return leagueModel, err
}

func GetAllLeagues() ([]models.League, error) {
	db := GetOpenConnection()

	var leagueEntities []entities.LeagueEntity

	err := db.Select(&leagueEntities, "SELECT * FROM league")
	if err != nil {
		return nil, err
	}

	leagueModels := make([]models.League, len(leagueEntities))

	for i, value := range leagueEntities {
		leagueModels[i] = models.League{
			Id:        value.Id,
			Name:      value.Name,
			StartDate: models.CustomTime{Time: value.StartDate},
			EndDate:   models.CustomTime{Time: value.EndDate},
			Website:   value.Website,
		}
	}

	return leagueModels, nil
}

func SaveLeague(league *models.League) *models.League {
	db := GetOpenConnection()

	tx := db.MustBegin()

	var err error
	if league.Id == uuid.Nil {
		err = tx.QueryRow("INSERT INTO league (name, start_date, end_date, website) VALUES ($1, $2, $3, $4) RETURNING id",
			league.Name,
			league.StartDate.Time,
			league.EndDate.Time,
			league.Website,
		).Scan(&league.Id)
	} else {
		err = tx.QueryRow("INSERT INTO league (id, name, start_date, end_date, website) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			league.Id,
			league.Name,
			league.StartDate.Time,
			league.EndDate.Time,
			league.Website,
		).Scan(&league.Id)
	}

	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return league
}
