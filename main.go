package main

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/vp"
	"github.com/Henry19910227/fitness-go/internal/v1/access"
	"github.com/Henry19910227/fitness-go/internal/v1/controller"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v2/router/action"
	"github.com/Henry19910227/fitness-go/internal/v2/router/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/router/banner"
	bodyImage "github.com/Henry19910227/fitness-go/internal/v2/router/body_image"
	body "github.com/Henry19910227/fitness-go/internal/v2/router/body_record"
	"github.com/Henry19910227/fitness-go/internal/v2/router/course"
	"github.com/Henry19910227/fitness-go/internal/v2/router/feedback"
	"github.com/Henry19910227/fitness-go/internal/v2/router/food"
	"github.com/Henry19910227/fitness-go/internal/v2/router/food_category"
	"github.com/Henry19910227/fitness-go/internal/v2/router/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/router/order"
	"github.com/Henry19910227/fitness-go/internal/v2/router/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/router/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/router/review"
	"github.com/Henry19910227/fitness-go/internal/v2/router/review_image"
	"github.com/Henry19910227/fitness-go/internal/v2/router/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/router/user"
	"github.com/Henry19910227/fitness-go/internal/v2/router/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/router/workout"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/router/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/router/workout_set_order"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var (
	rootPath string
)

var (
	mysqlTool     tool.Mysql
	gormTool      tool.Gorm
	viperTool     *viper.Viper
	migrateTool   tool.Migrate
	redisTool     tool.Redis
	jwtTool       tool.JWT
	logTool       tool.Logger
	otpTool       tool.OTP
	resTool       tool.Resource
	schedulerTool = cron.New(cron.WithSeconds())
)

var (
	logHandler    handler.Logger
	ssoHandler    handler.SSO
	uploadHandler handler.Uploader
	resHandler    handler.Resource
)

var (
	migrateService                         service.Migrate
	swagService                            service.Swagger
	loginService                           service.Login
	regService                             service.Register
	userService                            service.User
	trainerService                         service.Trainer
	courseService                          service.Course
	planService                            service.Plan
	workoutService                         service.Workout
	workoutSetService                      service.WorkoutSet
	actionService                          service.Action
	saleService                            service.Sale
	storeService                           service.Store
	reviewService                          service.Review
	paymentService                         service.Payment
	workoutLogService                      service.WorkoutLog
	favoriteService                        service.Favorite
	workoutSetLogService                   service.WorkoutSetLog
	orderService                           service.Order
	courseUsageStatisticService            service.CourseUsageStatistic
	userCourseUsageMonthlyStatisticService service.UserCourseUsageMonthlyStatistic
	userIncomeMonthlyStatisticService      service.UserIncomeMonthlyStatistic
	rdaService                             service.RDA
	dietService                            service.Diet
	foodCategoryService                    service.FoodCategory
	foodService                            service.Food
	mealService                            service.Meal
)

var (
	trainerAccess    access.Trainer
	courseAccess     access.Course
	planAccess       access.Plan
	workoutAccess    access.Workout
	workoutSetAccess access.WorkoutSet
	actionAccess     access.Action
)

