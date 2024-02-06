package youtube

import (
	"time"

	"github.com/sirupsen/logrus"
	ytApi "google.golang.org/api/youtube/v3"
)

// Ref: https://developers.google.com/youtube/v3/docs/search/list#go

func (svc *Youtube) SearchByQuery(query string, maxResults int64, after time.Time) ([]*ytApi.SearchResult, error) {
	request := svc.youtubeClient.Search.List([]string{"id", "snippet"}).
		Q(query).MaxResults(maxResults).Type("video").Order("date").
		PublishedAfter(after.Format(time.RFC3339))

	response, err := request.Do()
	if err != nil {
		// TODO: Could add a check here on which type of error it is and then change the api key accordingly
		logrus.Error("Error occured while searching by query: ", err)
		logrus.Info("Changing the api key")
		svc.changeApiKey()
		return nil, err
	}

	return response.Items, err
}
