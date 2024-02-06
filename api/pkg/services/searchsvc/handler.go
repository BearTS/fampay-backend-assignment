package searchsvc

import (
	"context"

	"github.com/BearTS/fampay-backend-assignment/api/pkg/routes"
	"github.com/BearTS/fampay-backend-assignment/pkg/db"
)

type SearchSvc struct {
	DB DB
}

type DB interface {
	GetAllVideosPaginated(offset int, limit int, title *string, description *string) ([]*db.Videos, error)
	GetAllVideosCount() (int64, error)
}

type Interface interface {
	GetVideos(ctx context.Context, params routes.GetVideosParams) (routes.GetVideosResponse, int, error)
}

func NewService(db DB) *SearchSvc {
	return &SearchSvc{DB: db}
}
