package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/spf13/viper"
)

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
