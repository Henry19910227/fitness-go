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

func NewReviewService(viperTool *viper.Viper, gormTool tool.Gorm) Review {
	courseRepo := repository.NewCourse(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	loggerTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(loggerTool, jwtTool))
	return NewReview(courseRepo, errHandler)
}

func NewCourseService(viperTool *viper.Viper, gormTool tool.Gorm) Course {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	courseRepo := repository.NewCourse(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	resTool := tool.NewResource(setting.NewResource(viperTool))
	uploader := handler.NewUploader(resTool, setting.NewUploadLimit(viperTool))
	resHandler := handler.NewResource(resTool)
	logTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	logger := handler.NewLogger(logTool, jwtTool)
	errHandler := errcode.NewErrHandler(handler.NewLogger(logTool, jwtTool))
	return NewCourse(courseRepo, trainerRepo, uploader, resHandler, logger, jwtTool, errHandler)
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
