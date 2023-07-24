package entities

import "github.com/google/uuid"

type Group struct {
	Id           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	LeagueId     uuid.UUID `db:"league_id"`
	Admin        string    `db:"admin"`
	Vetos        int       `db:"vetos"`
	MapsPerMatch int       `db:"maps_per_match"`
	Status       string    `db:"status"`
}
