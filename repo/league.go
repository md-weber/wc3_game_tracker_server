package repo

import (
	"log"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo/entities"
)

func GetAllLeagues() ([]models.League, error) {
	db := GetOpenConnection()

	var leagueEntity []entities.LeagueEntity

	err := db.Select(&leagueEntity, "SELECT * FROM league")
	if err != nil {
		return nil, err
	}

	leagueModels := make([]models.League, len(leagueEntity))

	for i, value := range leagueEntity {
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
	err := tx.QueryRow("INSERT INTO league (name, start_date, end_date, website) VALUES ($1, $2, $3, $4) RETURNING id",
		league.Name,
		league.StartDate.Time,
		league.EndDate.Time,
		league.Website,
	).Scan(&league.Id)

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
