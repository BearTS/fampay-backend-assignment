package cmd

import (
	"fmt"
	"os"

	api "github.com/BearTS/fampay-backend-assignment/api/cmd"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

// Execute - starts the CLI
func init() {
	cmd.AddCommand(api.RootCmd())
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error occured while starting the app: ", err)
		os.Exit(1)
	}
}
