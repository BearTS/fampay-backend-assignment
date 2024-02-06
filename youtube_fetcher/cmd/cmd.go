package cmd

import (
	"fmt"
	"time"

	"github.com/BearTS/fampay-backend-assignment/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "youtube-fetcher",
		Short: "fetches data from youtube every x seconds",
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.Info("youtube-fetcher running")

			fetch := func() {
				// TODO: Fetch data from youtube using the youtube pkg
				fmt.Println("Fetching data from youtube")
			}

			fetch()

			fmt.Println(config.Config.YoutubeFetchInterval, " is the interval of seconds")

			// Run the fetch function every x seconds
			ticker := time.NewTicker(time.Duration(config.Config.YoutubeFetchInterval) * time.Second)
			go func() {
				for range ticker.C {
					fetch()
				}
			}()

			// Keep the main goroutine running
			select {}
		},
	}

	return c
}
