package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type CustomTime struct {
	time.Time
}

type Tournament struct {
	Id              uuid.UUID  `json:"id" db:"id"`
	Name            string     `json:"name" db:"name"`
	StartDate       CustomTime `json:"startDate" db:"start_date"`
	EndDate         CustomTime `json:"endDate" db:"end_date"`
	Type            string     `json:"type" db:"type"`
	MaxParticipants int        `json:"maxParticipants" db:"max_participants"`
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	const shortForm = "2006-01-02T15:04:05Z07:00"

	// Unmarshal the bytes to a string to remove the quotes
	var dateString string
	if err = json.Unmarshal(b, &dateString); err != nil {
		return err
	}

	fmt.Println(dateString)
	date, err := time.Parse(shortForm, dateString)
	if err != nil {
		return err
	}
	t.Time = date
	return
}
