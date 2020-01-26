package main

import (
	"encoding/json"
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

// State represents game state
type State struct {
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
func createState() *State {
	state := &State{Clicks: 0, NextPrize: PrizeSmallClicks, Env: "dev"}
	state.Players = make(map[string]*Player)
	state.Port = os.Getenv("PORT")
	if state.Port == "" {
		state.Port = "3000"
	} else {
		state.Env = "production"
	}

	return state
}

// Create new Player
func createPlayer(score int, clicksLeft int, nextPrize int) *Player {
	p := &Player{Score: score, ClicksLeft: clicksLeft, NextPrize: nextPrize}
	return p
}

// Get amount of prize
func (s *State) getPrize() int {
	if s.Clicks%PrizeBigClicks == 0 {
		return PrizeBig
	} else if s.Clicks%PrizeMediumClicks == 0 {
		return PrizeMedium
	} else if s.Clicks%PrizeSmallClicks == 0 {
		// s.Players[ip].Score += PrizeSmall
		return PrizeSmall
	}
	return 0
}

// Check if prize is won, and adds prize's score to player identified by ip
// Only every 10:nth (smallest count for prize) need's to be checked
func (s *State) checkPrize(ip string) {
	if s.NextPrize > 1 {
		s.NextPrize--
		return
	}
	s.Players[ip].Score += s.getPrize()
	s.NextPrize = PrizeSmallClicks
}

// Set headers according to ENV
func (s *State) setHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if s.Env != "production" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

// Parse ip if ENV is dev
func parseIP(addr string) string {
	arr := strings.Split(addr, ":")
	return arr[0]
}

// Get right ip according to ENV
func (s *State) getIP(w http.ResponseWriter, r *http.Request) string {
	var ip string
	if s.Env != "production" {
		ip = parseIP(r.RemoteAddr)
	} else {
		ip = r.Header.Get("X-Forwarded-For")
	}
	return ip
}

// Called when page is refreshed
func (s *State) getState(w http.ResponseWriter, r *http.Request) {
	var p *Player
	ip := s.getIP(w, r)
	if _, ok := s.Players[ip]; ok {
		// Player exist, copy data from State
		p = createPlayer(s.Players[ip].Score, s.Players[ip].ClicksLeft, s.Players[ip].NextPrize)
	} else {
		// Add new player
		p = createPlayer(0, StartingPoints, s.NextPrize)
		s.Players[ip] = p
	}
	log.Println("/STATE:", s.Players[ip], ip)
	json.NewEncoder(w).Encode(s.Players[ip])
}

// Update game state
func (s *State) update(ip string) {
	s.Clicks++
	s.Players[ip].ClicksLeft--
	s.checkPrize(ip)
	s.Players[ip].NextPrize = s.NextPrize
}

// Called when button is clicked
func (s *State) handleClick(w http.ResponseWriter, r *http.Request) {
	ip := s.getIP(w, r)
	if s.Players[ip].ClicksLeft == 0 {
		json.NewEncoder(w).Encode(s.Players[ip])
		return
	}
	s.update(ip)
	log.Println("/ACTION:", s.Players[ip], ip)
	json.NewEncoder(w).Encode(s.Players[ip])
}

func main() {
	state := createState()
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/click", state.handleClick)
	http.HandleFunc("/state", state.getState)
	http.Handle("/", fs)
	log.Println("Server started on", state.Env)
	err := http.ListenAndServe("0.0.0.0:"+state.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
