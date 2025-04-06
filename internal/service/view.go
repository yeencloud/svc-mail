package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/yeencloud/bpt-service/internal/domain"
)

type ViewedEvent struct {
	Origin domain.ViewOrigin
}

func (v ViewedEvent) Channel() string {
	return "other-channel"
}

func (v ViewedEvent) EventType() string {
	return "PING_EVENT"
}

func (s service) Viewed(ctx context.Context, origin domain.ViewOrigin) ([]domain.ViewOrigin, error) {
	err := s.ports.ViewOriginRepo.AddViewOrigin(ctx, origin)
	if err != nil {
		return nil, err
	}

	log.WithContext(ctx).Warn("Publishing event to other channel")

	err = s.ports.EventPublisher.Publish(ctx, ViewedEvent{
		Origin: origin,
	})
	if err != nil {
		return nil, err
	}

	return s.ports.ViewOriginRepo.ListViewOrigins(ctx)
}
