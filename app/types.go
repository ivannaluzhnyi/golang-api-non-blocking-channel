package app

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
	Game   Game
}

type Game struct {
	State State `json:state`
}

type Channels struct {
	updateStateChan chan State
}

type State struct {
	PlOnline int
	PlDead   int
	PlAlive  int
	Players  map[string]Player
}

type Player struct {
	Nickname string
	Score    int
	Online   bool
	Alive    bool
}

type GameRequest struct {
	PlOnline int `json:"pl_online"`
	PlDead   int `json:"pl_dead"`
	PlAlive  int `json:"pl_alive"`
}

type Response struct {
	Message string
}

type PlayerRequest struct {
	Nickname string `json:"nickname"`
	Score    int    `json:"score"`
	Online   bool   `json:"online"`
	Alive    bool   `json:"alive"`
}
