package interfaces

import (
	"github.com/kabaluyot/googol/app/http/models"
)

type PlayerRepositoryInterface interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
