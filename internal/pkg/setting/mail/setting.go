package mail

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/spf13/viper"
)

type setting struct {
	vp   *viper.Viper
	mode string
}

func New() Setting {
	return &setting{vp: vp.Shared(), mode: build.RunMode()}
}

func (s *setting) SMTPHostName() string {
	return s.vp.GetString("Mail.SMTPHostName")
}

func (s *setting) Port() string {
	if s.mode == "debug" {
		return s.vp.GetString("Mail.Debug.Port")
	}
	return s.vp.GetString("Mail.Release.Port")
}

func (s *setting) Sender() string {
	if s.mode == "debug" {
		return s.vp.GetString("Mail.Debug.Sender")
	}
	return s.vp.GetString("Mail.Release.Sender")
}

func (s *setting) Password() string {
	if s.mode == "debug" {
		return s.vp.GetString("Mail.Debug.Password")
	}
	return s.vp.GetString("Mail.Release.Password")
}
