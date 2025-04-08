package ports

import "context"

type Smtp interface {
	SendMail(ctx context.Context, to string, subject string, body string) error
}
