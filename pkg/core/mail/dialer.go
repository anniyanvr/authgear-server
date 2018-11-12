package mail

import (
	"github.com/go-gomail/gomail"
	"github.com/skygeario/skygear-server/pkg/core/config"
)

func NewDialer(c config.SMTPConfiguration) *gomail.Dialer {
	if c.Host == "" {
		panic("mail server is not configured")
	}

	return gomail.NewPlainDialer(c.Host, c.Port, c.Login, c.Password)
}
