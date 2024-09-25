package cmd

import (
	"chat-app/handlers"
	"chat-app/models"
	"chat-app/structs"
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login user",
	Long:  "Login user using email and password",

	Run: func(cmd *cobra.Command, args []string) {
		var email, password string

		fmt.Println("Enter email: ")
		fmt.Scan(&email)

		fmt.Println("Enter password: ")
		fmt.Scan(&password)

		user := structs.LoginUser{
			Email:    email,
			Password: password,
		}

		existingUser, err := models.LoginUser(user)

		if err != nil {
			fmt.Println("Invalid email or password.")
		}

		handlers.StartClient(*existingUser)
	},
}
