package artists

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
)

func GetAllArtists() ([]models.Artist, error) {
	var err error
	artists, err := repository.GetAllArtists()
	if err != nil {
		logrus.Errorf("error retrieving artists : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return artists, nil
}

func GetArtistById(id uuid.UUID) (*models.Artist, error) {
	artist, err := repository.GetArtistById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "artist not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving artists : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return artist, err
}

func PostArtist(newArtist models.Artist) (uuid.UUID, error) {
	artistId, err := repository.PostArtist(newArtist)
	if err != nil {
		logrus.Errorf("error posting artist: %s", err.Error())
		return uuid.Nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return artistId, nil
}

func PutArtist(id uuid.UUID, newArtist models.Artist) error {
	err := repository.PutArtist(id, newArtist)
	if err != nil {
		logrus.Errorf("error putting artist: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}
