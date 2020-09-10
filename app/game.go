package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func hanleUpdateState(state State, payload GameRequest) State {
	return State{
		PlAlive:  payload.PlAlive,
		PlDead:   payload.PlDead,
		PlOnline: payload.PlOnline,
		Players:  state.Players,
	}
}

func hanleUpdatePlayer(state State, payload PlayerRequest) State {
	newPlayers := make(map[string]Player)
	for k, v := range state.Players {
		newPlayers[k] = v
	}

	u1 := uuid.NewV4().String()

	newPlayers[u1] = Player{
		Alive:    payload.Alive,
		Nickname: payload.Nickname,
		Online:   payload.Online,
		Score:    payload.Score,
	}

	return State{
		PlAlive:  state.PlAlive,
		PlDead:   state.PlDead,
		PlOnline: state.PlOnline,
		Players:  newPlayers,
	}
}

func editState(channels *Channels, newState State) {
	channels.updateStateChan <- newState
}

// GetGame => get game
func GetGame(channels *Channels, game *Game, w http.ResponseWriter, r *http.Request) {
	fmt.Println("API : GET /game ")
	respondJSON(w, http.StatusOK, game)
}

// UpdateGame update props
func UpdateGame(channels *Channels, game *Game, w http.ResponseWriter, r *http.Request) {
	fmt.Println("API : POST /game ")
	var payload GameRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	newState := hanleUpdateState(game.State, payload)
	go editState(channels, newState)

	respondJSON(w, http.StatusCreated, Response{
		Message: "OK",
	})
}

// GetAllPlayers => get all players
func GetAllPlayers(channels *Channels, game *Game, w http.ResponseWriter, r *http.Request) {
	fmt.Println("API  : GET /game/players ")
	respondJSON(w, http.StatusOK, game.State.Players)
}

// AddNewPlayer => AddNewPlayer
func AddNewPlayer(channels *Channels, game *Game, w http.ResponseWriter, r *http.Request) {
	fmt.Println("API  : POST /game/players ")
	var payload PlayerRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	newState := hanleUpdatePlayer(game.State, payload)

	go editState(channels, newState)

	respondJSON(w, http.StatusCreated, Response{
		Message: "OK",
	})
}
