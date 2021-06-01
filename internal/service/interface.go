package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/coursedto"
	"github.com/Henry19910227/fitness-go/internal/dto/logindto"
	"github.com/Henry19910227/fitness-go/internal/dto/registerdto"
	"github.com/Henry19910227/fitness-go/internal/dto/userdto"
	"github.com/gin-gonic/gin"
)

type Migrate interface {
	Version() (uint, bool, errcode.Error)
	Up() (uint, bool, errcode.Error)
	UpStep(step int) (uint, bool, errcode.Error)
	Down() errcode.Error
	DownStep(step int) errcode.Error
	Force(version int) (uint, bool, errcode.Error)
	Migrate(version uint) (uint, bool, errcode.Error)
}

type Swagger interface {
	WrapHandler() gin.HandlerFunc
}

type Login interface {
	UserLoginByEmail(c *gin.Context, email string, password string) (*logindto.User, string, errcode.Error)
	AdminLoginByEmail(c *gin.Context, email string, password string) (*logindto.Admin, string, errcode.Error)
	UserLogoutByToken(c *gin.Context, token string) errcode.Error
	AdminLogoutByToken(c *gin.Context, token string) errcode.Error
}

type Register interface {
	SendEmailOTP(c *gin.Context, email string) (*registerdto.OTP, errcode.Error)
	EmailRegister(c *gin.Context, otp string, email string, nickname string, password string) (*registerdto.Register, errcode.Error)
	ValidateNicknameDup(c *gin.Context, nickname string) errcode.Error
	ValidateEmailDup(c *gin.Context, email string) errcode.Error
}

type User interface {
	UpdateUserByUID(c *gin.Context, uid int64, param *userdto.UpdateUserParam) (*userdto.User, errcode.Error)
	UpdateUserByToken(c *gin.Context, token string, param *userdto.UpdateUserParam) (*userdto.User, errcode.Error)
	GetUserByUID(c *gin.Context, uid int64) (*userdto.User, errcode.Error)
	GetUserByToken(c *gin.Context, token string) (*userdto.User, errcode.Error)
	CreateTrainer(c *gin.Context, uid int64, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error)
	CreateTrainerByToken(c *gin.Context, token string, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error)
	GetTrainerInfo(c *gin.Context, uid int64) (*userdto.Trainer, errcode.Error)
	GetTrainerInfoByToken(c *gin.Context, token string) (*userdto.Trainer, errcode.Error)
}

type Course interface {
	CreateCourseByToken(c *gin.Context, token string, param *coursedto.CreateCourseParam) (*coursedto.CreateResult, errcode.Error)
	CreateCourse(c *gin.Context, uid int64, param *coursedto.CreateCourseParam) (*coursedto.CreateResult, errcode.Error)
	UpdateCourse(c *gin.Context, courseID int64, param *coursedto.UpdateCourseParam) (*coursedto.Course, errcode.Error)
	GetCourseByTokenAndCourseID(c *gin.Context, token string, courseID int64) (*coursedto.Course, errcode.Error)
	GetCourseByID(c *gin.Context, courseID int64) (*coursedto.Course, errcode.Error)
	UploadCourseCoverByID(c *gin.Context, courseID int64, param *coursedto.UploadCourseCoverParam) (*coursedto.CourseCover, errcode.Error)
}