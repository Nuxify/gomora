package interfaces

import (
	"github.com/kabaluyot/gomora/app/http/models"
)

type PlayerRepositoryInterface interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
