package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	socketio "github.com/googollee/go-socket.io"
)

// Rulebook represents all game rule
type Rulebook struct {
	Counter int
	// [click_count]prize_points
	// prizes  map[int]int
	// players map[string]Player
}

func main() {
	port := os.Getenv("PORT")
	// rb := CreateRulebook()
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
	// http.HandleFunc("/inc", rb.handleInc)
	// http.HandleFunc("/get", rb.handleGet)
	// http.HandleFunc("/ws", rb.wsEndpoint)
	http.Handle("/", fs)
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
