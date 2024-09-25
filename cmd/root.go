package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chat-app",
	Short: "Chat App",
	Long:  "Chat App to connect with people who are logged in",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chat App")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func init(){
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(serverCmd)
}
