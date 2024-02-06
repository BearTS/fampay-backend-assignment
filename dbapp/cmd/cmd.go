package cmd

import (
	"fmt"

	"github.com/BearTS/fampay-backend-assignment/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "dbapp",
	Short: "run different db operations",
	Run: func(cmd *cobra.Command, args []string) {
		// Current environment
		fmt.Println("Current Environment: ", config.Config.Environment)
	},
}

func init() {
	RootCmd.AddCommand(Migrate())
}
