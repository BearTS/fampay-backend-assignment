package cmd

import (
	"fmt"
	"time"

	"github.com/BearTS/fampay-backend-assignment/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	// Created a separate service to showcase how we can have multiple services running
	// Another way would be including the same in the main api service and running it inside a go routine
	// By running this in a separate service, we can have a separate config for this service and also have a separate log file
	// This also makes sure that if we scale the services, we can scale them independently
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
