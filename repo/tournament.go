package repo

import (
	"log"
	"wc3_game_tracker/api/models"
	"wc3_game_tracker/repo/entities"
)

func GetAllTournaments() ([]models.Tournament, error) {
	db := GetOpenConnection()

	var tournamentEntity []entities.TournamentEntity

	err := db.Select(&tournamentEntity, "SELECT * FROM tournament")
	if err != nil {
		return nil, err
	}

	tournamentModels := make([]models.Tournament, len(tournamentEntity))

	for i, value := range tournamentEntity {
		tournamentModels[i] = models.Tournament{
			Id:              value.Id,
			Name:            value.Name,
			StartDate:       models.CustomTime{Time: value.StartDate},
			EndDate:         models.CustomTime{Time: value.EndDate},
			Type:            value.Type,
			MaxParticipants: value.MaxParticipants,
		}
	}

	return tournamentModels, nil
}

func SaveTournament(tournament *models.Tournament) *models.Tournament {
	db := GetOpenConnection()

	tx := db.MustBegin()
	err := tx.QueryRow("INSERT INTO tournament (name, type, start_date, end_date, max_participants) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		tournament.Name,
		tournament.Type,
		tournament.StartDate.Time,
		tournament.EndDate.Time,
		tournament.MaxParticipants,
	).Scan(&tournament.Id)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return tournament
}
