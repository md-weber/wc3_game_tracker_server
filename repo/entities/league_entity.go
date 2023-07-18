package entities

import (
	"github.com/google/uuid"
	"time"
)

type LeagueEntity struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	StartDate time.Time `json:"startDate" db:"start_date"`
	EndDate   time.Time `json:"endDate" db:"end_date"`
	Website   string    `json:"website" db:"website"`
}
