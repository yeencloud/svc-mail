package hermes

import (
	"context"
	"fmt"
	"time"

	"github.com/matcornic/hermes"
	shared "github.com/yeencloud/lib-shared"
	"github.com/yeencloud/svc-mail/internal/domain"
)

type Hermes struct {
	engine *hermes.Hermes
}

func NewHermes() Hermes {
	h := hermes.Hermes{
		Theme: new(hermes.Flat),

		Product: hermes.Product{
			Name: shared.AppName,
			Link: "https://yeencloud.io",

			Logo: "https://avatars.githubusercontent.com/u/132744176?s=400&u=f70c42e5c6fb824c04f1c593d183c0f84ffccc85&v=4",
		},
	}

	return Hermes{
		engine: &h,
	}
}

func (h Hermes) beforeGeneration() {
	h.engine.Product.Copyright = fmt.Sprintf("Copyright © %d %s. All rights reserved.", time.Now().Year(), shared.AppName)
}

func (h Hermes) RenderUserCreatedTemplate(ctx context.Context, creation domain.UserCreated) (string, error) {
	h.beforeGeneration()

	email := hermes.Email{
		Body: hermes.Body{
			Name: creation.Username,
			Intros: []string{
				fmt.Sprintf("Welcome to %s!", shared.AppName),
				"To continue, please verify your email address.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "You can use this code to verify your account",
					InviteCode:   creation.Code,
				},
			},
			Signature: "Thanks",
		},
	}

	return h.engine.GenerateHTML(email)
}
