package service

import (
	"github.com/yeencloud/svc-mail/internal/ports"
)

type service struct {
	templater ports.Templater
	smtp      ports.Smtp
}

func NewUsecases(templater ports.Templater, smtp ports.Smtp) service {
	return service{
		templater: templater,
		smtp:      smtp,
	}
}
