package handlers

import (
	"bufio"
	"chat-app/structs"
	"fmt"
	"log"
	"os"
	"strconv"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func StartClient(user structs.User) {
	serverHost := os.Getenv("DB_HOST")
	serverPort := os.Getenv("PORT")

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	opts.Query["id"] = strconv.Itoa(user.ID)
	opts.Query["username"] = user.Username
	opts.Query["password"] = user.Password
	opts.Query["room"] = "chatApp"

	uri := fmt.Sprintf("http://%v:%v/socket.io/", serverHost, serverPort)
	
	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("Error Occured \n")
	})

	client.On("connection", func() {
		log.Printf("Connected to the server \n")
	})

	client.On("chat_message", func(msg map[string]interface{}) {
		userName := msg["username"].(string)
		content := msg["content"].(string)
		log.Printf("%s: %s\n", userName, content)
		fmt.Print("Send Message: ")
	})

	client.On("disconnection", func() {
		log.Printf("Server Disconnected\n")
	})

	reader := bufio.NewReader(os.Stdin)

	for {
		data, _, _ := reader.ReadLine()
		command := string(data)

		if command != "" {
			client.Emit("chat_message", command)
		}
	}
}
