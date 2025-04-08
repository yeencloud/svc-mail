package ports

import (
	"context"

	"github.com/yeencloud/svc-mail/internal/domain"
)

type Usecases interface {
	UserCreated(ctx context.Context, origin domain.UserCreated) error
}
