package config

// TODO: Add a configuration for ssl/tls
type SmtpConfig struct {
	Host     string `config:"SMTP_HOST" default:"localhost"`
	Port     int    `config:"SMTP_PORT" default:"587"`
	Username string `config:"SMTP_USERNAME" default:""`
	// TODO: Set Password to a secret
	Password string `config:"SMTP_PASSWORD" default:""`
	From     string `config:"SMTP_FROM" default:"noreply@localhost"`
}
