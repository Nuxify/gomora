package interfaces

import (
	"github.com/kabaluyot/googol/app/http/models"
)

type IPlayerRepository interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
