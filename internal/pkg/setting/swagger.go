package setting

import "github.com/spf13/viper"

type swagger struct {
	vp   *viper.Viper
	mode string
}

func NewSwagger(viperTool *viper.Viper) Swagger {
	return &swagger{viperTool, viperTool.GetString("Server.RunMode")}
}

func (s *swagger) GetProtocol() string {
	if s.mode == "debug" {
		return s.vp.GetString("Swagger.Debug.Protocol")
	}
	if s.mode == "release" {
		return s.vp.GetString("Swagger.Release.Protocol")
	}
	if s.mode == "production" {
		return s.vp.GetString("Swagger.Production.Protocol")
	}
	return ""
}

func (s *swagger) GetVersion() string {
	if s.mode == "debug" {
		return s.vp.GetString("App.Debug.Version")
	}
	if s.mode == "release" {
		return s.vp.GetString("App.Release.Version")
	}
	if s.mode == "production" {
		return s.vp.GetString("App.Production.Version")
	}
	return ""
}

func (s *swagger) GetHost() string {
	if s.mode == "debug" {
		return s.vp.GetString("Swagger.Debug.Host")
	}
	if s.mode == "release" {
		return s.vp.GetString("Swagger.Release.Host")
	}
	if s.mode == "production" {
		return s.vp.GetString("Swagger.Production.Host")
	}
	return ""
}

func (s *swagger) GetBasePath() string {
	if s.mode == "debug" {
		return s.vp.GetString("Swagger.Debug.BasePath")
	}
	if s.mode == "release" {
		return s.vp.GetString("Swagger.Release.BasePath")
	}
	if s.mode == "production" {
		return s.vp.GetString("Swagger.Production.BasePath")
	}
	return ""
}
