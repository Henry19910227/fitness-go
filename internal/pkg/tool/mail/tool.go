package mail

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/mail"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type tool struct {
	setting mail.Setting
}

func New(setting mail.Setting) Tool {
	return &tool{setting}
}

func (t *tool) Send(recp string, subject string, body string) error {
	auth := smtp.PlainAuth("", t.setting.Sender(), t.setting.Password(), t.setting.SMTPHostName())
	smtpServer := t.setting.SMTPHostName() + ":" + t.setting.Port()
	e := email.NewEmail()
	e.From = fmt.Sprintf("健身平台平台 <%s>", t.setting.Sender())
	e.To = []string{recp}
	e.Subject = subject
	e.Text = []byte(body)
	if err := e.Send(smtpServer, auth); err != nil {
		return err
	}
	return nil
}
