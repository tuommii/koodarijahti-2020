package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// Player represents player
type Player struct {
	Name   string
	Points int
}

// Rulebook represents all game rule
type Rulebook struct {
	// [click_count]prize_points
	prizes  map[int]int
	players map[string]Player
}

// CreateRulebook ...
func CreateRulebook() Rulebook {
	rb := Rulebook{}
	rb.prizes = make(map[int]int)
	rb.players = make(map[string]Player)
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

func parseIP(addr string) string {
	arr := strings.Split(addr, " ")
	return arr[len(arr)-1]
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
	// 2020/01/25 12:01:52 10.12.5.12:50995
	log.Printf("%v\n", r.RemoteAddr)
	log.Println(parseIP(r.RemoteAddr))
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(resp)
	w.Write(json)
}

func main() {
	port := os.Getenv("PORT")
	// rb := CreateRulebook()
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
