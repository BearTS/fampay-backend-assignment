// Package routes provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// GetVideosResponse defines model for GetVideosResponse.
type GetVideosResponse struct {
	Limit  int     `json:"limit"`
	Page   int     `json:"page"`
	Total  int64   `json:"total"`
	Videos []Video `json:"videos"`
}

// Video defines model for Video.
type Video struct {
	Description string    `json:"description"`
	Id          string    `json:"id"`
	PublishedAt time.Time `json:"publishedAt"`
	Thumbnail   string    `json:"thumbnail"`
	Title       string    `json:"title"`
}

// GetVideosParams defines parameters for GetVideos.
type GetVideosParams struct {
	// Page page number
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// Limit number of videos per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Title title of the video
	Title *string `form:"title,omitempty" json:"title,omitempty"`

	// Description description of the video
	Description *string `form:"description,omitempty" json:"description,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Ping
	// (GET /ping)
	Ping(ctx echo.Context) error
	// Get videos
	// (GET /videos)
	GetVideos(ctx echo.Context, params GetVideosParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Ping converts echo context to params.
func (w *ServerInterfaceWrapper) Ping(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Ping(ctx)
	return err
}

// GetVideos converts echo context to params.
func (w *ServerInterfaceWrapper) GetVideos(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetVideosParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "title" -------------

	err = runtime.BindQueryParameter("form", true, false, "title", ctx.QueryParams(), &params.Title)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter title: %s", err))
	}

	// ------------- Optional query parameter "description" -------------

	err = runtime.BindQueryParameter("form", true, false, "description", ctx.QueryParams(), &params.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter description: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetVideos(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/ping", wrapper.Ping)
	router.GET(baseURL+"/videos", wrapper.GetVideos)

}
