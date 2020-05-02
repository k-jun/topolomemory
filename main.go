package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
)

var Server socketio.Server

func main() {
	Server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	EmitStatus := func() {
		status := "user number: " + strconv.Itoa(Server.RoomLen("/", "default"))
		Server.BroadcastToRoom("/", "default", "status", status)
	}

	// room/index
	Server.OnEvent("/", "room/index", func(s socketio.Conn, payload string) error {
		fmt.Println("room/index")
		EmitStatus()
		return nil
	})

	// room/create
	Server.OnEvent("/", "room/create", func(s socketio.Conn, payload string) error {
		fmt.Println("room/create")
		fmt.Println(Server.Rooms("/"))
		EmitStatus()
		return nil
	})

	// room/join
	Server.OnEvent("/", "room/join", func(s socketio.Conn, payload string) error {
		fmt.Println("room/join")
		return nil
	})

	// room/leave
	Server.OnEvent("/", "room/leave", func(s socketio.Conn, payload string) error {
		fmt.Println("room/leave")
		return nil
	})

	Server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		Server.JoinRoom("/", "default", s)
		return nil

	})
	Server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
		s.LeaveAll()
	})
	Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		s.LeaveAll()
	})
	go Server.Serve()
	defer Server.Close()
	http.Handle("/socket.io/", Server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Status() {
}
