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
	Username      string `validate:"required,username"`
	Email         string `validate:"required,email"`
	Code          string `validate:"required,validation_code"`
	CodeExpiresAt string `validate:"required,date_time"`
}

func (e *EventHandler) Subscribe(ctx context.Context) error {
	myChannelReceiver := e.subscriber.Subscribe("user_events")
	myChannelReceiver.Handle("USER_CREATED", func(ctx context.Context, eventJson string) error {
		createdUserEvent, err := events.DecodeEvent[UserCreatedEventBody](e.subscriber.Validator, ctx, eventJson)
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

	return nil
}
