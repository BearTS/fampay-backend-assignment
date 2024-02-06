package api

import (
	"context"
	"fmt"
	"time"

	"github.com/BearTS/fampay-backend-assignment/api/pkg/routes"
	"github.com/BearTS/fampay-backend-assignment/pkg/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type EchoContext interface {
	echo.Context
}

type EchoServer interface {
	Start(string) error
	Shutdown(ctx context.Context) error
}

type Options struct {
	Path                string
	Port                int
	ShutdownGracePeriod time.Duration
}

type ApiSvc struct {
	ctx      context.Context
	opts     *Options
	server   EchoServer
	services Services
}

type Services struct {
	SearchService SearchService
}

type Dependencies struct {
	DB       *db.DB
	Services Services
}

// NewService - constructor for Service
func NewService(ctx context.Context, opts *Options, deps Dependencies) (*ApiSvc, error) {
	svc := &ApiSvc{
		ctx:      ctx,
		opts:     opts,
		services: deps.Services,
	}
	svc.server = svc.createServer()
	return svc, nil
}

// Start starts the API
func (svc *ApiSvc) Start() {
	go func() {
		addr := fmt.Sprintf(":%d", svc.opts.Port)
		if err := svc.server.Start(addr); err != nil {
			logrus.Fatal(err)
		}
	}()
}

// Close closes the API
func (svc *ApiSvc) Close() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), svc.opts.ShutdownGracePeriod)
	defer cancel()
	return svc.server.Shutdown(ctx)
}

func (svc *ApiSvc) createServer() EchoServer {
	server := echo.New()
	// Default CORS
	server.Use(middleware.CORS())
	apiGroup := server.Group("")
	routes.RegisterHandlersWithBaseURL(apiGroup, svc, svc.opts.Path)
	return server
}
