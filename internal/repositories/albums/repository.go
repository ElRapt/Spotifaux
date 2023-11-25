package albums

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllAlbums() ([]models.Album, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var albums []models.Album
	rows, err := db.Query("SELECT * FROM Album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Album
		if err := rows.Scan(&a.Id, &a.Name, &a.ArtistId); err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}

	return albums, nil
}

func GetAlbumById(id uuid.UUID) (*models.Album, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var a models.Album
	err = db.QueryRow("SELECT * FROM Album WHERE id = ?", id).Scan(&a.Id, &a.Name, &a.ArtistId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Or a custom error indicating not found
		}
		return nil, err
	}

	return &a, nil
}

func PostAlbum(newAlbum models.Album) (uuid.UUID, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return uuid.Nil, err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO Album (id, name, artistId) VALUES (?, ?, ?)", newAlbum.Id, newAlbum.Name, newAlbum.ArtistId)
	if err != nil {
		return uuid.Nil, err
	}

	return newAlbum.Id, nil
}

func Put(id uuid.UUID, updatedAlbum models.Album) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE Album SET name = ?, artistId = ? WHERE id = ?", updatedAlbum.Name, updatedAlbum.ArtistId, id)
	return err
}

func DeleteAlbum(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM Album WHERE id = ?", id)
	return err
}
