package ports

import (
	"context"

	"github.com/yeencloud/svc-mail/internal/domain"
)

type Templater interface {
	RenderUserCreatedTemplate(ctx context.Context, creation domain.UserCreated) (string, error)
}
