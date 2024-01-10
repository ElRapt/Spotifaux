package models

import "github.com/gofrs/uuid"

type Rating struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	SongID  uuid.UUID `json:"song_id"`
	Rating  int       `json:"rating"`
	Comment string    `json:"comment,omitempty"`
}
