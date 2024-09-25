package structs

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUser struct{
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}
