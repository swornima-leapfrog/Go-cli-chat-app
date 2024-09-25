package models

import (
	"chat-app/db"
	"chat-app/structs"
)

func SaveChatMessage(chatHistory structs.ChatHistory) error {
	dbConn := db.DbConnect()
	defer dbConn.Close()

	_, err := dbConn.Model(&chatHistory).Insert()

	if err != nil {
		return err
	}

	return nil
}
