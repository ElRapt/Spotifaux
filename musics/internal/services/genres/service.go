package genres

import (
	"database/sql"
	"errors"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/genres"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
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

func PutGenre(partialGenre models.Genre) (models.Genre, error) {

	var genre models.Genre
	var err2 error

	genre, err2 = repository.PutGenre(partialGenre.Id, partialGenre.Name)

	if err2 != nil {
		logrus.Errorf("error updating Genre: %s", err2.Error())
		return models.Genre{}, &models.CustomError{
			Message: "Error updating Genre",
			Code:    http.StatusInternalServerError,
		}
	}

	return genre, nil
}
