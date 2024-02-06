package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/BearTS/fampay-backend-assignment/pkg/config"
	"github.com/BearTS/fampay-backend-assignment/pkg/db"
	"github.com/BearTS/fampay-backend-assignment/pkg/youtube"
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
			apiKey := strings.Split(config.Config.YoutubeApiKeys, ",")
			gormDB, _ := db.Connection()
			database := db.NewDB(gormDB)
			youtubeSvc := youtube.NewYoutubeClient(apiKey)
			timeafter := time.Now().AddDate(0, -1, 0)

			// Fetch the data from youtube
			fetch := func() {
				fmt.Println("Fetching data from youtube")
				data, err := youtubeSvc.SearchByQuery(config.Config.YoutubeQuery, 10, timeafter)
				if err != nil {
					logrus.Error("Error fetching data from youtube: ", err)
					return
				}
				var dbVideos []*db.Videos
				for _, video := range data {
					fmt.Println("Title: ", video.Snippet.Title)
					fmt.Println("Description: ", video.Snippet.Description)
					fmt.Println("PublishedAt: ", video.Snippet.PublishedAt)
					fmt.Println("Thumbnail: ", video.Snippet.Thumbnails.Default.Url)
					fmt.Println("VideoId: ", video.Id.VideoId)
					fmt.Println()

					var dbVideo db.Videos
					dbVideo.VideoId = video.Id.VideoId
					dbVideo.Title = video.Snippet.Title
					dbVideo.Description = video.Snippet.Description
					dbVideo.Thumbnail = video.Snippet.Thumbnails.Default.Url
					dbVideo.ChannelTitle = video.Snippet.ChannelTitle
					timeParsed, _ := time.Parse(time.RFC3339, video.Snippet.PublishedAt)
					dbVideo.PublishedAt = timeParsed

					dbVideos = append(dbVideos, &dbVideo)
				}

				// Save the data to the database
				err = database.CreateVideosBulk(dbVideos)
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
