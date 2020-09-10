package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"game/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app *App

const initGameString string = `{"State":{"PlOnline":0,"PlDead":0,"PlAlive":0,"Players":{"admin-uniq-key":{"Nickname":"inzh","Score":0,"Online":false,"Alive":false}}}}`
const gameAfterUpdate string = `{"State":{"PlOnline":10,"PlDead":3,"PlAlive":2,"Players":{"admin-uniq-key":{"Nickname":"inzh","Score":0,"Online":false,"Alive":false}}}}`
const stateAfterUpdate string = `{"PlOnline":10,"PlDead":3,"PlAlive":2,"Players":{"admin-uniq-key":{"Nickname":"inzh","Score":0,"Online":false,"Alive":false}}}`
const reposneData string = `{"Message":"OK"}`

const initPlayersString string = `{"admin-uniq-key":{"Nickname":"inzh","Score":0,"Online":false,"Alive":false}}`

var jsonRequest = []byte(`{
	"pl_online": 10,
	"pl_dead": 3,
	"pl_alive": 2
}`)

func TestMain(t *testing.T) {
	config := config.GetConfig()

	app = &App{}
	app.Initialize(config)

}

func TestGetGame(t *testing.T) {

	req, _ := http.NewRequest("GET", "/game", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()

	if body != initGameString {
		t.Errorf("Initial state not good, lol. Got %s", body)
	}

}

func updateTest(t *testing.T) string {
	req, _ := http.NewRequest("POST", "/game", bytes.NewBuffer(jsonRequest))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	body := response.Body.String()

	return body

}

func TestUpdateGame(t *testing.T) {
	body := updateTest(t)
	if body != reposneData {
		t.Errorf(" Not goot updated data. Got %s", body)
	}
}

func TestUpdateAndGetAfterUpdate(t *testing.T) {
	body := updateTest(t)
	if body != reposneData {
		t.Errorf(" Not goot updated data. Got %s", body)
	}

	req, _ := http.NewRequest("GET", "/game", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	respAfterUpdate := response.Body.String()

	if respAfterUpdate != gameAfterUpdate {
		t.Errorf("Response after update not good, lol. Got %s", body)
	}

}

func TestGetAllPlayers(t *testing.T) {

	req, _ := http.NewRequest("GET", "/game/players", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()

	if body != initPlayersString {
		t.Errorf("Initial playesr not good, lol. Got %s", body)
	}

}

func TestHandleUpdateState(t *testing.T) {
	state := initState()

	payload := GameRequest{
		PlAlive:  2,
		PlDead:   3,
		PlOnline: 10,
	}

	newState := hanleUpdateState(state, payload)

	expectedStateJson, err := json.Marshal(newState)

	if err != nil {
		fmt.Println(err)
		t.Errorf("Error")

	}

	strExp := string(expectedStateJson)

	if strExp != stateAfterUpdate {
		t.Errorf("Expected not same")
	}

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}
