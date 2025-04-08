package templater

import (
	"errors"

	"github.com/yeencloud/svc-mail/internal/adapters/templater/hermes"
	"github.com/yeencloud/svc-mail/internal/domain/config"
	"github.com/yeencloud/svc-mail/internal/ports"
)

func NewTemplater(config *config.TemplaterConfig) (ports.Templater, error) {

	switch config.Engine {
	case "HERMES":
		return hermes.NewHermes(), nil
	}
	return nil, errors.New("Templater engine not supported")
}
