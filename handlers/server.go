package handlers

import (
	"chat-app/models"
	"chat-app/structs"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

var server *socketio.Server

func StartServer() {
	server = socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		queryParams, err := url.ParseQuery(s.URL().RawQuery)
		if err != nil {
			return fmt.Errorf("error parsing query parameters: %v", err)
		}
		userID := queryParams.Get("id")
		room := queryParams.Get("room")
		s.Join(room)
		fmt.Println("New user connected:", s.ID(), "with userID:", userID, "in room:", room)
		return nil
	})

	server.OnEvent("/", "chat_message", func(s socketio.Conn, msg string) {
		queryParams, err := url.ParseQuery(s.URL().RawQuery)
		if err != nil {
			log.Printf("Error parsing query: %v", err)
			return
		}

		stringID := queryParams.Get("id")

		// Convert the string ID to an int
		userID, _ := strconv.Atoi(stringID)
		userName := queryParams.Get("username")
		room := queryParams.Get("room")

		message := map[string]string{
			"username": userName,
			"content":  msg,
		}

		server.BroadcastToRoom("/", room, "chat_message", message)

		chat := structs.ChatHistory{
			SenderID:  userID,
			Message:   msg,
			CreatedAt: time.Now(),
		}

		err = models.SaveChatMessage(chat)
		if err != nil {
			log.Printf("Error saving chat message: %v", err)
		}
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("Socket.IO error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("User disconnected:", s.ID())
	})

	go server.Serve()

	http.Handle("/socket.io/", server)
}
