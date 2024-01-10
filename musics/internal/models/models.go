package models

import (
	"github.com/gofrs/uuid"
)

type Genre struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Artist struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Album struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ArtistId uuid.UUID `json:"artistId"`
}

type Music struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	GenreId  uuid.UUID `json:"genreId"`
	ArtistId uuid.UUID `json:"artistId"`
	AlbumId  uuid.UUID `json:"albumId"`
}
