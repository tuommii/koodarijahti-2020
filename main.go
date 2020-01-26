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
	NextPrize  int `json:"nextPrize"`
}

// Prizes
const (
	PrizeBig    = 500
	PrizeMedium = 100
	PrizeSmall  = 10
)

// Prize ...
type Prize map[int]int

// State represents game state
type State struct {
	Clicks    int
	Players   map[string]*Player
	NextPrize int
	Prizes    map[int]int
}

func (s *State) checkPrize(ip string) {
	if s.NextPrize > 1 {
		s.NextPrize--
		return
	}
	if s.Clicks%PrizeBig == 0 {
		s.Players[ip].Score += 250
	} else if s.Clicks%PrizeMedium == 0 {
		s.Players[ip].Score += 40
	} else if s.Clicks%PrizeSmall == 0 {
		s.Players[ip].Score += 5
	}
	s.NextPrize = 10
}

// secod arr like [::1]:49046
func parseIP(addr string) string {
	arr := strings.Split(addr, ":")
	// ip := strings.Split(arr[len(arr)-1], "]")
	return arr[0]
}

func (s *State) getState(w http.ResponseWriter, r *http.Request) {
	var p *Player
	ip := parseIP(r.RemoteAddr)
	if _, ok := s.Players[ip]; ok {
		// Player exist, take data form State
		p = &Player{
			Score:      s.Players[ip].Score,
			ClicksLeft: s.Players[ip].ClicksLeft,
			NextPrize:  s.Players[ip].NextPrize,
		}
	} else {
		p = &Player{Score: 0, ClicksLeft: 2000, NextPrize: s.NextPrize}
		s.Players[ip] = p
		s.Players[ip].NextPrize = s.NextPrize
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(s.Players[ip])
}

// Handle
func (s *State) handleAction(w http.ResponseWriter, r *http.Request) {
	ip := parseIP(r.RemoteAddr)

	if s.Players[ip].ClicksLeft == 0 {
		json.NewEncoder(w).Encode(s.Players[ip])
		return
	}

	s.Clicks++
	s.Players[ip].ClicksLeft--
	s.checkPrize(ip)
	s.Players[ip].NextPrize = s.NextPrize

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(s.Players[ip])
}

func main() {
	state := &State{Clicks: 0, NextPrize: 10}
	state.Players = make(map[string]*Player)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/action", state.handleAction)
	http.HandleFunc("/state", state.getState)
	http.Handle("/", fs)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
