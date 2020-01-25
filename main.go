package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Player represents player
type Player struct {
	Name   string
	Points int
}

// Rulebook represents all game rule
type Rulebook struct {
	// [click_count]prize_points
	prizes map[int]int
}

// CreateRulebook ...
func CreateRulebook() Rulebook {
	rb := Rulebook{}
	rb.prizes = make(map[int]int)
	rb.prizes[10] = 5
	rb.prizes[100] = 40
	rb.prizes[250] = 500
	return rb
}

// CreatePlayer creates new player
func CreatePlayer() Player {
	// TODO: Maybe static int player_count, and add to it player name
	p := Player{Points: 20, Name: "Default"}
	return p
}

func dataHandle(w http.ResponseWriter, r *http.Request) {
	// response := struct {
	// 	Players []model.Player
	// 	Games   []model.Game
	// 	Message string
	// }{
	// 	players,
	// 	games,
	// 	msg,
	// }
	//s.templates["home.html"].ExecuteTemplate(w, "base", response)
	resp := struct {
		Player
	}{
		CreatePlayer(),
	}
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(resp)
	w.Write(json)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/data", dataHandle)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
