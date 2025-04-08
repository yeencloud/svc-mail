package service

import (
	"context"

	"github.com/yeencloud/svc-mail/internal/domain"
)

func (s service) UserCreated(ctx context.Context, userCreated domain.UserCreated) error {
	mailBody, err := s.templater.RenderUserCreatedTemplate(ctx, userCreated)
	if err != nil {
		return err
	}

	err = s.smtp.SendMail(ctx, userCreated.Email, "Verify your email", mailBody)

	return err
}
