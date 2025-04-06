package service

import (
	"github.com/yeencloud/bpt-service/internal/ports"
	"github.com/yeencloud/bpt-service/internal/ports/repository"
	"github.com/yeencloud/lib-base/events"
)

type service struct {
	ports *ports.Ports
}

func NewUsecases(viewRepository database.ViewOriginRepository, eventPublisher events.Publisher) service {
	return service{
		ports: &ports.Ports{
			ViewOriginRepo: viewRepository,

			EventPublisher: eventPublisher,
		},
	}
}
