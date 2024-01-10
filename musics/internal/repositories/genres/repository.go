package genres

import (
	"database/sql"
	"errors"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"net/http"

	"github.com/gofrs/uuid"
)

func GetAllGenres() ([]models.Genre, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var genres []models.Genre
	rows, err := db.Query("SELECT * FROM Genre")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var g models.Genre
		if err := rows.Scan(&g.Id, &g.Name); err != nil {
			return nil, err
		}
		genres = append(genres, g)
	}

	return genres, nil
}

func GetGenreById(id uuid.UUID) (*models.Genre, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var g models.Genre
	err = db.QueryRow("SELECT * FROM Genre WHERE id = ?", id).Scan(&g.Id, &g.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			// When no genre is found, return nil and a custom error
			return nil, &models.CustomError{
				Message: "Genre not found",
				Code:    http.StatusNotFound,
			}
		}
		// Handle other types of errors
		return nil, err
	}

	return &g, nil
}

func PostGenre(newGenre models.Genre) (uuid.UUID, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return uuid.Nil, err
	}
	defer helpers.CloseDB(db)

	// Generate a new UUID if not provided
	if newGenre.Id == uuid.Nil {
		newId, err := uuid.NewV4()
		if err != nil {
			return uuid.Nil, err
		}
		newGenre.Id = newId
	}

	_, err = db.Exec("INSERT INTO Genre (id, name) VALUES (?, ?)", newGenre.Id, newGenre.Name)
	if err != nil {
		return uuid.Nil, err
	}

	return newGenre.Id, nil
}

func PutGenre(genreId uuid.UUID, name string) (models.Genre, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return models.Genre{}, err
	}
	defer db.Close()

	if name == "" {
		return models.Genre{}, errors.New("no fields to update")
	}

	// Execute the update query
	_, err = db.Exec("UPDATE Genre SET name = ? WHERE id = ?", name, genreId)
	if err != nil {
		return models.Genre{}, err
	}

	// Retrieve the updated genre
	var updatedGenre models.Genre
	err = db.QueryRow("SELECT id, name FROM Genre WHERE id = ?", genreId).Scan(&updatedGenre.Id, &updatedGenre.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Genre{}, errors.New("genre not found")
		}
		return models.Genre{}, err
	}

	return updatedGenre, nil
}

func DeleteGenre(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM Genre WHERE id = ?", id)
	return err
}
