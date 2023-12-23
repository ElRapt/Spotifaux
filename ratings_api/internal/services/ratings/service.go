package ratings

import (
	"database/sql"
	"errors"
	"middleware/example3/internal/models"
	repository "middleware/example3/internal/repositories/ratings"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllRatings() ([]models.Rating, error) {
	ratings, err := repository.GetAllRatings()
	if err != nil {
		logrus.Errorf("error retrieving ratings: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return ratings, nil
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	rating, err := repository.GetRatingById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "Rating not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving rating: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return rating, nil
}

func PostRating(newRating models.Rating) (uuid.UUID, error) {
	ratingId, err := repository.PostRating(newRating)
	if err != nil {
		logrus.Errorf("error posting rating: %s", err.Error())
		return uuid.Nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return ratingId, nil
}

func PutRating(id uuid.UUID, newRating models.Rating) error {
	err := repository.PutRating(id, newRating)
	if err != nil {
		logrus.Errorf("error updating rating: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func DeleteRating(id uuid.UUID) error {
	err := repository.DeleteRating(id)
	if err != nil {
		logrus.Errorf("error deleting rating: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}
