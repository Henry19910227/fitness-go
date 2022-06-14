package service

import (
	"github.com/Henry19910227/fitness-go/docs"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type swagger struct {
	setting setting.Swagger
}

func NewSwagger (setting setting.Swagger) Swagger {
	docs.SwaggerInfo.Version = setting.GetVersion()
	docs.SwaggerInfo.Host = setting.GetHost()
	docs.SwaggerInfo.BasePath = setting.GetBasePath()
	return &swagger{setting}
}

func (s *swagger) WrapHandler() gin.HandlerFunc {
	url := s.setting.GetProtocol() + "://" + s.setting.GetHost() + "/api/swagger/doc.json"
	return ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(url))
}
