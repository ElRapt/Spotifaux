package albums

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/albums"
	"net/http"
)

func GetAllAlbums() ([]models.Album, error) {
	var err error
	albums, err := repository.GetAllAlbums()
	if err != nil {
		logrus.Errorf("error retrieving albums : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return albums, nil
}

func GetAlbumById(id uuid.UUID) (*models.Album, error) {
	album, err := repository.GetAlbumById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "album not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving albums : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return album, err
}

func PostAlbum(newAlbum models.Album) (uuid.UUID, error) {
	albumId, err := repository.PostAlbum(newAlbum)
	if err != nil {
		logrus.Errorf("error posting album: %s", err.Error())
		return uuid.Nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return albumId, nil
}

func PutAlbum(id uuid.UUID, newAlbum models.Album) error {
	err := repository.PutAlbum(id, newAlbum)
	if err != nil {
		logrus.Errorf("error putting album: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}
