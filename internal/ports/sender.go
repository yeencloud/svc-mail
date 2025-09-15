package ports

import "context"

type Sender interface {
	SendMail(ctx context.Context, to string, subject string, body string) error
}
