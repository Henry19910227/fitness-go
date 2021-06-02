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
	return s.vp.GetString("Swagger.Release.Protocol")
}

func (s *swagger) GetVersion() string {
	if s.mode == "debug" {
		return s.vp.GetString("App.Debug.Version")
	}
	return s.vp.GetString("App.Release.Version")
}

func (s *swagger) GetHost() string {
	if s.mode == "debug" {
		return s.vp.GetString("Swagger.Debug.Host")
	}
	return s.vp.GetString("Swagger.Release.Host")
}

func (s *swagger) GetBasePath() string {
	if s.mode == "debug" {
		return s.vp.GetString("Swagger.Debug.BasePath")
	}
	return s.vp.GetString("Swagger.Release.BasePath")
}

