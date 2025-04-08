package event

import (
	"context"

	events "github.com/yeencloud/lib-events"
	"github.com/yeencloud/svc-mail/internal/domain"
	"github.com/yeencloud/svc-mail/internal/ports"
)

type EventHandler struct {
	subscriber *events.Subscriber

	usecases ports.Usecases
}

func NewEventHandler(subscriber *events.Subscriber, usecases ports.Usecases) *EventHandler {
	eventHandler := &EventHandler{
		subscriber: subscriber,
		usecases:   usecases,
	}

	return eventHandler
}

type UserCreatedEventBody struct {
	Username      string
	Email         string
	Code          string
	CodeExpiresAt string
}

func (e *EventHandler) Listen(ctx context.Context) error {
	myChannelReceiver := e.subscriber.Subscribe("user_events")
	myChannelReceiver.Handle("USER_CREATED", func(ctx context.Context, event any) error {
		createdUserEvent, err := events.DecodeEvent[UserCreatedEventBody](event)
		if err != nil {
			return err
		}

		user := domain.UserCreated{
			Username:      createdUserEvent.Username,
			Email:         createdUserEvent.Email,
			Code:          createdUserEvent.Code,
			CodeExpiresAt: createdUserEvent.CodeExpiresAt,
		}
		return e.usecases.UserCreated(ctx, user)
	})

	return e.subscriber.Listen(ctx)
}
