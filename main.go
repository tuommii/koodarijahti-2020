package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

<<<<<<< HEAD
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

// TODO: Handle Reset
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
=======
	socketio "github.com/googollee/go-socket.io"
)

// Rulebook represents all game rule
type Rulebook struct {
	Counter int
	// [click_count]prize_points
	// prizes  map[int]int
	// players map[string]Player
>>>>>>> 244175d752c5984b647b6a43f57a1c3062c6aea1
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("ws://0.0.0.0:3000", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()
	fs := http.FileServer(http.Dir("./public"))
<<<<<<< HEAD
	http.HandleFunc("/inc", incrementCounter)
	http.HandleFunc("/state", getState)
	http.Handle("/", fs)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
=======
	// http.HandleFunc("/inc", rb.handleInc)
	// http.HandleFunc("/get", rb.handleGet)
	// http.HandleFunc("/ws", rb.wsEndpoint)
	http.Handle("/", fs)
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
>>>>>>> 244175d752c5984b647b6a43f57a1c3062c6aea1
	if err != nil {
		log.Fatal(err)
	}
}
