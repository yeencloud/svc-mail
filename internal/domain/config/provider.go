package config

import "github.com/yeencloud/lib-shared/config"

// TODO: Add a configuration for ssl/tls
type SmtpConfig struct {
	Host     string        `config:"SMTP_HOST" default:"localhost"`
	Port     int           `config:"SMTP_PORT" default:"587"`
	Username string        `config:"SMTP_USERNAME" default:""`
	Password config.Secret `config:"SMTP_PASSWORD" default:""`
	From     string        `config:"SMTP_FROM" default:"noreply@localhost"`
}

type MailGun struct {
	Sender string        `config:"MAILGUN_SENDER"`
	Domain string        `config:"MAILGUN_DOMAIN"`
	ApiUrl string        `config:"MAILGUN_ENDPOINT" default:"https://api.mailgun.net/"`
	ApiKey config.Secret `config:"MAILGUN_APIKEY"`
}

type MailProvider struct {
	Provider string `config:"MAIL_PROVIDER"`
}
