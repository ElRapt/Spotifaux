package musics

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
)

func GetAllMusics() ([]models.Music, error) {
	musics, err := repository.GetAllMusics()
	if err != nil {
		logrus.Errorf("error retrieving musics: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return musics, nil
}

func GetMusicById(id uuid.UUID) (*models.Music, error) {
	music, err := repository.GetMusicById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "Music not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving music: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return music, nil
}

func PostMusic(newMusic models.Music) (uuid.UUID, error) {
	musicId, err := repository.PostMusic(newMusic)
	if err != nil {
		logrus.Errorf("error posting music: %s", err.Error())
		return uuid.Nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return musicId, nil
}

func PutMusic(id uuid.UUID, newMusic models.Music) error {
	err := repository.PutMusic(id, newMusic)
	if err != nil {
		logrus.Errorf("error putting music: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func DeleteMusic(id uuid.UUID) error {
	err := repository.DeleteMusic(id)
	if err != nil {
		logrus.Errorf("error deleting music: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}
