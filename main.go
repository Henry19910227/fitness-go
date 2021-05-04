package main

import (
	"flag"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/controller"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
)

var (
	rootPath   string
)

var (
	mysqlTool   tool.Mysql
	viperTool   *viper.Viper
	migrateTool tool.Migrate
	redisTool   tool.Redis
	jwtTool     tool.JWT
)

var (
	ssoHandler  handler.SSO
)

var (
	migrateService  service.Migrate
	swagService     service.Swagger
)

var (
	userMiddleware gin.HandlerFunc
	trainerMiddleware gin.HandlerFunc
	adminLV1Middleware  gin.HandlerFunc
	adminLV2Middleware  gin.HandlerFunc
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		rootPath = path.Dir(filename)
	}
	setupTool()
	setupHandler()
	setupService()
	userMiddleware = middleware.UserJWT(ssoHandler, errcode.NewCommon())
	trainerMiddleware = middleware.TrainerJWT(ssoHandler, errcode.NewCommon())
	adminLV1Middleware = middleware.AdminLV1JWT(ssoHandler, errcode.NewCommon())
	adminLV2Middleware = middleware.AdminLV2JWT(ssoHandler, errcode.NewCommon())
}


func main() {
	router := gin.New()
	router.Use(gin.Logger()) //加入路由Logger
	baseGroup := router.Group("/api/v1")
	controller.NewMigrate(baseGroup, migrateService, adminLV2Middleware)
	controller.NewSwaggerController(router, swagService)
	router.Run(":"+viperTool.GetString("Server.HttpPort"))
}

/** Tool */
func setupTool() {
	setupViper()
	setupMysqlTool()
	setupMigrateTool()
	jwtTool = tool.NewJWT(setting.NewJWT(viperTool))
	redisTool = tool.NewRedis(setting.NewRedis(viperTool))
}

func setupViper() {
	vp := viper.New()
	vp.SetConfigFile("./config/config.yaml")
	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf(err.Error())
	}
	var mode string
	flag.StringVar(&mode, "m", "debug", "獲取運行模式")
	flag.Parse()
	vp.Set("Server.RunMode", mode)
	viperTool = vp
}

func setupMysqlTool() {
	setting := setting.NewMysql(viperTool)
	tool, err := tool.NewMysql(setting)
	if err != nil {
		log.Fatalf(err.Error())
	}
	mysqlTool = tool
}


func setupMigrateTool()  {
	mysqlSetting := setting.NewMysql(viperTool)
	migSetting := setting.NewMigrate(rootPath)
	migrate := tool.NewMigrate(mysqlSetting, migSetting)
	migrateTool = migrate
}

/** Handler */
func setupHandler() {
	ssoHandler = handler.NewSSO(jwtTool, redisTool, setting.NewUser(viperTool))
}

/** Service */
func setupService() {
	setupMigrateService()
	setupSwagService()
}

func setupMigrateService()  {
	migrateService = service.NewMigrate(migrateTool, errcode.NewCommon())
}

func setupSwagService()  {
	swagService = service.NewSwagger(setting.NewSwagger(viperTool))
}

