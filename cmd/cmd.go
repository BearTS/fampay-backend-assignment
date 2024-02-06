package cmd

import (
	"fmt"
	"os"

	api "github.com/BearTS/fampay-backend-assignment/api/cmd"
	dbapp "github.com/BearTS/fampay-backend-assignment/dbapp/cmd"
	youtubeFetcher "github.com/BearTS/fampay-backend-assignment/youtube_fetcher/cmd"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute - starts the CLI
func init() {
	cmd.AddCommand(api.RootCmd())
	cmd.AddCommand(youtubeFetcher.RootCmd())
	cmd.AddCommand(dbapp.RootCmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error occured while starting the app: ", err)
		os.Exit(1)
	}
}
