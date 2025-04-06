package ports

import (
	"context"

	"github.com/yeencloud/bpt-service/internal/domain"
)

type Usecases interface {
	Viewed(ctx context.Context, origin domain.ViewOrigin) ([]domain.ViewOrigin, error)
}
