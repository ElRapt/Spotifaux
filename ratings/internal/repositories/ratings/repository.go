package ratings

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllRatingsForAMusic(musicID uuid.UUID) ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings WHERE music_id=?", musicID.String())
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	ratings := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id, &data.Comment, &data.Rating, &data.RatingDate, &data.MusicID, &data.UserID)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return ratings, err
}

func GetMusicRating(musicID uuid.UUID, ratingID uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM ratings WHERE music_id=? AND id=?", musicID.String(), ratingID.String())
	helpers.CloseDB(db)

	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.Comment, &rating.Rating, &rating.RatingDate, &rating.MusicID, &rating.UserID)
	if err != nil {
		return nil, err
	}
	return &rating, err
}

func AddMusicRating(newRating models.Rating) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO ratings (id, comment, rating, rating_date, music_id, user_id) VALUES (?, ?, ?, ?, ?, ?);", newRating.Id.String(), newRating.Comment, newRating.Rating, newRating.RatingDate, newRating.MusicID.String(), newRating.UserID.String())
	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return nil
}

func ModifyMusicRating(musicID uuid.UUID, ratingID uuid.UUID, newRatingData models.RatingRequest) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	if newRatingData.Rating != nil {
		_, err = tx.Exec("UPDATE ratings SET rating=? WHERE music_id=? AND id=?;", newRatingData.Rating, musicID.String(), ratingID.String())
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if newRatingData.Comment != nil {
		_, err = tx.Exec("UPDATE ratings SET comment=? WHERE music_id=? AND id=?;", newRatingData.Comment, musicID.String(), ratingID.String())
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if newRatingData.UserID != nil {
		_, err = tx.Exec("UPDATE ratings SET user_id=? WHERE music_id=? AND id=?;", newRatingData.UserID, musicID.String(), ratingID.String())
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	row := tx.QueryRow("SELECT * FROM ratings WHERE music_id=? AND id=?", musicID.String(), ratingID.String())
	var rating models.Rating
	err = row.Scan(&rating.Id, &rating.Comment, &rating.Rating, &rating.RatingDate, &rating.MusicID, &rating.UserID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	helpers.CloseDB(db)

	return &rating, err
}

func DeleteMusicRating(musicID uuid.UUID, ratingID uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM ratings WHERE music_id=? AND id=?", musicID.String(), ratingID.String())
	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return err
}
