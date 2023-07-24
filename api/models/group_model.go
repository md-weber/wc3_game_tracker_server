package models

import "github.com/google/uuid"

type Group struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	LeagueId     uuid.UUID `json:"league_id"`
	Admin        string    `json:"admin"`
	Vetos        int       `json:"vetos"`
	MapsPerMatch int       `json:"maps_per_match"`
	Status       string    `json:"status"`
}
