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
	if s.mode == "release" {
		return s.vp.GetString("Database.Release.UserName")
	}
	if s.mode == "production" {
		return s.vp.GetString("Database.Production.UserName")
	}
	return ""
}

func (s *setting) GetPassword() string {
	if s.mode == "debug" {
		return s.vp.GetString("Database.Debug.Password")
	}
	if s.mode == "release" {
		return s.vp.GetString("Database.Release.Password")
	}
	if s.mode == "production" {
		return s.vp.GetString("Database.Production.Password")
	}
	return ""
}

func (s *setting) GetHost() string {
	if s.mode == "debug" {
		return s.vp.GetString("Database.Debug.Host")
	}
	if s.mode == "release" {
		return s.vp.GetString("Database.Release.Host")
	}
	if s.mode == "production" {
		return s.vp.GetString("Database.Production.Host")
	}
	return ""
}

func (s *setting) GetDatabase() string {
	if s.mode == "debug" {
		return s.vp.GetString("Database.Debug.DBName")
	}
	if s.mode == "release" {
		return s.vp.GetString("Database.Release.DBName")
	}
	if s.mode == "production" {
		return s.vp.GetString("Database.Production.DBName")
	}
	return ""
}
