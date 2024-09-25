package main

import (
	"chat-app/cmd"
	"chat-app/db"
)
func main() {
	db.DbConnect()

	cmd.Execute()

}
