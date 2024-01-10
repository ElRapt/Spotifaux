package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Rating struct {
	Id         uuid.UUID `json:"id"`
	Comment    string    `json:"comment"`
	Rating     int       `json:"rating"`
	RatingDate time.Time `json:"rating_date"`
	MusicID    uuid.UUID `json:"music_id"`
	UserID     uuid.UUID `json:"user_id"`
}

type RatingRequest struct {
	Comment *string `json:"comment"`
	Rating  *int    `json:"rating"`
	UserID  *string `json:"user_id"`
}
