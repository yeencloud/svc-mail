package service

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	metrics "github.com/yeencloud/lib-metrics"
	"github.com/yeencloud/svc-mail/internal/domain"
	mailMetrics "github.com/yeencloud/svc-mail/internal/domain/metrics"
)

func (s service) buildToField(username string, recipient string) string {
	if username == "" {
		return recipient
	}

	return fmt.Sprintf("%s <%s>", username, recipient)
}

func (s service) UserCreated(ctx context.Context, userCreated domain.UserCreated) error {
	mailBody, err := s.templater.RenderUserCreatedTemplate(ctx, userCreated)
	if err != nil {
		return err
	}

	to := s.buildToField(userCreated.Username, userCreated.Email)
	subject := "Verify your email"
	log.WithContext(ctx).
		WithField("to", userCreated.Email).WithField("subject", subject).
		Info("Sending email")
	err = s.sender.SendMail(ctx, to, subject, mailBody)

	sentMetric := mailMetrics.MailSentMetrics{
		Address: userCreated.Email,
		Subject: subject,
	}
	
	if err != nil {
		sentMetric.Status = "Error: " + err.Error()
	} else {
		sentMetric.Status = "Success"
	}

	_ = metrics.WritePoint(ctx, "mail", sentMetric)

	return err
}
