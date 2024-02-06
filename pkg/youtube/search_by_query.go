package youtube

import (
	"time"

	"github.com/sirupsen/logrus"
	ytApi "google.golang.org/api/youtube/v3"
)

// Ref: https://developers.google.com/youtube/v3/docs/search/list#go

func (svc *Youtube) SearchByQuery(query string, maxResults int64, after time.Time) ([]*ytApi.SearchResult, error) {
	request := svc.youtubeClient.Search.List([]string{"id", "snippet"}).
		Q(query).MaxResults(maxResults).
		PublishedAfter(after.Format(time.RFC3339))

	response, err := request.Do()
	if err != nil {
		// TODO: Upon a ratelimit error call a function which changes the api key
		logrus.Error("Error occured while searching by query: ", err)
		return nil, nil
	}

	return response.Items, err
}
