package artists

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllArtists() ([]models.Artist, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var artists []models.Artist
	rows, err := db.Query("SELECT * FROM Artist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Artist
		if err := rows.Scan(&a.Id, &a.Name); err != nil {
			return nil, err
		}
		artists = append(artists, a)
	}

	return artists, nil
}

func GetArtistById(id uuid.UUID) (*models.Artist, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var a models.Artist
	err = db.QueryRow("SELECT * FROM Artist WHERE id = ?", id).Scan(&a.Id, &a.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Or a custom error indicating not found
		}
		return nil, err
	}

	return &a, nil
}

func PostArtist(newArtist models.Artist) (uuid.UUID, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return uuid.Nil, err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO Artist (id, name) VALUES (?, ?)", newArtist.Id, newArtist.Name)
	if err != nil {
		return uuid.Nil, err
	}

	return newArtist.Id, nil
}

func PutArtist(id uuid.UUID, updatedArtist models.Artist) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE Artist SET name = ? WHERE id = ?", updatedArtist.Name, id)
	return err
}

func DeleteArtist(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM Artist WHERE id = ?", id)
	return err
}
