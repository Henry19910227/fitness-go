package main

import (
	"flag"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/access"
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
	workoutSetService service.WorkoutSet
	actionService   service.Action
	saleService     service.Sale
	storeService    service.Store
	reviewService   service.Review
)

var (
	trainerAccess access.Trainer
	courseAccess access.Course
	planAccess access.Plan
	workoutAccess access.Workout
	workoutSetAccess access.WorkoutSet
	actionAccess access.Action
)

var (
	userMiddleware gin.HandlerFunc
	trainerMiddleware gin.HandlerFunc
	adminLV1Middleware  gin.HandlerFunc
	adminLV2Middleware  gin.HandlerFunc

	userMidd middleware.User
	courseMidd middleware.Course
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
	setupAccess()
	userMiddleware = middleware.UserJWT(ssoHandler, errcode.NewHandler())
	adminLV1Middleware = middleware.AdminLV1JWT(ssoHandler, errcode.NewHandler())
	adminLV2Middleware = middleware.AdminLV2JWT(ssoHandler, errcode.NewHandler())

	userMidd = middleware.NewUserMiddleware(viperTool, gormTool)
	courseMidd = middleware.NewCourseMiddleware(viperTool, gormTool)
}

// @title fitness api
// @description 健身平台 api

// @securityDefinitions.apikey fitness_token
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
	controller.NewTrainer(baseGroup, trainerService, userMiddleware, userMidd)
	controller.NewCourse(baseGroup, courseService, planService, actionService, reviewService, userMidd, courseMidd)
	controller.NewPlan(baseGroup, planService, workoutService, workoutSetAccess, userMidd, courseMidd)
	controller.NewWorkout(baseGroup, workoutService, workoutSetService, userMidd, courseMidd)
	controller.NewWorkoutSet(baseGroup, workoutSetService, userMidd, courseMidd)
	controller.NewAction(baseGroup, actionService, actionAccess, trainerAccess, userMidd, courseMidd)
	controller.NewSale(baseGroup, saleService, userMiddleware)
	controller.NewStore(baseGroup, storeService, courseService, courseMidd)
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
	resTool = tool.NewResource(setting.NewResource(viperTool))
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
	setupRegService()
	setupUserService()
	setupPlanService()
	setupWorkoutSetService()
	setupSaleService()
	courseService = service.NewCourseService(viperTool, gormTool)
	workoutService = service.NewWorkoutService(viperTool, gormTool)
	trainerService = service.NewTrainerService(viperTool, gormTool)
	actionService = service.NewActionService(viperTool, gormTool)
	loginService = service.NewLoginService(viperTool, gormTool)
	storeService = service.NewStoreService(viperTool, gormTool)
	reviewService = service.NewReviewService(viperTool, gormTool)
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

func setupPlanService()  {
	planRepo := repository.NewPlan(gormTool)
	planService = service.NewPlan(planRepo, logHandler, jwtTool, errcode.NewHandler())
}

func setupWorkoutSetService()  {
	workoutSetRepo := repository.NewWorkoutSet(gormTool)
	workoutSetService = service.NewWorkoutSet(workoutSetRepo, uploadHandler, resHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupSaleService()  {
	saleRepo := repository.NewSale(gormTool)
	saleService = service.NewSale(saleRepo, logHandler, jwtTool, errcode.NewHandler())
}

func setupSwagService()  {
	swagService = service.NewSwagger(setting.NewSwagger(viperTool))
}

/** Access */
func setupAccess()  {
	setupTrainerAccess()
	setupCourseAccess()
	setupPlanAccess()
	setupWorkoutAccess()
	setupWorkoutSetAccess()
	setupActionAccess()
}

func setupTrainerAccess() {
	trainerRepo := repository.NewTrainer(gormTool)
	trainerAccess = access.NewTrainer(trainerRepo, logHandler, jwtTool, errcode.NewHandler())
}

func setupCourseAccess() {
	courseRepo := repository.NewCourse(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	courseAccess = access.NewCourse(courseRepo, trainerRepo, logHandler, jwtTool, errcode.NewHandler())
}

func setupPlanAccess() {
	courseRepo := repository.NewCourse(gormTool)
	planAccess = access.NewPlan(courseRepo, logHandler, jwtTool, errcode.NewHandler())
}

func setupWorkoutAccess() {
	courseRepo := repository.NewCourse(gormTool)
	workoutAccess = access.NewWorkout(courseRepo, logHandler, jwtTool, errcode.NewHandler())
}

func setupWorkoutSetAccess() {
	courseRepo := repository.NewCourse(gormTool)
	logger := handler.NewLogger(logTool, jwtTool)
	workoutSetAccess = access.NewWorkoutSet(courseRepo, logHandler, jwtTool, errcode.NewErrHandler(logger))
}

func setupActionAccess() {
	courseRepo := repository.NewCourse(gormTool)
	actionAccess = access.NewAction(courseRepo, logHandler, jwtTool, errcode.NewHandler())
}


