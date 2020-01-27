package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

// Player represents a player and contains all data that are sended to client
type Player struct {
	Score      int `json:"score"`
	ClicksLeft int `json:"clicksLeft"`
	NextPrize  int `json:"nextPrize,omitempty"`
}

// GameState represents game state
type GameState struct {
	// Total clicks
	Clicks int
	// Hold's all players. Key is IP like: ["127.0.0.1"]
	Players map[string]*Player
	// How many clicks to win next prize
	NextPrize int
	// production/development, maybe others also
	Env string
	// Port to listen
	Port string
}

// Clicks required to win prizes, and their values
const (
	PrizeBigClicks    = 500
	PrizeMediumClicks = 100
	PrizeSmallClicks  = 10
	PrizeBig          = 250
	PrizeMedium       = 40
	PrizeSmall        = 5
)

// StartingPoints for player
const StartingPoints = 20

// Create new State
func createGameState() *GameState {
	gs := &GameState{Clicks: 0, NextPrize: PrizeSmallClicks, Env: "dev"}
	gs.Players = make(map[string]*Player)
	gs.Port = os.Getenv("PORT")
	if gs.Port == "" {
		gs.Port = "3000"
	} else {
		gs.Env = "production"
	}

	return gs
}

// Create new Player
func createPlayer(score int, clicksLeft int, nextPrize int) *Player {
	p := &Player{Score: score, ClicksLeft: clicksLeft, NextPrize: nextPrize}
	return p
}

// Get amount of prize
func (gs *GameState) getPrize() int {
	if gs.Clicks%PrizeBigClicks == 0 {
		gs.Clicks = 0
		return PrizeBig
	} else if gs.Clicks%PrizeMediumClicks == 0 {
		return PrizeMedium
	} else if gs.Clicks%PrizeSmallClicks == 0 {
		return PrizeSmall
	}
	return 0
}

// Check if prize is won, and adds prize's score to player identified by ip
// Only every 10:nth (smallest count for prize) need's to be checked
func (gs *GameState) checkPrize(ip string) {
	if gs.NextPrize > 1 {
		gs.NextPrize--
		return
	}
	gs.Players[ip].Score += gs.getPrize()
	gs.NextPrize = PrizeSmallClicks
}

// Update game state
func (gs *GameState) update(ip string) {
	gs.Clicks++
	gs.Players[ip].ClicksLeft--
	gs.checkPrize(ip)
	gs.Players[ip].NextPrize = gs.NextPrize
}

// Set headers according to ENV
func (gs *GameState) setHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if gs.Env != "production" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

// Parse ip if ENV is dev
func parseIP(addr string) string {
	arr := strings.Split(addr, ":")
	return arr[0]
}

// Get right ip according to ENV
func (gs *GameState) getIP(w http.ResponseWriter, r *http.Request) string {
	var ip string
	if gs.Env != "production" {
		ip = parseIP(r.RemoteAddr)
	} else {
		ip = r.Header.Get("X-Forwarded-For")
	}
	return ip
}

func main() {
	state := createGameState()
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/click", state.handleClick)
	http.HandleFunc("/state", state.handleGetState)
	http.HandleFunc("/reset", state.handleReset)
	http.Handle("/", fs)
	log.Println("Server started on", state.Env)
	err := http.ListenAndServe("0.0.0.0:"+state.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
