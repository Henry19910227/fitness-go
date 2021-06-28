package main

import (
	"flag"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/controller"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	rootPath   string
)

var (
	mysqlTool   tool.Mysql
	gormTool    tool.Gorm
	viperTool   *viper.Viper
	migrateTool tool.Migrate
	redisTool   tool.Redis
	jwtTool     tool.JWT
	logTool     tool.Logger
	otpTool     tool.OTP
	resTool     tool.Resource
)

var (
	logHandler  handler.Logger
	ssoHandler  handler.SSO
	uploadHandler handler.Uploader
	resHandler handler.Resource
)

var (
	migrateService  service.Migrate
	swagService     service.Swagger
	loginService    service.Login
	regService      service.Register
	userService     service.User
	trainerService  service.Trainer
	courseService   service.Course
	planService     service.Plan
	workoutService  service.Workout
	actionService   service.Action
)

var (
	userMiddleware gin.HandlerFunc
	trainerMiddleware gin.HandlerFunc
	adminLV1Middleware  gin.HandlerFunc
	adminLV2Middleware  gin.HandlerFunc
)

func init() {
	os.Setenv("TZ", "Asia/Taipei")
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		rootPath = path.Dir(filename)
	}
	setupTool()
	setupHandler()
	setupService()
	userMiddleware = middleware.UserJWT(ssoHandler, errcode.NewHandler())
	adminLV1Middleware = middleware.AdminLV1JWT(ssoHandler, errcode.NewHandler())
	adminLV2Middleware = middleware.AdminLV2JWT(ssoHandler, errcode.NewHandler())
}

// @title fitness api
// @description 健身平台 api

// @securityDefinitions.apikey fitness_user_token
// @in header
// @name Token

// @securityDefinitions.apikey fitness_admin_token
// @in header
// @name Token


func main() {
	router := gin.New()
	router.Use(gin.Logger()) //加入路由Logger
	router.Use(gin.CustomRecovery(middleware.Recover(logHandler)))
	baseGroup := router.Group("/api/v1")
	controller.NewMigrate(baseGroup, migrateService, adminLV2Middleware)
	controller.NewManager(baseGroup)
	controller.NewRegister(baseGroup, regService)
	controller.NewLogin(baseGroup, loginService, userMiddleware, adminLV1Middleware)
	controller.NewUser(baseGroup, userService, userMiddleware)
	controller.NewTrainer(baseGroup, trainerService, userMiddleware)
	controller.NewCourse(baseGroup, courseService, planService, actionService, userMiddleware)
	controller.NewPlan(baseGroup, planService, workoutService, userMiddleware)
	controller.NewWorkout(baseGroup, workoutService, userMiddleware)
	controller.NewAction(baseGroup, actionService, userMiddleware)
	controller.NewSwagger(router, swagService)
	controller.NewHealthy(router)

	router.Run(":"+viperTool.GetString("Server.HttpPort"))
}

/** Tool */
func setupTool() {
	setupViper()
	setupLogTool()
	setupMysqlTool()
	setupMigrateTool()
	setupGormTool()
	jwtTool = tool.NewJWT(setting.NewJWT(viperTool))
	redisTool = tool.NewRedis(setting.NewRedis(viperTool))
	otpTool = tool.NewOTP()
	resTool = tool.NewFile(setting.NewUploader(viperTool))
}

func setupLogTool() {
	logSetting := setting.NewLogger(viperTool)
	logger, err := tool.NewLogger(logSetting)
	if err != nil {
		log.Fatalf(err.Error())
	}
	logTool = logger
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

func setupGormTool()  {
	setting := setting.NewMysql(viperTool)
	tool, err := tool.NewGorm(setting)
	if err != nil {
		log.Fatalf(err.Error())
	}
	gormTool = tool
}


func setupMigrateTool()  {
	mysqlSetting := setting.NewMysql(viperTool)
	migSetting := setting.NewMigrate(rootPath)
	migrate := tool.NewMigrate(mysqlSetting, migSetting)
	migrateTool = migrate
}

/** Handler */
func setupHandler() {
	logHandler = handler.NewLogger(logTool, jwtTool)
	ssoHandler = handler.NewSSO(jwtTool, redisTool, setting.NewUser(viperTool))
	uploadHandler = handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler = handler.NewResource(resTool)
}

/** Service */
func setupService() {
	setupMigrateService()
	setupSwagService()
	setupLoginService()
	setupRegService()
	setupUserService()
	setupTrainerService()
	setupCourseService()
	setupPlanService()
	setupActionService()
	setupWorkoutService()
}

func setupLoginService() {
	adminRepo := repository.NewAdmin(gormTool)
	userRepo := repository.NewUser(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	loginService = service.NewLogin(adminRepo, userRepo, trainerRepo, ssoHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupMigrateService()  {
	migrateService = service.NewMigrate(migrateTool, errcode.NewHandler())
}

func setupRegService()  {
	userRepo := repository.NewUser(gormTool)
	regService = service.NewRegister(userRepo, logHandler, jwtTool, otpTool, viperTool, errcode.NewHandler())
}

func setupUserService()  {
	userRepo := repository.NewUser(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	userService = service.NewUser(userRepo, trainerRepo, uploadHandler, resHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupTrainerService()  {
	trainerRepo := repository.NewTrainer(gormTool)
	trainerService = service.NewTrainer(trainerRepo, uploadHandler, resHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupCourseService()  {
	courseRepo := repository.NewCourse(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	courseService = service.NewCourse(courseRepo, trainerRepo, uploadHandler, resHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupPlanService()  {
	planRepo := repository.NewPlan(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	planService = service.NewPlan(planRepo, courseRepo, logHandler, jwtTool, errcode.NewHandler())
}

func setupWorkoutService()  {
	workoutRepo := repository.NewWorkout(gormTool)
	planRepo := repository.NewPlan(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	workoutService = service.NewWorkout(workoutRepo, planRepo, courseRepo, uploadHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupActionService()  {
	actionRepo := repository.NewAction(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	actionService = service.NewAction(actionRepo, courseRepo, uploadHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupSwagService()  {
	swagService = service.NewSwagger(setting.NewSwagger(viperTool))
}

