package musics

import (
	"database/sql"
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

	var musics []models.Music
	rows, err := db.Query("SELECT * FROM Music")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Music
		if err := rows.Scan(&m.Id, &m.Title, &m.GenreId, &m.ArtistId, &m.AlbumId); err != nil {
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
		if err == sql.ErrNoRows {
			return nil, nil // Or a custom error indicating not found
		}
		return nil, err
	}

	return &m, nil
}

func PostMusic(newMusic models.Music) (uuid.UUID, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return uuid.Nil, err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO Music (id, title, genreId, artistId, albumId) VALUES (?, ?, ?, ?, ?)", newMusic.Id, newMusic.Title, newMusic.GenreId, newMusic.ArtistId, newMusic.AlbumId)
	if err != nil {
		return uuid.Nil, err
	}

	return newMusic.Id, nil
}

func PutMusic(id uuid.UUID, updatedMusic models.Music) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE Music SET title = ?, genreId = ?, artistId = ?, albumId = ? WHERE id = ?", updatedMusic.Title, updatedMusic.GenreId, updatedMusic.ArtistId, updatedMusic.AlbumId, id)
	return err
}

func DeleteMusic(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM Music WHERE id = ?", id)
	return err
}
