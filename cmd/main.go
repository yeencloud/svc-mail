package main

import (
	"context"
	"errors"

	"github.com/yeencloud/bpt-service/internal/adapters/database"
	"github.com/yeencloud/bpt-service/internal/adapters/http"
	"github.com/yeencloud/bpt-service/internal/service"
	baseservice "github.com/yeencloud/lib-base"
)

func main() {
	baseservice.Run("base-service", baseservice.Options{
		UseDatabase: true,
		UseEvents:   true,
	}, func(ctx context.Context, svc *baseservice.BaseService) error {
		dbEngine, err := svc.GetDatabase()
		if err != nil {
			return err
		}

		db, err := database.NewDatabase(ctx, dbEngine.Gorm)
		if err != nil {
			return err
		}

		httpServer, err := svc.GetHttpServer()
		if err != nil {
			return err
		}

		mqPublisher, err := svc.GetMqPublisher()
		if err != nil {
			return err
		}

		mqSubscriber, err := svc.GetMqSubscriber()
		if err != nil {
			return err
		}

		myChannelReceiver := mqSubscriber.Subscribe("my-channel")
		myChannelReceiver.Handle("PING_EVENT", func(ctx context.Context, event any) error {
			return nil
		})
		myChannelReceiver.Handle("PONG_EVENT", func(ctx context.Context, event any) error {
			return nil
		})

		otherChannelReceiver := mqSubscriber.Subscribe("other-channel")
		otherChannelReceiver.Handle("PING_EVENT", func(ctx context.Context, event any) error {
			return nil
		})
		otherChannelReceiver.Handle("PONG_EVENT", func(ctx context.Context, event any) error {
			return errors.New("oh no")
		})

		usecases := service.NewUsecases(database.NewViewOriginRepo(), mqPublisher)
		http.NewHTTPServer(httpServer, usecases, db.Gorm)

		return nil
	})
}
