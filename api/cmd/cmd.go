package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	api "github.com/BearTS/fampay-backend-assignment/api/pkg/controllers"
	"github.com/BearTS/fampay-backend-assignment/api/pkg/services/searchsvc"
	"github.com/BearTS/fampay-backend-assignment/pkg/config"
	"github.com/BearTS/fampay-backend-assignment/pkg/db"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func RootCmd() *cobra.Command {
	configuration := config.ReadFromEnv()
	opts := &api.Options{
		Path:                "/v1",
		Port:                configuration.Port,
		ShutdownGracePeriod: 5 * time.Second,
	}

	c := &cobra.Command{
		Use:   "api",
		Short: "serves the api",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			gormDB, _ := db.Connection()
			database := db.NewDB(gormDB)

			var deps api.Dependencies
			deps.DB = &database
			searchSvc := searchsvc.NewService(database)
			deps.Services.SearchService = searchSvc

			svc, err := api.NewService(ctx, opts, deps)
			if err != nil {
				return Cancel(err, cancel, svc)
			}

			svc.Start()

			logrus.Info("api serving")
			signals := make(chan os.Signal, 1)
			signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
			select {
			case <-ctx.Done():
				logrus.Info("terminating: context canceled")
			case <-signals:
				logrus.Info("terminating: via signal")
			}
			return Cancel(nil, cancel, svc)
		},
	}

	return c
}

func Cancel(err error, cancel context.CancelFunc, closers ...io.Closer) error {
	if cancel != nil {
		cancel()
	}
	var eg errgroup.Group
	// Close all closers
	for i := range closers {
		closer := closers[i]
		if !isNil(closer) {
			eg.Go(closer.Close)
		}
	}
	waitErr := eg.Wait()
	if waitErr == nil {
		return err
	}
	if err == nil {
		return waitErr
	}
	return fmt.Errorf("%v: %v", err, waitErr.Error())
}

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}
