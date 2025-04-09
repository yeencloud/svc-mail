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
