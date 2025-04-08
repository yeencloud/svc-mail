package service

import (
	"context"
	"os"

	"github.com/yeencloud/svc-mail/internal/domain"
)

func (s service) UserCreated(ctx context.Context, userCreated domain.UserCreated) error {
	mailBody, err := s.templater.RenderUserCreatedTemplate(ctx, userCreated)
	if err != nil {
		return err
	}

	fileToWrite := "user_created.html"
	err = os.WriteFile(fileToWrite, []byte(mailBody), 0644)
	if err != nil {
		return err
	}

	err = s.smtp.SendMail(ctx, userCreated.Email, "Verify your email", mailBody)

	return err
}
