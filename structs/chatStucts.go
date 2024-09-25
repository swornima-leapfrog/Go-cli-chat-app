package structs

import "time"

type ChatHistory struct {
	ID        int       `json:"id"`
	SenderID  int       `json:"sender_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