var (
	userMiddleware     gin.HandlerFunc
	trainerMiddleware  gin.HandlerFunc
	adminLV1Middleware gin.HandlerFunc
	adminLV2Middleware gin.HandlerFunc

	userMidd    middleware.User
	courseMidd  middleware.Course
	planMidd    middleware.Plan
	reviewMidd  middleware.Review
	workoutMidd middleware.Workout
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
	planMidd = middleware.NewPlanMiddleware(viperTool, gormTool)
	reviewMidd = middleware.NewReviewMiddleware(viperTool, gormTool)
	workoutMidd = middleware.NewWorkoutMiddleware(viperTool, gormTool)
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
	router.Use(middleware.CORS()) //加入解決跨域中間層
	//gin.SetMode(gin.ReleaseMode)
	baseGroup := router.Group("/api")

	//v1
	v1 := baseGroup.Group("/v1")
	controller.NewMigrate(v1, migrateService, adminLV2Middleware)
	controller.NewRegister(v1, regService)
	controller.NewLogin(v1, loginService, userMiddleware, adminLV1Middleware)
	controller.NewUser(v1, userService, userMiddleware)
	controller.NewTrainer(v1, trainerService, courseService, userMiddleware, userMidd)
	controller.NewCourse(v1, courseService, planService, actionService, reviewService, userMidd, courseMidd)
	controller.NewCourseProduct(v1, courseService, planService, workoutSetService, courseMidd, userMidd)
	controller.NewCourseAsset(v1, courseService, planService, userMidd, courseMidd)
	controller.NewCourseStatistic(v1, courseService, userMidd)
	controller.NewPlan(v1, planService, workoutService, workoutSetAccess, userMidd, courseMidd)
	controller.NewPlanProduct(v1, planService, workoutService, planMidd, userMidd)
	controller.NewPlanAsset(v1, planService, workoutService, planMidd, userMidd)
	controller.NewWorkout(v1, workoutService, workoutSetService, userMidd, courseMidd)
	controller.NewWorkoutProduct(v1, workoutService, workoutSetService, workoutLogService, workoutMidd, userMidd)
	controller.NewWorkoutAsset(v1, workoutService, workoutSetService, workoutLogService, workoutMidd, userMidd)
	controller.NewWorkoutSet(v1, workoutSetService, userMidd, courseMidd)
	controller.NewWorkoutLog(v1, workoutLogService, userMidd)
	controller.NewAction(v1, actionService, workoutSetLogService, userMidd, courseMidd)
	controller.NewSale(v1, saleService, userMidd)
	controller.NewStore(v1, storeService, courseService, planService, workoutService, workoutSetService, courseMidd, planMidd)
	controller.NewReview(v1, courseService, reviewService, userMidd, courseMidd, reviewMidd)
	controller.NewPayment(v1, paymentService, courseService, userMidd)
	controller.NewFavorite(v1, favoriteService, courseService, userMidd, courseMidd)
	controller.NewCMSLogin(v1, loginService, userMidd)
	controller.NewCMSUser(v1, userService, userMidd)
	controller.NewCMSTrainer(v1, trainerService, courseService, userMidd)
	controller.NewOrder(v1, orderService, userMidd)
	controller.NewStatistic(v1, userIncomeMonthlyStatisticService, userCourseUsageMonthlyStatisticService, userMidd)
	controller.NewRDA(v1, rdaService, userMidd)
	controller.NewDiet(v1, dietService, userMidd)
	controller.NewFoodCategory(v1, foodCategoryService, userMidd)
	controller.NewFood(v1, foodService, userMidd)
	controller.NewMeal(v1, mealService, userMidd)
	controller.NewScheduler(schedulerTool, courseUsageStatisticService, userCourseUsageMonthlyStatisticService, userIncomeMonthlyStatisticService)
	controller.NewSwagger(router, swagService)
	controller.NewHealthy(router)
	schedulerTool.Start()

	// v2
	v2 := baseGroup.Group("/v2")
	course.SetRoute(v2)
	plan.SetRoute(v2)
	workout.SetRoute(v2)
	workoutSet.SetRoute(v2)
	workout_set_order.SetRoute(v2)
	food.SetRoute(v2)
	food_category.SetRoute(v2)
	meal.SetRoute(v2)
	trainer.SetRoute(v2)
	action.SetRoute(v2)
	body.SetRoute(v2)
	bodyImage.SetRoute(v2)
	feedback.SetRoute(v2)
	user.SetRoute(v2)
	order.SetRoute(v2)
	receipt.SetRoute(v2)
	review.SetRoute(v2)
	review_image.SetRoute(v2)
	banner.SetRoute(v2)
	user_subscribe_info.SetRoute(v2)
	bank_account.SetRoute(v2)
	router.Run(":" + vp.Shared().GetString("Server.HttpPort"))
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
	fullPath, err := filepath.Abs("./config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	vp := viper.New()
	vp.SetConfigFile(fullPath)
	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf(err.Error())
	}
	vp.Set("Server.RunMode", build.RunMode())
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

func setupGormTool() {
	setting := setting.NewMysql(viperTool)
	tool, err := tool.NewGorm(setting)
	if err != nil {
		log.Fatalf(err.Error())
	}
	gormTool = tool
}

func setupMigrateTool() {
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
	setupWorkoutSetService()
	courseService = service.NewCourseService(viperTool, gormTool)
	planService = service.NewPlanService(viperTool, gormTool)
	workoutService = service.NewWorkoutService(viperTool, gormTool)
	trainerService = service.NewTrainerService(viperTool, gormTool)
	actionService = service.NewActionService(viperTool, gormTool)
	loginService = service.NewLoginService(viperTool, gormTool)
	storeService = service.NewStoreService(viperTool, gormTool)
	reviewService = service.NewReviewService(viperTool, gormTool)
	paymentService = service.NewPaymentService(viperTool, gormTool)
	saleService = service.NewSaleService(viperTool, gormTool)
	userService = service.NewUserService(viperTool, gormTool)
	workoutLogService = service.NewWorkoutLogService(viperTool, gormTool)
	favoriteService = service.NewFavoriteService(viperTool, gormTool)
	workoutSetLogService = service.NewWorkoutSetLogService(viperTool, gormTool)
	orderService = service.NewOrderService(viperTool, gormTool)
	courseUsageStatisticService = service.NewCourseUsageStatisticService(viperTool, gormTool)
	userCourseUsageMonthlyStatisticService = service.NewUserCourseUsageMonthlyStatisticService(viperTool, gormTool)
	userIncomeMonthlyStatisticService = service.NewUserIncomeMonthlyStatisticService(viperTool, gormTool)
	rdaService = service.NewRDAService(viperTool, gormTool)
	dietService = service.NewDietService(viperTool, gormTool)
	foodCategoryService = service.NewFoodCategoryService(viperTool, gormTool)
	foodService = service.NewFoodService(viperTool, gormTool)
	mealService = service.NewMealService(viperTool, gormTool)
}

func setupMigrateService() {
	migrateService = service.NewMigrate(migrateTool, errcode.NewHandler())
}

func setupRegService() {
	userRepo := repository.NewUser(gormTool)
	regService = service.NewRegister(userRepo, logHandler, jwtTool, otpTool, viperTool, errcode.NewHandler())
}

func setupWorkoutSetService() {
	workoutSetRepo := repository.NewWorkoutSet(gormTool)
	workoutSetService = service.NewWorkoutSet(workoutSetRepo, uploadHandler, resHandler, logHandler, jwtTool, errcode.NewHandler())
}

func setupSwagService() {
	swagService = service.NewSwagger(setting.NewSwagger(viperTool))
}

/** Access */
func setupAccess() {
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
