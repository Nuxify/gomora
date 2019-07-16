package main

import (
	"sync"

	"database/sql"
	"github.com/kabaluyot/gomora/app/http/controllers"
	"github.com/kabaluyot/gomora/app/http/repositories"
	"github.com/kabaluyot/gomora/app/http/services"
	"github.com/kabaluyot/gomora/infrastructures"
)

//ServiceContainerInterface - contains the interfaces for the service container
type ServiceContainerInterface interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

//region dependency injection
func (k *kernel) InjectPlayerController() controllers.PlayerController {

	sqlConn, _ := sql.Open("sqlite3", "/var/tmp/tennis.db")
	sqliteHandler := &infrastructures.SQLiteHandler{}
	sqliteHandler.Conn = sqlConn

	playerRepository := &repositories.PlayerRepository{sqliteHandler}
	playerService := &services.PlayerService{&repositories.PlayerRepositoryWithCircuitBreaker{playerRepository}}
	playerController := controllers.PlayerController{playerService}

	return playerController
}

//endregion dependency injection

//ServiceContainer - export instantiated service container once
func ServiceContainer() ServiceContainerInterface {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
