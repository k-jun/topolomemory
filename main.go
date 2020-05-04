package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"topolomemory/models"

	socketio "github.com/googollee/go-socket.io"
)

var Server socketio.Server

type RoomIndex struct {
	Rooms []string
}
type GameStatus struct {
	Field []models.Card
	Score string
}
type GameScore struct {
	ScoreBoard string
}

func main() {
	Game := models.NewGame()
	Score := models.NewScore()
	Game.Start()
	Server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	BroadcastStatus := func() error {
		board, err := Score.Status()
		gameStatus := GameStatus{
			Field: Game.Status(),
			Score: board,
		}
		encodeData, err := json.Marshal(gameStatus)

		if err != nil {
			fmt.Println("err", err)
			return err
		}
		Server.BroadcastToRoom("/", "all", "game/status", string(encodeData))
		return nil
	}

	Server.OnEvent("/", "game/start", func(s socketio.Conn) error {
		fmt.Println("game/start")
		Game.Start()
		Score.Reset()
		BroadcastStatus()
		return nil
	})

	Server.OnEvent("/", "game/draw", func(s socketio.Conn) error {
		fmt.Println("game/draw")
		err := Game.Draw()
		if err != nil {
			fmt.Println("err", err)
			return err
		}
		err = BroadcastStatus()
		if err != nil {
			return err
		}
		return nil
	})
	Server.OnEvent("/", "game/hit", func(s socketio.Conn, msg string) error {
		fmt.Println("game/hit")
		ids := strings.Split(msg, ",")
		aID, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Println("err", err)
			return err
		}
		bID, err := strconv.Atoi(ids[1])
		if err != nil {
			fmt.Println("err", err)
			return err
		}
		IsGetCard, err := Game.Hit(aID, bID)
		if err != nil {
			fmt.Println("err", err)
			return err
		}

		if IsGetCard {
			Score.GetPoint(s.ID())
			s.Emit("game/win", "")
			err = BroadcastStatus()
		} else {
			s.Emit("game/lose", "")
		}

		if err != nil {
			fmt.Println("err", err)
			return err
		}

		return nil
	})

	Server.OnEvent("/", "username/change", func(s socketio.Conn, payload string) {
		fmt.Println("username/change")
		Score.SetUserName(s.ID(), payload)
		BroadcastStatus()
		s.Emit("username/change", Score.GetUserName(s.ID()))
	})

	Server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		Server.JoinRoom("/", "all", s)
		Score.SetUserName(s.ID(), "player-"+s.ID())
		s.Emit("username/change", Score.GetUserName(s.ID()))
		return nil
	})
	Server.OnError("/", func(s socketio.Conn, e error) {
		Score.ClearUserName(s.ID())
		s.LeaveAll()
	})
	Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		Score.ClearUserName(s.ID())
		s.LeaveAll()
	})
	go Server.Serve()
	defer Server.Close()
	http.Handle("/socket.io/", Server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
