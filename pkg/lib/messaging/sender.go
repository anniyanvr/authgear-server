package messaging

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/authgear/authgear-server/pkg/api/event"
	"github.com/authgear/authgear-server/pkg/api/event/nonblocking"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/appdb"
	"github.com/authgear/authgear-server/pkg/lib/infra/mail"
	"github.com/authgear/authgear-server/pkg/lib/infra/sms"
	"github.com/authgear/authgear-server/pkg/lib/infra/whatsapp"
	"github.com/authgear/authgear-server/pkg/lib/translation"
	"github.com/authgear/authgear-server/pkg/util/log"
	"github.com/authgear/authgear-server/pkg/util/phone"
)

type Logger struct{ *log.Logger }

func NewLogger(lf *log.Factory) Logger {
	return Logger{lf.New("messaging")}
}

type EventService interface {
	DispatchEventImmediately(ctx context.Context, payload event.NonBlockingPayload) error
}

type MailSender interface {
	Send(opts mail.SendOptions) error
}

type SMSSender interface {
	Send(ctx context.Context, opts sms.SendOptions) error
}

type WhatsappSender interface {
	SendAuthenticationOTP(Ctx context.Context, opts *whatsapp.SendAuthenticationOTPOptions) error
}

type Sender struct {
	Logger         Logger
	Limits         Limits
	Events         EventService
	MailSender     MailSender
	SMSSender      SMSSender
	WhatsappSender WhatsappSender
	Database       *appdb.Handle

	DevMode config.DevMode

	MessagingFeatureConfig *config.MessagingFeatureConfig

	FeatureTestModeEmailSuppressed config.FeatureTestModeEmailSuppressed
	TestModeEmailConfig            *config.TestModeEmailConfig

	FeatureTestModeSMSSuppressed config.FeatureTestModeSMSSuppressed
	TestModeSMSConfig            *config.TestModeSMSConfig

	FeatureTestModeWhatsappSuppressed config.FeatureTestModeWhatsappSuppressed
	TestModeWhatsappConfig            *config.TestModeWhatsappConfig
}

func (s *Sender) SendEmailInNewGoroutine(ctx context.Context, msgType translation.MessageType, opts *mail.SendOptions) error {
	err := s.Limits.checkEmail(ctx, opts.Recipient)
	if err != nil {
		return err
	}

	if s.FeatureTestModeEmailSuppressed {
		return s.testModeSendEmail(ctx, msgType, opts)
	}

	if s.TestModeEmailConfig.Enabled {
		if r, ok := s.TestModeEmailConfig.MatchTarget(opts.Recipient); ok && r.Suppressed {
			return s.testModeSendEmail(ctx, msgType, opts)
		}
	}

	if s.DevMode {
		return s.devModeSendEmail(ctx, msgType, opts)
	}

	go func() {
		// Detach the deadline so that the context is not canceled along with the request.
		ctx = context.WithoutCancel(ctx)

		err := s.MailSender.Send(*opts)
		if err != nil {
			s.Logger.WithError(err).WithFields(logrus.Fields{
				"email": mail.MaskAddress(opts.Recipient),
			}).Error("failed to send email")
			return
		}

		err = s.Database.WithTx(ctx, func(ctx context.Context) error {
			return s.Events.DispatchEventImmediately(ctx, &nonblocking.EmailSentEventPayload{
				Sender:    opts.Sender,
				Recipient: opts.Recipient,
				Type:      string(msgType),
			})
		})
		if err != nil {
			s.Logger.Error("failed to emit email.sent event")
		}
	}()

	return nil
}

func (s *Sender) testModeSendEmail(ctx context.Context, msgType translation.MessageType, opts *mail.SendOptions) error {
	s.Logger.
		WithField("message_type", string(msgType)).
		WithField("recipient", opts.Recipient).
		WithField("body", opts.TextBody).
		WithField("sender", opts.Sender).
		WithField("subject", opts.Subject).
		WithField("reply_to", opts.ReplyTo).
		Warn("email is suppressed by test mode")

	desc := fmt.Sprintf("email (%v) to %v is suppressed by test mode.", msgType, opts.Recipient)
	return s.Events.DispatchEventImmediately(ctx, &nonblocking.EmailSuppressedEventPayload{
		Description: desc,
	})
}

func (s *Sender) devModeSendEmail(ctx context.Context, msgType translation.MessageType, opts *mail.SendOptions) error {
	s.Logger.
		WithField("message_type", string(msgType)).
		WithField("recipient", opts.Recipient).
		WithField("body", opts.TextBody).
		WithField("sender", opts.Sender).
		WithField("subject", opts.Subject).
		WithField("reply_to", opts.ReplyTo).
		Warn("email is suppressed by development mode")

	desc := fmt.Sprintf("email (%v) to %v is suppressed by development mode", msgType, opts.Recipient)
	return s.Events.DispatchEventImmediately(ctx, &nonblocking.EmailSuppressedEventPayload{
		Description: desc,
	})
}

