package ports

import (
	"github.com/yeencloud/bpt-service/internal/ports/repository"
	"github.com/yeencloud/lib-base/events"
)

type Ports struct {
	ViewOriginRepo database.ViewOriginRepository
	EventPublisher events.Publisher
}
