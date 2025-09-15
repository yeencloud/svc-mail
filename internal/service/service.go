package service

import (
	"github.com/yeencloud/svc-mail/internal/ports"
)

type service struct {
	templater ports.Templater
	sender    ports.Sender
}

func NewUsecases(templater ports.Templater, smtp ports.Sender) service {
	return service{
		templater: templater,
		sender:    smtp,
	}
}
