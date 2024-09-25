package cmd

import (
	"chat-app/models"
	"chat-app/structs"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register user",
	Long:  "Register user using name, email, password",

	Run: func(cmd *cobra.Command, args []string) {
		var username, email, password string

		fmt.Println("Enter username: ")
		fmt.Scan(&username)

		fmt.Println("Enter email: ")
		fmt.Scan(&email)

		fmt.Println("Enter password: ")
		fmt.Scan(&password)

		user := structs.User{
			Username: username,
			Email:    email,
			Password: password,
			CreatedAt: time.Now(),
		}

		err := models.RegisterUser(user)
		
		if err != nil {
			fmt.Print("Could not register user: ", err)
		}else{

			fmt.Print("User registered successfully")
		}

	},
}
