package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// handleGetState returns current state
func (gs *GameState) handleGetState(w http.ResponseWriter, r *http.Request) {
	var p *Player
	ip := gs.getIP(w, r)
	if _, ok := gs.Players[ip]; ok {
		// Player exist, copy data from GameState
		p = createPlayer(gs.Players[ip].Score, gs.Players[ip].ClicksLeft, gs.Players[ip].NextPrize)
	} else {
		// Add new player
		p = createPlayer(0, StartingPoints, gs.NextPrize)
		gs.Players[ip] = p
	}
	log.Println("/STATE:", gs.Players[ip], ip, gs.Clicks)
	json.NewEncoder(w).Encode(gs.Players[ip])
}

// Update's game state
func (gs *GameState) handleClick(w http.ResponseWriter, r *http.Request) {
	ip := gs.getIP(w, r)
	if gs.Players[ip].ClicksLeft == 0 {
		json.NewEncoder(w).Encode(gs.Players[ip])
		return
	}
	gs.update(ip)
	log.Println("/ACTION:", gs.Players[ip], ip, gs.Clicks)
	json.NewEncoder(w).Encode(gs.Players[ip])
}

// Reset player's data
func (gs *GameState) handleReset(w http.ResponseWriter, r *http.Request) {
	ip := gs.getIP(w, r)
	if gs.Players[ip].ClicksLeft == 0 {
		gs.Players[ip].ClicksLeft = StartingPoints
		gs.Players[ip].Score = 0
		gs.Players[ip].NextPrize = gs.NextPrize
		json.NewEncoder(w).Encode(gs.Players[ip])
		return
	}
	log.Println("/RESET:", gs.Players[ip], ip)
	json.NewEncoder(w).Encode(gs.Players[ip])
}
