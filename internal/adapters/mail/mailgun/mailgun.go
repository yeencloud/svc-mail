package mailgun

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/mailgun/mailgun-go/v5"
	"github.com/yeencloud/svc-mail/internal/domain/config"
)

type MailGunClient struct {
	config *config.MailGun

	client *mailgun.Client
}

func (s MailGunClient) SendMail(ctx context.Context, to string, subject string, body string) error {
	m := mailgun.NewMessage(s.config.Domain, s.config.Sender, subject, "", to)
	m.SetHTML(body)

	_, err := s.client.Send(ctx, m)
	return err
}

func NewMailgunClient(config *config.MailGun) (*MailGunClient, error) {
	client := mailgun.NewMailgun(config.ApiKey.Value)
	err := client.SetAPIBase(config.ApiUrl)
	if err != nil {
		return nil, err
	}

	return &MailGunClient{
		config: config,
		client: client,
	}, nil
}
