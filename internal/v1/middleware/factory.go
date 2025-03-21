package middleware

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
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

func NewPlanMiddleware(viperTool *viper.Viper, gormTool tool.Gorm) Plan {
	courseRepo := repository.NewCourse(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	loggerTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(loggerTool, jwtTool))
	return NewPlan(courseRepo, jwtTool, errHandler)
}


func NewReviewMiddleware(viperTool *viper.Viper, gormTool tool.Gorm) Review {
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	loggerTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(loggerTool, jwtTool))
	return NewReview(errHandler)
}

func NewWorkoutMiddleware(viperTool *viper.Viper, gormTool tool.Gorm) Workout {
	courseRepo := repository.NewCourse(gormTool)
	jwtTool := tool.NewJWT(setting.NewJWT(viperTool))
	loggerTool, _ := tool.NewLogger(setting.NewLogger(viperTool))
	errHandler := errcode.NewErrHandler(handler.NewLogger(loggerTool, jwtTool))
	return NewWorkout(courseRepo, jwtTool, errHandler)
}
