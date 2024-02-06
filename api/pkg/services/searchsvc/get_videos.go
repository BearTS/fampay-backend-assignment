package searchsvc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/BearTS/fampay-backend-assignment/api/pkg/routes"
)

func (s *SearchSvc) GetVideos(ctx context.Context, params routes.GetVideosParams) (routes.GetVideosResponse, int, error) {
	var response routes.GetVideosResponse

	limit := 5
	if params.Limit != nil {
		limit = *params.Limit
	}

	offset := 0
	if params.Page != nil {
		offset = *params.Page * limit
	}

	totalVideos, err := s.DB.GetAllVideosCount()
	if err != nil {
		return response, http.StatusInternalServerError, fmt.Errorf("unable to get total videos: %w", err)
	}

	videos, err := s.DB.GetAllVideosPaginated(offset, limit, params.Title, params.Description)
	if err != nil {
		return response, http.StatusInternalServerError, fmt.Errorf("unable to get videos: %w", err)
	}

	for _, video := range videos {
		response.Videos = append(response.Videos, routes.Video{
			Id:          video.VideoId,
			Title:       video.Title,
			Description: video.Description,
			Thumbnail:   video.Thumbnail,
			PublishedAt: video.PublishedAt,
		})
	}

	response.Limit = limit
	response.Page = offset / limit
	response.Total = totalVideos

	return response, http.StatusOK, nil
}
