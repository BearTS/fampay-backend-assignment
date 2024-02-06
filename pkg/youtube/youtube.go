package youtube

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	ytApi "google.golang.org/api/youtube/v3"
)

type Youtube struct {
	ApiKey []string

	currentIndexOfApiKey int
	youtubeClient        *ytApi.Service
}

type YoutubeService interface {
	SearchByQuery(query string, maxResults int64, after time.Time) ([]*ytApi.SearchResult, error)
}

func NewYoutubeClient(apiKey []string) *Youtube {
	if len(apiKey) < 1 {
		logrus.Fatal("Api Keys not present")
		return nil
	}

	client, err := ytApi.NewService(context.Background(), option.WithAPIKey(apiKey[0]))
	if err != nil {
		logrus.Fatal("Unable to create a yt client")
		return nil
	}
	return &Youtube{
		ApiKey:               apiKey,
		youtubeClient:        client,
		currentIndexOfApiKey: 0,
	}
}

// TODO: Upon a ratelimit error call a function which changes the api key
func (svc *Youtube) changeApiKey() error {
	return nil
}
