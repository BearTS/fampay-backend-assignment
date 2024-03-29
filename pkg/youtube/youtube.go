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
		logrus.Fatal("Unable to create a yt client", err)
		return nil
	}
	return &Youtube{
		ApiKey:               apiKey,
		youtubeClient:        client,
		currentIndexOfApiKey: 0,
	}
}

func (svc *Youtube) changeApiKey() error {
	if svc.currentIndexOfApiKey == len(svc.ApiKey)-1 {
		svc.currentIndexOfApiKey = 0
	} else {
		svc.currentIndexOfApiKey++
	}

	client, err := ytApi.NewService(context.Background(), option.WithAPIKey(svc.ApiKey[svc.currentIndexOfApiKey]))
	if err != nil {
		logrus.Error("Unable to create a yt client", err)
		return err
	}
	svc.youtubeClient = client
	return nil
}
