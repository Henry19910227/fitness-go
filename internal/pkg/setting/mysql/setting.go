package mysql

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

func (s *setting) GetUserName() string {
	if s.mode == "debug" {
		return s.vp.GetString("Database.Debug.UserName")
	}
	return s.vp.GetString("Database.Release.UserName")
}

func (s *setting) GetPassword() string {
	if s.mode == "debug" {
		return s.vp.GetString("Database.Debug.Password")
	}
	return s.vp.GetString("Database.Release.Password")
}

func (s *setting) GetHost() string {
	if s.mode == "debug" {
		return s.vp.GetString("Database.Debug.Host")
	}
	return s.vp.GetString("Database.Release.Host")
}

func (s *setting) GetDatabase() string {
	if s.mode == "debug" {
		return s.vp.GetString("Database.Debug.DBName")
	}
	return s.vp.GetString("Database.Release.DBName")
}
