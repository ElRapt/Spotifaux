package musics

import (
	"database/sql"
	"errors"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"strings"

	"github.com/gofrs/uuid"
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
			return nil, nil
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

	// Generate a new UUID if not provided
	if newMusic.Id == uuid.Nil {
		newId, err := uuid.NewV4()
		if err != nil {
			return uuid.Nil, err
		}
		newMusic.Id = newId
	}

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

	var updateParts []string
	var args []interface{}

	if updatedMusic.Title != "" {
		updateParts = append(updateParts, "title = ?")
		args = append(args, updatedMusic.Title)
	}
	if updatedMusic.GenreId != uuid.Nil {
		updateParts = append(updateParts, "genreId = ?")
		args = append(args, updatedMusic.GenreId)
	}
	if updatedMusic.ArtistId != uuid.Nil {
		updateParts = append(updateParts, "artistId = ?")
		args = append(args, updatedMusic.ArtistId)
	}
	if updatedMusic.AlbumId != uuid.Nil {
		updateParts = append(updateParts, "albumId = ?")
		args = append(args, updatedMusic.AlbumId)
	}

	if len(updateParts) == 0 {
		// No fields to update
		return errors.New("no fields provided for update")
	}

	updateQuery := "UPDATE Music SET " + strings.Join(updateParts, ", ") + " WHERE id = ?"
	args = append(args, id)

	_, err = db.Exec(updateQuery, args...)
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
