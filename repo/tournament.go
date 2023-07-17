package repo

import (
	"log"
	"wc3_game_tracker/api/models"
)

func SaveTournament(tournament *models.Tournament) *models.Tournament {
	db := getOpenConnection()

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
