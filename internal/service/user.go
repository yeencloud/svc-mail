package service

import (
	"context"

	metrics "github.com/yeencloud/lib-metrics"
	"github.com/yeencloud/svc-mail/internal/domain"
	mailMetrics "github.com/yeencloud/svc-mail/internal/domain/metrics"
)

func (s service) UserCreated(ctx context.Context, userCreated domain.UserCreated) error {
	mailBody, err := s.templater.RenderUserCreatedTemplate(ctx, userCreated)
	if err != nil {
		return err
	}

	subject := "Verify your email"
	sentMetric := mailMetrics.MailSentMetrics{
		Address: userCreated.Email,
		Subject: subject,
	}
	return nil
	err = s.smtp.SendMail(ctx, userCreated.Email, subject, mailBody)
	if err != nil {
		sentMetric.Status = "Error: " + err.Error()
	} else {
		sentMetric.Status = "Success"
	}

	_ = metrics.WritePoint(ctx, "mail", sentMetric)

	return err
}
