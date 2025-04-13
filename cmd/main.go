package main

import (
	"context"

	baseservice "github.com/yeencloud/lib-base"
	sharedConfig "github.com/yeencloud/lib-shared/config"
	lib_user "github.com/yeencloud/lib-user"
	"github.com/yeencloud/svc-mail/internal/adapters/event"
	"github.com/yeencloud/svc-mail/internal/adapters/smtp"
	"github.com/yeencloud/svc-mail/internal/adapters/templater"
	"github.com/yeencloud/svc-mail/internal/domain/config"
	"github.com/yeencloud/svc-mail/internal/service"
)

// TODO: Add metrics for sent mail
func main() {
	baseservice.Run("svc-mail", baseservice.Options{
		UseDatabase: false,
		UseEvents:   true,
	}, func(ctx context.Context, svc *baseservice.BaseService) error {
		templaterConfig, err := sharedConfig.FetchConfig[config.TemplaterConfig]()
		if err != nil {
			return err
		}

		err = svc.Validator.RegisterValidations(lib_user.Validations())
		if err != nil {
			return err
		}

		mqSubscriber, err := svc.GetMqSubscriber()
		if err != nil {
			return err
		}

		templateEngine, err := templater.NewTemplater(templaterConfig)
		if err != nil {
			return err
		}

		smtpConfig, err := sharedConfig.FetchConfig[config.SmtpConfig]()
		if err != nil {
			return err
		}

		smtpClient, err := smtp.NewSmtpClient(smtpConfig)
		if err != nil {
			return err
		}

		usecases := service.NewUsecases(templateEngine, smtpClient)

		subscriber := event.NewEventHandler(mqSubscriber, usecases)

		return subscriber.Subscribe(ctx)
	})
}