func (s *Sender) SendSMSInNewGoroutine(ctx context.Context, msgType translation.MessageType, opts *sms.SendOptions) error {
	err := s.Limits.checkSMS(ctx, opts.To)
	if err != nil {
		return err
	}

	if s.FeatureTestModeSMSSuppressed {
		return s.testModeSendSMS(ctx, msgType, opts)
	}

	if s.TestModeSMSConfig.Enabled {
		if r, ok := s.TestModeSMSConfig.MatchTarget(opts.To); ok && r.Suppressed {
			return s.testModeSendSMS(ctx, msgType, opts)
		}
	}

	if s.DevMode {
		return s.devModeSendSMS(ctx, msgType, opts)
	}

	go func() {
		// Detach the deadline so that the context is not canceled along with the request.
		ctx = context.WithoutCancel(ctx)

		err := s.SMSSender.Send(ctx, *opts)
		if err != nil {
			s.Logger.WithError(err).WithFields(logrus.Fields{
				"phone": phone.Mask(opts.To),
			}).Error("failed to send SMS")
			return
		}

		err = s.Database.WithTx(ctx, func(ctx context.Context) error {
			return s.Events.DispatchEventImmediately(ctx, &nonblocking.SMSSentEventPayload{
				Sender:              opts.Sender,
				Recipient:           opts.To,
				Type:                string(msgType),
				IsNotCountedInUsage: s.MessagingFeatureConfig.SMSUsageCountDisabled,
			})
		})
		if err != nil {
			s.Logger.Error("failed to emit sms.sent event")
		}
	}()

	return nil
}

func (s *Sender) testModeSendSMS(ctx context.Context, msgType translation.MessageType, opts *sms.SendOptions) error {
	s.Logger.
		WithField("message_type", string(msgType)).
		WithField("recipient", opts.To).
		WithField("sender", opts.Sender).
		WithField("body", opts.Body).
		WithField("app_id", opts.AppID).
		WithField("template_name", opts.TemplateName).
		WithField("language_tag", opts.LanguageTag).
		WithField("template_variables", opts.TemplateVariables).
		Warn("SMS is suppressed in test mode")

	desc := fmt.Sprintf("SMS (%v) to %v is suppressed by test mode.", msgType, opts.To)
	return s.Events.DispatchEventImmediately(ctx, &nonblocking.SMSSuppressedEventPayload{
		Description: desc,
	})
}

func (s *Sender) devModeSendSMS(ctx context.Context, msgType translation.MessageType, opts *sms.SendOptions) error {
	s.Logger.
		WithField("message_type", string(msgType)).
		WithField("recipient", opts.To).
		WithField("sender", opts.Sender).
		WithField("body", opts.Body).
		WithField("app_id", opts.AppID).
		WithField("template_name", opts.TemplateName).
		WithField("language_tag", opts.LanguageTag).
		WithField("template_variables", opts.TemplateVariables).
		Warn("SMS is suppressed in development mode")

	desc := fmt.Sprintf("SMS (%v) to %v is suppressed by development mode.", msgType, opts.To)
	return s.Events.DispatchEventImmediately(ctx, &nonblocking.SMSSuppressedEventPayload{
		Description: desc,
	})
}

func (s *Sender) SendWhatsappImmediately(ctx context.Context, msgType translation.MessageType, opts *whatsapp.SendAuthenticationOTPOptions) error {
	err := s.Limits.checkWhatsapp(ctx, opts.To)
	if err != nil {
		return err
	}

	if s.FeatureTestModeWhatsappSuppressed {
		return s.testModeSendWhatsapp(ctx, msgType, opts)
	}

	if s.TestModeWhatsappConfig.Enabled {
		if r, ok := s.TestModeWhatsappConfig.MatchTarget(opts.To); ok && r.Suppressed {
			return s.testModeSendWhatsapp(ctx, msgType, opts)
		}
	}

	if s.DevMode {
		return s.devModeSendWhatsapp(ctx, msgType, opts)
	}

	// Send immediately.
	err = s.WhatsappSender.SendAuthenticationOTP(ctx, opts)
	if err != nil {
		return err
	}

	err = s.Events.DispatchEventImmediately(ctx, &nonblocking.WhatsappSentEventPayload{
		Recipient:           opts.To,
		Type:                string(msgType),
		IsNotCountedInUsage: s.MessagingFeatureConfig.WhatsappUsageCountDisabled,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Sender) testModeSendWhatsapp(ctx context.Context, msgType translation.MessageType, opts *whatsapp.SendAuthenticationOTPOptions) error {
	s.Logger.
		WithField("message_type", string(msgType)).
		WithField("recipient", opts.To).
		WithField("otp", opts.OTP).
		Warn("Whatsapp is suppressed in test mode")

	desc := fmt.Sprintf("Whatsapp (%v) to %v is suppressed by test mode.", msgType, opts.To)
	return s.Events.DispatchEventImmediately(ctx, &nonblocking.WhatsappSuppressedEventPayload{
		Description: desc,
	})
}

func (s *Sender) devModeSendWhatsapp(ctx context.Context, msgType translation.MessageType, opts *whatsapp.SendAuthenticationOTPOptions) error {
	s.Logger.
		WithField("recipient", opts.To).
		WithField("otp", opts.OTP).
		Warn("Whatsapp is suppressed in development mode")

	desc := fmt.Sprintf("Whatsapp (%v) to %v is suppressed by development mode.", msgType, opts.To)
	return s.Events.DispatchEventImmediately(ctx, &nonblocking.WhatsappSuppressedEventPayload{
		Description: desc,
	})
}
