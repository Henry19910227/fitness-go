package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
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
	ssoHandler := handler.NewSSO(jwtTool, redisTool, setting.NewUser(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewLogin(adminRepo, userRepo, trainerRepo, albumRepo, cerRepo, ssoHandler, logger, jwtTool, errHandler)
}

func NewCourseService(viperTool *viper.Viper, gormTool tool.Gorm) Course {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	courseRepo := repository.NewCourse(gormTool)
	userCourseAsset := repository.NewUserCourseAsset(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	planRepo := repository.NewPlan(gormTool)
	saleRepo := repository.NewSale(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewCourse(courseRepo, userCourseAsset, trainerRepo, planRepo, saleRepo, uploader, resHandler, logger, jwtTool, errHandler)
}

func NewPlanService(viperTool *viper.Viper, gormTool tool.Gorm) Plan {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	courseRepo := repository.NewCourse(gormTool)
	planRepo := repository.NewPlan(gormTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewPlan(planRepo, courseRepo, logger, jwtTool, errHandler)
}

func NewWorkoutService(viperTool *viper.Viper, gormTool tool.Gorm) Workout {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	courseRepo := repository.NewCourse(gormTool)
	workoutRepo := repository.NewWorkout(gormTool)
	workoutSetRepo := repository.NewWorkoutSet(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewWorkout(courseRepo, workoutRepo, workoutSetRepo, resHandler, uploader, logger, jwtTool, errHandler)
}

func NewTrainerService(viperTool *viper.Viper, gormTool tool.Gorm) Trainer {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	trainerRepo := repository.NewTrainer(gormTool)
	albumRepo := repository.NewTrainerAlbum(gormTool)
	cerRepo := repository.NewCertificate(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewTrainer(trainerRepo, albumRepo, cerRepo, uploader, resHandler, logger, jwtTool, errHandler)
}

func NewActionService(viperTool *viper.Viper, gormTool tool.Gorm) Action {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	actionRepo := repository.NewAction(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewAction(actionRepo, courseRepo, uploader, resHandler, errHandler)
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
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewReview(reviewRepo, uploader, resHandler, errHandler)
}

func NewPaymentService(viperTool *viper.Viper, gormTool tool.Gorm) Payment {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	reqTool := tool.NewRequest()

	orderRepo := repository.NewOrder(gormTool)
	saleRepo := repository.NewSale(gormTool)
	subscribePlanRepo := repository.NewSubscribePlan(gormTool)
	courseRepo := repository.NewCourse(gormTool)
	receiptRepo := repository.NewReceipt(gormTool)
	purchaseRepo := repository.NewUserCourseAsset(gormTool)
	subscribeLogRepo := repository.NewSubscribeLog(gormTool)
	purchaseLogRepo := repository.NewPurchaseLog(gormTool)
	memberRepo := repository.NewSubscribeInfo(gormTool)
	transactionRepo := repository.NewTransaction(gormTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewPayment(orderRepo, saleRepo, subscribePlanRepo, courseRepo, receiptRepo, purchaseRepo, subscribeLogRepo, purchaseLogRepo, memberRepo, transactionRepo, reqTool, jwtTool, errHandler)
}
