package templater

import (
	"errors"

	"github.com/yeencloud/svc-mail/internal/adapters/templater/hermes"
	"github.com/yeencloud/svc-mail/internal/domain/config"
	"github.com/yeencloud/svc-mail/internal/ports"
)

func NewTemplater(config *config.TemplaterConfig) (ports.Templater, error) {
	if config.Engine == "HERMES" {
		return hermes.NewHermes(), nil
	}

	// TODO: Custom error
	return nil, errors.New("Templater engine not supported")
}
