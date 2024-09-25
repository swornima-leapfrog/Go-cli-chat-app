package cmd

import (
	"chat-app/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start the server",

	Run: func(cmd *cobra.Command, args []string) {
		handlers.StartServer()

	fmt.Println("Server started at port: 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))

	},
}
