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
		// p.Points = gs.Players[ip].Points
		// p.NextPrize = gs.Players[ip].NextPrize
		// gs.Players[ip] = p
	} else {
		// Add new player
		p = createPlayer(StartingPoints, gs.NextPrize)
		gs.Players[ip] = p
	}
	log.Println("/STATE:", gs.Players[ip], ip, gs.Clicks)
	json.NewEncoder(w).Encode(gs.Players[ip])
}

// Update's game state
func (gs *GameState) handleClick(w http.ResponseWriter, r *http.Request) {
	ip := gs.getIP(w, r)
	p := gs.Players[ip]
	log.Printf("MIIKKA: %+v, %+v, %s", p, gs.Players, ip)
	if gs.Players[ip] != nil && gs.Players[ip].Points == 0 {
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
	if gs.Players[ip].Points == 0 {
		gs.Players[ip].Points = StartingPoints
		gs.Players[ip].NextPrize = gs.NextPrize
		json.NewEncoder(w).Encode(gs.Players[ip])
		return
	}
	log.Println("/RESET:", gs.Players[ip], ip)
	json.NewEncoder(w).Encode(gs.Players[ip])
}
