package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Middleware logger, called with every request
func (gs *GameState) logger(next http.HandlerFunc) http.HandlerFunc {
	var ip string
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		ip = gs.getIP(r)
		log.Println(r.URL, "IP:", ip, "Player:", gs.Players[ip], "Clicks:", gs.Clicks, "Players:", len(gs.Players))
	}
}

// handleGetState returns current state
func (gs *GameState) handleGetState(w http.ResponseWriter, r *http.Request) {
	var p *Player
	ip := gs.getIP(r)
	if _, ok := gs.Players[ip]; ok {
		// Player exist
	} else {
		// Add new player
		p = createPlayer(StartingPoints, gs.NextPrize)
		gs.Players[ip] = p
	}
	json.NewEncoder(w).Encode(gs.Players[ip])
}

// Update's game state
func (gs *GameState) handleClick(w http.ResponseWriter, r *http.Request) {
	ip := gs.getIP(r)
	if gs.Players[ip].Points == 0 {
		json.NewEncoder(w).Encode(gs.Players[ip])
	} else {
		gs.update(ip)
		json.NewEncoder(w).Encode(gs.Players[ip])
	}
}

// Reset player's data
func (gs *GameState) handleReset(w http.ResponseWriter, r *http.Request) {
	ip := gs.getIP(r)
	if gs.Players[ip].Points == 0 {
		gs.resetPlayer(ip)
		json.NewEncoder(w).Encode(gs.Players[ip])
	} else {
		json.NewEncoder(w).Encode(gs.Players[ip])
	}
}
