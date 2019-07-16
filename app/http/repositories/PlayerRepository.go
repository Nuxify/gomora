package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/kabaluyot/gomora/app/http/interfaces"
	"github.com/kabaluyot/gomora/app/http/models"

	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepositoryWithCircuitBreaker struct {
	PlayerRepository interfaces.PlayerRepositoryInterface
}

func (repository *PlayerRepositoryWithCircuitBreaker) GetPlayerByName(name string) (models.PlayerModel, error) {

	output := make(chan models.PlayerModel, 1)
	hystrix.ConfigureCommand("get_player_by_name", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_player_by_name", func() error {

		player, _ := repository.PlayerRepository.GetPlayerByName(name)

		output <- player
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.PlayerModel{}, err
	}
}

type PlayerRepository struct {
	interfaces.DBHandlerInterface
}

func (repository *PlayerRepository) GetPlayerByName(name string) (models.PlayerModel, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM player_models WHERE name = '%s'", name))
	if err != nil {
		return models.PlayerModel{}, err
	}

	var player models.PlayerModel

	row.Next()
	row.Scan(&player.Id, &player.Name, &player.Score)

	return player, nil
}
