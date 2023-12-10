package ratings

import (
	"database/sql"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllRatings() ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var ratings []models.Rating
	rows, err := db.Query("SELECT * FROM ratings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Rating
		if err := rows.Scan(&r.ID, &r.UserID, &r.SongID, &r.Rating, &r.Comment); err != nil {
			return nil, err
		}
		ratings = append(ratings, r)
	}

	return ratings, nil
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	var r models.Rating
	err = db.QueryRow("SELECT * FROM ratings WHERE id = ?", id).Scan(&r.ID, &r.UserID, &r.SongID, &r.Rating, &r.Comment)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &r, nil
}

func PostRating(newRating models.Rating) (uuid.UUID, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return uuid.Nil, err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO ratings (id, user_id, song_id, rating, comment) VALUES (?, ?, ?, ?, ?)", newRating.ID, newRating.UserID, newRating.SongID, newRating.Rating, newRating.Comment)
	if err != nil {
		return uuid.Nil, err
	}

	return newRating.ID, nil
}

func PutRating(id uuid.UUID, updatedRating models.Rating) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE ratings SET rating = ?, comment = ? WHERE id = ?", updatedRating.Rating, updatedRating.Comment, id)
	return err
}

func DeleteRating(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM ratings WHERE id = ?", id)
	return err
}
