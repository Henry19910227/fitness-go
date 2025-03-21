package service

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/spf13/viper"
)

func NewLoginService(viperTool *viper.Viper, gormTool tool.Gorm) Login {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	redisTool := tool.NewRedis(setting.NewRedis(viperTool))
	adminRepo := repository.NewAdmin(gormTool)
	userRepo := repository.NewUser(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	albumRepo := repository.NewTrainerAlbum(gormTool)
	cerRepo := repository.NewCertificate(gormTool)
	subscribeInfoRepo := repository.NewSubscribeInfo(gormTool)
	ssoHandler := handler.NewSSO(jwtTool, redisTool, setting.NewUser(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewLogin(adminRepo, userRepo, trainerRepo, albumRepo, cerRepo, subscribeInfoRepo, ssoHandler, logger, jwtTool, errHandler)
}

func NewCourseService(viperTool *viper.Viper, gormTool tool.Gorm) Course {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	courseRepo := repository.NewCourse(gormTool)
	userCourseAsset := repository.NewUserCourseAsset(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	trainerStatRepo := repository.NewTrainerStatistic(gormTool)
	planRepo := repository.NewPlan(gormTool)
	workoutRepo := repository.NewWorkout(gormTool)
	workoutSetRepo := repository.NewWorkoutSet(gormTool)
	saleRepo := repository.NewSale(gormTool)
	subscribeInfoRepo := repository.NewSubscribeInfo(gormTool)
	userCourseStatisticRepo := repository.NewUserCourseStatistic(gormTool)
	favoriteRepo := repository.NewFavorite(gormTool)
	reviewStatisticRepo := repository.NewReviewStatistic(gormTool)
	courseUsageStatisticRepo := repository.NewCourseUsageStatistic(gormTool)
	transactionRepo := repository.NewTransaction(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewCourse(courseRepo, userCourseAsset, trainerRepo, trainerStatRepo,
		planRepo, workoutRepo, workoutSetRepo, saleRepo, subscribeInfoRepo,
		userCourseStatisticRepo, favoriteRepo, reviewStatisticRepo,
		courseUsageStatisticRepo, transactionRepo, uploader, resHandler, logger, jwtTool, errHandler)
}

func NewPlanService(viperTool *viper.Viper, gormTool tool.Gorm) Plan {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	courseRepo := repository.NewCourse(gormTool)
	planRepo := repository.NewPlan(gormTool)
	planStatisticRepo := repository.NewUserPlanStatistic(gormTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewPlan(planRepo, courseRepo, planStatisticRepo, logger, jwtTool, errHandler)
}

func NewWorkoutService(viperTool *viper.Viper, gormTool tool.Gorm) Workout {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	courseRepo := repository.NewCourse(gormTool)
	workoutRepo := repository.NewWorkout(gormTool)
	workoutSetRepo := repository.NewWorkoutSet(gormTool)
	workoutLogRepo := repository.NewWorkoutLog(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewWorkout(courseRepo, workoutRepo, workoutSetRepo, workoutLogRepo, resHandler, uploader, logger, jwtTool, errHandler)
}

func NewTrainerService(viperTool *viper.Viper, gormTool tool.Gorm) Trainer {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	trainerRepo := repository.NewTrainer(gormTool)
	albumRepo := repository.NewTrainerAlbum(gormTool)
	cerRepo := repository.NewCertificate(gormTool)
	favoriteRepo := repository.NewFavorite(gormTool)
	cardRepo := repository.NewCard(gormTool)
	bankRepo := repository.NewBankAccount(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewTrainer(trainerRepo, albumRepo, cerRepo, favoriteRepo, cardRepo, bankRepo, uploader, resHandler, logger, jwtTool, errHandler)
}

func NewActionService(viperTool *viper.Viper, gormTool tool.Gorm) Action {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	actionRepo := repository.NewAction(gormTool)
	actionPRRepo := repository.NewActionPR(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewAction(actionRepo, actionPRRepo, courseRepo, uploader, resHandler, errHandler)
}

func NewStoreService(viperTool *viper.Viper, gormTool tool.Gorm) Store {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))

	courseRepo := repository.NewCourse(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	reviewRepo := repository.NewReview(gormTool)

	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewStore(courseRepo, trainerRepo, reviewRepo, errHandler)
}

func NewReviewService(viperTool *viper.Viper, gormTool tool.Gorm) Review {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))

	reviewRepo := repository.NewReview(gormTool)
	reviewImageRepo := repository.NewReviewImage(gormTool)
	reviewStatRepo := repository.NewReviewStatistic(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	trainerStatisticRepo := repository.NewTrainerStatistic(gormTool)
	transactionRepo := repository.NewTransaction(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewReview(reviewRepo, reviewImageRepo, reviewStatRepo, courseRepo, trainerStatisticRepo, transactionRepo, uploader, resHandler, errHandler)
}

func NewPaymentService(viperTool *viper.Viper, gormTool tool.Gorm) Payment {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	reqTool := tool.NewRequest()
	iapTool := tool.NewIAP(setting.NewIAP(viperTool))
	iabTool := tool.NewIAB(setting.NewIAB(viperTool))

	userRepo := repository.NewUser(gormTool)
	orderRepo := repository.NewOrder(gormTool)
	orderSPrepo := repository.NewOrderSubscribePlan(gormTool)
	saleRepo := repository.NewSale(gormTool)
	subscribePlanRepo := repository.NewSubscribePlan(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	receiptRepo := repository.NewReceipt(gormTool)
	purchaseRepo := repository.NewUserCourseAsset(gormTool)
	subscribeLogRepo := repository.NewSubscribeLog(gormTool)
	purchaseLogRepo := repository.NewPurchaseLog(gormTool)
	memberRepo := repository.NewSubscribeInfo(gormTool)
	transactionRepo := repository.NewTransaction(gormTool)
	iapHandler := handler.NewIAP(iapTool, reqTool)
	iabHandler := handler.NewIAB(iabTool, reqTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewPayment(userRepo, orderRepo, orderSPrepo, saleRepo, subscribePlanRepo,
		courseRepo, receiptRepo, purchaseRepo, subscribeLogRepo,
		purchaseLogRepo, memberRepo, transactionRepo,
		iapHandler, iabHandler, reqTool, jwtTool, errHandler)
}

func NewSaleService(viperTool *viper.Viper, gormTool tool.Gorm) Sale {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	saleRepo := repository.NewSale(gormTool)
	subscribePlanRepo := repository.NewSubscribePlan(gormTool)
	return NewSale(saleRepo, subscribePlanRepo, jwtTool, errHandler)
}

func NewUserService(viperTool *viper.Viper, gormTool tool.Gorm) User {
	userRepo := repository.NewUser(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	subscribeInfoRepo := repository.NewSubscribeInfo(gormTool)
	albumRepo := repository.NewTrainerAlbum(gormTool)
	cerRepo := repository.NewCertificate(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewUser(userRepo, trainerRepo, subscribeInfoRepo, albumRepo, cerRepo, uploader, resHandler, logger, jwtTool, errHandler)
}

func NewWorkoutLogService(viperTool *viper.Viper, gormTool tool.Gorm) WorkoutLog {
	workoutLogRepo := repository.NewWorkoutLog(gormTool)
	workoutSetLogRepo := repository.NewWorkoutSetLog(gormTool)
	actionPRRepo := repository.NewActionPR(gormTool)
	transactionRepo := repository.NewTransaction(gormTool)
	workoutSetRepo := repository.NewWorkoutSet(gormTool)
	courseAssetRepo := repository.NewUserCourseAsset(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	subscribeInfoRepo := repository.NewSubscribeInfo(gormTool)
	courseStatisticRepo := repository.NewUserCourseStatistic(gormTool)
	planStatisticRepo := repository.NewUserPlanStatistic(gormTool)
	trainerStatisticRepo := repository.NewTrainerStatistic(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewWorkoutLog(workoutLogRepo, workoutSetLogRepo, workoutSetRepo, actionPRRepo, courseRepo,
		courseAssetRepo, subscribeInfoRepo, courseStatisticRepo,
		planStatisticRepo, trainerStatisticRepo, transactionRepo, errHandler)
}

func NewFavoriteService(viperTool *viper.Viper, gormTool tool.Gorm) Favorite {
	favoriteRepo := repository.NewFavorite(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewFavorite(favoriteRepo, errHandler)
}

func NewWorkoutSetLogService(viperTool *viper.Viper, gormTool tool.Gorm) WorkoutSetLog {
	workoutSetLogRepo := repository.NewWorkoutSetLog(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewWorkoutSetLog(workoutSetLogRepo, errHandler)
}

func NewOrderService(viperTool *viper.Viper, gormTool tool.Gorm) Order {
	orderRepo := repository.NewOrder(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewOrder(orderRepo, errHandler)
}

func NewCourseUsageStatisticService(viperTool *viper.Viper, gormTool tool.Gorm) CourseUsageStatistic {
	transactionRepo := repository.NewTransaction(gormTool)
	statisticRepo := repository.NewCourseUsageStatistic(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewCourseUsageStatistic(transactionRepo, statisticRepo, errHandler)
}

func NewUserCourseUsageMonthlyStatisticService(viperTool *viper.Viper, gormTool tool.Gorm) UserCourseUsageMonthlyStatistic {
	transactionRepo := repository.NewTransaction(gormTool)
	statisticRepo := repository.NewUserCourseUsageMonthlyStatistic(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewUserCourseUsageMonthlyStatistic(transactionRepo, statisticRepo, errHandler)
}

func NewUserIncomeMonthlyStatisticService(viperTool *viper.Viper, gormTool tool.Gorm) UserIncomeMonthlyStatistic {
	transactionRepo := repository.NewTransaction(gormTool)
	statisticRepo := repository.NewUserIncomeMonthlyStatistic(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewUserIncomeMonthlyStatistic(transactionRepo, statisticRepo, errHandler)
}

func NewRDAService(viperTool *viper.Viper, gormTool tool.Gorm) RDA {
	rdaRepo := repository.NewRDA(gormTool)
	dietRepo := repository.NewDiet(gormTool)
	transactionRepo := repository.NewTransaction(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewRDA(rdaRepo, dietRepo, transactionRepo, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), errHandler)
}

func NewDietService(viperTool *viper.Viper, gormTool tool.Gorm) Diet {
	dietRepo := repository.NewDiet(gormTool)
	rdaRepo := repository.NewRDA(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewDiet(dietRepo, rdaRepo, errHandler)
}

func NewFoodCategoryService(viperTool *viper.Viper, gormTool tool.Gorm) FoodCategory {
	foodCategoryRepo := repository.NewFoodCategory(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewFoodCategory(foodCategoryRepo, errHandler)
}

func NewFoodService(viperTool *viper.Viper, gormTool tool.Gorm) Food {
	foodRepo := repository.NewFood(gormTool)
	foodCategoryRepo := repository.NewFoodCategory(gormTool)
	calorieTool := tool.NewCalorie()
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewFood(foodRepo, foodCategoryRepo, calorieTool, errHandler)
}

func NewMealService(viperTool *viper.Viper, gormTool tool.Gorm) Meal {
	mealRepo := repository.NewMeal(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewMeal(mealRepo, errHandler)
}
