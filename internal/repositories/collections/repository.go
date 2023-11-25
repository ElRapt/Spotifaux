package collections

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllMusics() ([]models.Music, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)
	rows, err := db.Query("SELECT * FROM Music")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var musics []models.Music
	for rows.Next() {
		var m models.Music
		err := rows.Scan(&m.Id, &m.Title, &m.GenreId, &m.ArtistId, &m.AlbumId)
		if err != nil {
			return nil, err
		}
		musics = append(musics, m)
	}
	return musics, nil
}

func GetMusicById(id uuid.UUID) (*models.Music, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)
	var m models.Music
	err = db.QueryRow("SELECT * FROM Music WHERE id = ?", id).Scan(&m.Id, &m.Title, &m.GenreId, &m.ArtistId, &m.AlbumId)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
