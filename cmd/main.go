package main

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	baseservice "github.com/yeencloud/lib-base"
	sharedConfig "github.com/yeencloud/lib-shared/config"
	libuser "github.com/yeencloud/lib-user"
	"github.com/yeencloud/svc-mail/internal/adapters/event"
	"github.com/yeencloud/svc-mail/internal/adapters/mail/mailgun"
	"github.com/yeencloud/svc-mail/internal/adapters/mail/smtp"
	"github.com/yeencloud/svc-mail/internal/adapters/templater"
	"github.com/yeencloud/svc-mail/internal/domain"
	"github.com/yeencloud/svc-mail/internal/domain/config"
	"github.com/yeencloud/svc-mail/internal/ports"
	"github.com/yeencloud/svc-mail/internal/service"
)

func main() {
	log.Info("Will run service")

	baseservice.Run("svc-mail", baseservice.Options{
		UseDatabase: false,
		UseEvents:   true,
	}, func(ctx context.Context, svc *baseservice.BaseService) error {
		templaterConfig, err := sharedConfig.FetchConfig[config.TemplaterConfig]()
		if err != nil {
			return err
		}

		err = svc.Validator.RegisterValidations(libuser.Validations())
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

		provider, err := sharedConfig.FetchConfig[config.MailProvider]()
		if err != nil {
			return err
		}

		var mailSender ports.Sender
		switch provider.Provider {
		case "SMTP":
			smtpConfig, err := sharedConfig.FetchConfig[config.SmtpConfig]()
			if err != nil {
				return err
			}

			mailSender, err = smtp.NewSmtpClient(smtpConfig)
			if err != nil {
				return err
			}
		case "Mailgun":
			mailgunConfig, err := sharedConfig.FetchConfig[config.MailGun]()
			if err != nil {
				return err
			}

			mailSender, err = mailgun.NewMailgunClient(mailgunConfig)
			if err != nil {
				return err
			}
		default:
			return errors.Join(domain.ErrUnknownMailProvider, errors.New("available: SMTP, Mailgun"))
		}

		log.Info("Mail provider set to: " + provider.Provider)

		usecases := service.NewUsecases(templateEngine, mailSender)

		subscriber := event.NewEventHandler(mqSubscriber, usecases)

		return subscriber.Subscribe(ctx)
	})
}
