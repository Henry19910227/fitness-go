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
	return setting.vp.GetString("Database.Release.UserName")
}

func (setting *mysql) GetPassword() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.Password")
	}
	return setting.vp.GetString("Database.Release.Password")
}

func (setting *mysql) GetHost() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.Host")
	}
	return setting.vp.GetString("Database.Release.Host")
}

func (setting *mysql) GetDatabase() string {
	if setting.mode == "debug" {
		return setting.vp.GetString("Database.Debug.DBName")
	}
	return setting.vp.GetString("Database.Release.DBName")
}