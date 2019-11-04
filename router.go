package main

import (
	"sync"

	"github.com/go-chi/chi"
)

//IChiRouter - chi router interface
type IChiRouter interface {
	InitRouter() *chi.Mux
}

//create router struct
type router struct{}

var (
	m          *router
	routerOnce sync.Once
)

//Initialize router and routes
func (router *router) InitRouter() *chi.Mux {

	playerController := ServiceContainer().RegisterPlayerController()

	r := chi.NewRouter()
	r.HandleFunc("/getScore/{player1}/vs/{player2}", playerController.GetPlayerScore)

	return r
}

//ChiRouter - export instantiated chi router once
func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
