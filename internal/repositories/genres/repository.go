package genres

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
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
			return nil, nil // Or a custom error indicating not found
		}
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

	_, err = db.Exec("INSERT INTO Genre (id, name) VALUES (?, ?)", newGenre.Id, newGenre.Name)
	if err != nil {
		return uuid.Nil, err
	}

	return newGenre.Id, nil
}

func PutGenre(id uuid.UUID, updatedGenre models.Genre) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE Genre SET name = ? WHERE id = ?", updatedGenre.Name, id)
	return err
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
