package smtp

import (
	"context"
	"fmt"

	"github.com/wneessen/go-mail"
	shared "github.com/yeencloud/lib-shared/domain"
	"github.com/yeencloud/svc-mail/internal/domain/config"
)

type SmtpClient struct {
	config *config.SmtpConfig

	client *mail.Client
}

func (s SmtpClient) SendMail(ctx context.Context, to string, subject string, body string) error {
	message := mail.NewMsg()
	from := fmt.Sprintf("%s <%s>", shared.AppName, s.config.From)
	if err := message.From(from); err != nil {
		return err
	}
	if err := message.To(to); err != nil {
		return err
	}
	message.Subject(subject)
	message.SetBodyString(mail.TypeTextHTML, body)

	if err := s.client.DialAndSendWithContext(ctx, message); err != nil {
		return err
	}
	return nil
}

func NewSmtpClient(config *config.SmtpConfig) (*SmtpClient, error) {
	client, err := mail.NewClient(
		config.Host,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(config.Username),
		mail.WithPassword(config.Password.Value),
	)
	if err != nil {
		return nil, err
	}

	return &SmtpClient{
		config: config,
		client: client,
	}, nil
}
