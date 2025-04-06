package database

import (
	"context"

	"github.com/yeencloud/bpt-service/internal/domain"
)

type ViewOriginRepository interface {
	AddViewOrigin(ctx context.Context, origin domain.ViewOrigin) error
	ListViewOrigins(ctx context.Context) ([]domain.ViewOrigin, error)
	CountViewOrigin(ctx context.Context) (int64, error)
}
