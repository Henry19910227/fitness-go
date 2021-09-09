package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/dto/plandto"
	"github.com/Henry19910227/fitness-go/internal/dto/registerdto"
	"github.com/Henry19910227/fitness-go/internal/dto/saledto"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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
	UserLoginByEmail(c *gin.Context, email string, password string) (*dto.User, string, errcode.Error)
	AdminLoginByEmail(c *gin.Context, email string, password string) (*dto.Admin, string, errcode.Error)
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
	UpdateUserByUID(c *gin.Context, uid int64, param *dto.UpdateUserParam) (*dto.User, errcode.Error)
	UpdateUserByToken(c *gin.Context, token string, param *dto.UpdateUserParam) (*dto.User, errcode.Error)
	GetUserByUID(c *gin.Context, uid int64) (*dto.User, errcode.Error)
	GetUserByToken(c *gin.Context, token string) (*dto.User, errcode.Error)
	UploadUserAvatarByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.UserAvatar, errcode.Error)
	UploadUserAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*dto.UserAvatar, errcode.Error)
}

type Trainer interface {
	CreateTrainer(c *gin.Context, uid int64, param *dto.CreateTrainerParam) (*dto.Trainer, errcode.Error)
	GetTrainerInfo(c *gin.Context, uid int64) (*dto.Trainer, errcode.Error)
	GetTrainerInfoByToken(c *gin.Context, token string) (*dto.Trainer, errcode.Error)
	UpdateTrainer(c *gin.Context, uid int64, param *dto.UpdateTrainerParam) (*dto.Trainer, errcode.Error)
	UploadTrainerAvatarByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerAvatar, errcode.Error)
	UploadTrainerAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*dto.TrainerAvatar, errcode.Error)
}

type Course interface {
	CreateCourseByToken(c *gin.Context, token string, param *dto.CreateCourseParam) (*dto.Course, errcode.Error)
	CreateCourse(c *gin.Context, uid int64, param *dto.CreateCourseParam) (*dto.Course, errcode.Error)
	UpdateCourse(c *gin.Context, courseID int64, param *dto.UpdateCourseParam) (*dto.Course, errcode.Error)
	DeleteCourse(c *gin.Context, courseID int64) (*dto.CourseID, errcode.Error)
	GetCourseSummariesByToken(c *gin.Context, token string, status *int) ([]*dto.CourseSummary, errcode.Error)
	GetCourseSummariesByUID(c *gin.Context, uid int64, status *int) ([]*dto.CourseSummary, errcode.Error)
	GetCourseDetailByCourseID(c *gin.Context, courseID int64) (*dto.Course, errcode.Error)
	UploadCourseCoverByID(c *gin.Context, courseID int64, param *dto.UploadCourseCoverParam) (*dto.CourseCover, errcode.Error)
}

type Plan interface {
	CreatePlan(c *gin.Context, courseID int64, name string) (*plandto.Plan, errcode.Error)
	UpdatePlan(c *gin.Context, planID int64, name string) (*plandto.Plan, errcode.Error)
	DeletePlan(c *gin.Context, planID int64) (*plandto.PlanID, errcode.Error)
	GetPlansByCourseID(c *gin.Context, courseID int64) ([]*plandto.Plan, errcode.Error)
}

type Workout interface {
	CreateWorkout(c *gin.Context, planID int64, name string) (*dto.Workout, errcode.Error)
	GetWorkoutsByPlanID(c *gin.Context, planID int64) ([]*dto.Workout, errcode.Error)
	UpdateWorkout(c *gin.Context, workoutID int64, param *dto.UpdateWorkoutParam) (*dto.Workout, errcode.Error)
	DeleteWorkout(c *gin.Context, workoutID int64) (*dto.WorkoutID, errcode.Error)
	UploadWorkoutStartAudio(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	UploadWorkoutEndAudio(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	CreateWorkoutByTemplate(c *gin.Context, planID int64, name string, workoutTemplateID int64) (*dto.Workout, errcode.Error)
	DeleteWorkoutStartAudio(c *gin.Context, workoutID int64) errcode.Error
	DeleteWorkoutEndAudio(c *gin.Context, workoutID int64) errcode.Error
}

type WorkoutSet interface {
	CreateRestSet(c *gin.Context, workoutID int64) (*dto.WorkoutSet, errcode.Error)
	CreateWorkoutSets(c *gin.Context, workoutID int64, actionIDs []int64) ([]*dto.WorkoutSet, errcode.Error)
	DuplicateWorkoutSets(c *gin.Context, setID int64, count int) ([]*dto.WorkoutSet, errcode.Error)
	GetWorkoutSets(c *gin.Context, workoutID int64) ([]*dto.WorkoutSet, errcode.Error)
	UpdateWorkoutSet(c *gin.Context, setID int64, param *dto.UpdateWorkoutSetParam) (*dto.WorkoutSet, errcode.Error)
	DeleteWorkoutSet(c *gin.Context, setID int64) (*dto.WorkoutSetID, errcode.Error)
	UpdateWorkoutSetOrders(c *gin.Context, workoutID int64, params []*dto.WorkoutSetOrder) errcode.Error
	UploadWorkoutSetStartAudio(c *gin.Context, setID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	UploadWorkoutSetProgressAudio(c *gin.Context, setID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	DeleteWorkoutSetStartAudio(c *gin.Context, setID int64) errcode.Error
	DeleteWorkoutSetProgressAudio(c *gin.Context, setID int64) errcode.Error
}

type Action interface {
	CreateAction(c *gin.Context, courseID int64, param *dto.CreateActionParam) (*dto.Action, errcode.Error)
	UpdateAction(c *gin.Context, actionID int64, param *dto.UpdateActionParam) (*dto.Action, errcode.Error)
	SearchActions(c *gin.Context, courseID int64, param *dto.FindActionsParam) ([]*dto.Action, errcode.Error)
	DeleteAction(c *gin.Context, actionID int64) (*dto.ActionID, errcode.Error)
	UploadActionCover(c *gin.Context, actionID int64, coverNamed string, file multipart.File) (*dto.ActionCover, errcode.Error)
	UploadActionVideo(c *gin.Context, actionID int64, videoNamed string, file multipart.File) (*dto.ActionVideo, errcode.Error)
}

type Sale interface {
	GetSaleItems(c *gin.Context) ([]*saledto.SaleItem, errcode.Error)
}

type Review interface {
	CourseSubmit(c *gin.Context, courseID int64) errcode.Error
}