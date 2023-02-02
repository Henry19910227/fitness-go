package setting

import (
	"github.com/spf13/viper"
)

type mysql struct {
	vp   *viper.Viper
	mode string
}

func NewMysql(viperTool *viper.Viper) Mysql {
	return &mysql{viperTool, viperTool.GetString("Server.RunMode")}
}

func (setting *mysql) GetUserName() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.UserName")
	}
	if setting.mode == "release" {
		return setting.vp.GetString("Database.Release.UserName")
	}
	if setting.mode == "production" {
		return setting.vp.GetString("Database.Production.UserName")
	}
	return ""
}

func (setting *mysql) GetPassword() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.Password")
	}
	if setting.mode == "release" {
		return setting.vp.GetString("Database.Release.Password")
	}
	if setting.mode == "production" {
		return setting.vp.GetString("Database.Production.Password")
	}
	return ""
}

func (setting *mysql) GetHost() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.Host")
	}
	if setting.mode == "release" {
		return setting.vp.GetString("Database.Release.Host")
	}
	if setting.mode == "production" {
		return setting.vp.GetString("Database.Production.Host")
	}
	return ""
}

func (setting *mysql) GetDatabase() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.DBName")
	}
	if setting.mode == "release" {
		return setting.vp.GetString("Database.Release.DBName")
	}
	if setting.mode == "production" {
		return setting.vp.GetString("Database.Production.DBName")
	}
	return ""
}
