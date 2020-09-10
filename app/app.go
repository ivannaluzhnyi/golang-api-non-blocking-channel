package app

import (
	"game/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func channelsListener(channels *Channels, game *Game) {
	for {
		select {
		case newState := <-channels.updateStateChan:
			game.State = newState
		}
	}
}

func initState() State {
	return State{
		PlAlive:  0,
		PlDead:   0,
		PlOnline: 0,
		Players: map[string]Player{
			"admin-uniq-key": Player{
				Score:    0,
				Nickname: "inzh",
				Online:   false,
				Alive:    false,
			},
		},
	}
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	var channels *Channels
	var game *Game
	channels = &Channels{make(chan State)}
	game = &Game{State: initState()}

	go channelsListener(channels, game)
	a.Router = mux.NewRouter()
	a.setRouters(channels, game)

}

// setRouters sets the all required routers
func (a *App) setRouters(channels *Channels, game *Game) {
	// Routing for handling the projects
	a.Get("/game", a.handleRequest(GetGame, channels, game))
	a.Post("/game", a.handleRequest(UpdateGame, channels, game))

	a.Get("/game/players", a.handleRequest(GetAllPlayers, channels, game))
	a.Post("/game/players", a.handleRequest(AddNewPlayer, channels, game))

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(channels *Channels, game *Game, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction, channels *Channels, game *Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(channels, game, w, r)
	}
}
