package api

import (
	"context"

	"github.com/BearTS/fampay-backend-assignment/api/pkg/routes"
	"github.com/labstack/echo/v4"
)

type SearchService interface {
	GetVideos(ctx context.Context, params routes.GetVideosParams) (routes.GetVideosResponse, int, error)
}

func (s *ApiSvc) GetVideos(ctx echo.Context, params routes.GetVideosParams) error {
	response, status, err := s.services.SearchService.GetVideos(ctx.Request().Context(), params)
	if err != nil {
		return ctx.JSON(status, err.Error())
	}
	return ctx.JSON(status, response)
}

func (svc *ApiSvc) Ping(ctx echo.Context) error {
	return ctx.String(200, "pong")
}
