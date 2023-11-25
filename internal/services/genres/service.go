package genres

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
)

func GetAllGenres() ([]models.Genre, error) {
	var err error
	genres, err := repository.GetAllGenres()
	if err != nil {
		logrus.Errorf("error retrieving genres : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return genres, nil
}

func GetGenreById(id uuid.UUID) (*models.Genre, error) {
	genre, err := repository.GetGenreById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "genre not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving genres : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return genre, err
}

func PostGenre(newGenre models.Genre) (uuid.UUID, error) {
	genreId, err := repository.PostGenre(newGenre)
	if err != nil {
		logrus.Errorf("error posting genre: %s", err.Error())
		return uuid.Nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return genreId, nil
}

func PutGenre(id uuid.UUID, newGenre models.Genre) error {
	err := repository.PutGenre(id, newGenre)
	if err != nil {
		logrus.Errorf("error putting genre: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}
