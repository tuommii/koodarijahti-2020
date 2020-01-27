package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
	log.Println("/STATE", "IP:", ip, "Player:", gs.Players[ip], "Clicks:", gs.Clicks, len(gs.Players))
	json.NewEncoder(w).Encode(gs.Players[ip])
}

// Update's game state
func (gs *GameState) handleClick(w http.ResponseWriter, r *http.Request) {
	ip := gs.getIP(r)
	if gs.Players[ip].Points == 0 {
		json.NewEncoder(w).Encode(gs.Players[ip])
		return
	}
	gs.update(ip)
	log.Println("/CLICK\t", "IP:", ip, "Player:", gs.Players[ip], "Clicks:", gs.Clicks)
	json.NewEncoder(w).Encode(gs.Players[ip])
}

// Reset player's data
func (gs *GameState) handleReset(w http.ResponseWriter, r *http.Request) {
	ip := gs.getIP(r)
	if gs.Players[ip].Points == 0 {
		gs.Players[ip].Points = StartingPoints
		gs.Players[ip].NextPrize = gs.NextPrize
		json.NewEncoder(w).Encode(gs.Players[ip])
		return
	}
	log.Println("/RESET\t", "IP:", ip, "Player:", gs.Players[ip], "Clicks:", gs.Clicks)
	json.NewEncoder(w).Encode(gs.Players[ip])
}
