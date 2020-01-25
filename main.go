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
	Score      int `json:"score"`
	ClicksLeft int `json:"clicksLeft"`
}

var gClicks int
var gPlayers map[string]*Player

// [::1]:49046
func parseIP(addr string) string {
	arr := strings.Split(addr, " ")
	ip := strings.Split(arr[len(arr)-1], "]")
	return ip[0]
}

func init() {
	gPlayers = make(map[string]*Player)
	gClicks = 0
}

func getState(w http.ResponseWriter, r *http.Request) {
	ip := parseIP(r.RemoteAddr)
	if _, ok := gPlayers[ip]; ok {
		// Player exist
	} else {
		p := &Player{Score: 0, ClicksLeft: 20}
		gPlayers[ip] = p
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println(gPlayers[ip])
	json.NewEncoder(w).Encode(gPlayers)
}

// Handle REe
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	ip := parseIP(r.RemoteAddr)
	// Check prizes
	if (gClicks+1)%10 == 0 {
		gPlayers[ip].Score = 10
	}
	gPlayers[ip].ClicksLeft--
	gClicks++
	log.Println(ip, gPlayers[ip])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gPlayers)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/inc", incrementCounter)
	http.HandleFunc("/state", getState)
	http.Handle("/", fs)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
