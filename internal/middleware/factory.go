package middleware

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/spf13/viper"
)

func NewUserMiddleware(viperTool *viper.Viper, gormTool tool.Gorm) User {
	userRepo := repository.NewUser(gormTool)
	trainerRepo := repository.NewTrainer(gormTool)
	albumRepo := repository.NewTrainerAlbum(gormTool)
	cerRepo := repository.NewCertificate(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	redisTool := tool.NewRedis(setting.NewRedis(viperTool))
	loggerTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(loggerTool, jwtTool))
	return NewUser(userRepo, trainerRepo, albumRepo, cerRepo, jwtTool, redisTool, errHandler)
}

func NewCourseMiddleware(viperTool *viper.Viper, gormTool tool.Gorm) Course {
	courseRepo := repository.NewCourse(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	loggerTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(loggerTool, jwtTool))
	return NewCourse(courseRepo, jwtTool, errHandler)
}
