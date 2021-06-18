package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
)

type Admin interface {
	GetAdminID(email string, password string) (int64, error)
	GetAdmin(uid int64, entity interface{}) error
}

type User interface {
	CreateUser(accountType int, account string, nickname string, password string) (int64, error)
	UpdateUserByUID(uid int64, param *model.UpdateUserParam) error
	FindUserByUID(uid int64, entity interface{}) error
	FindUserByAccountAndPassword(account string, password string, entity interface{}) error
	FindUserIDByNickname(nickname string) (int64, error)
	FindUserIDByEmail(email string) (int64, error)
}

type Trainer interface {
	CreateTrainer(uid int64, param *model.CreateTrainerParam) error
	FindTrainerByUID(uid int64, entity interface{}) error
	UpdateTrainerByUID(uid int64, param *model.UpdateTrainerParam) error
}

type Course interface {
	CreateCourse(uid int64, param *model.CreateCourseParam) (int64, error)
	UpdateCourseByID(courseID int64, param *model.UpdateCourseParam) error
	FindCourses(uid int64, entity interface{}) error
	FindCourseByID(courseID int64, entity interface{}) error
	CheckCourseExistByIDAndUID(courseID int64, uid int64) (bool, error)
}

type Plan interface {
	CreatePlan(courseID int64, name string) (int64, error)
	FindPlanByID(planID int64, entity interface{}) error
	FindPlansByCourseID(courseID int64) ([]*model.Plan, error)
	UpdatePlanByID(planID int64, name string) error
	DeletePlanByID(planID int64) error
	CheckPlanExistByUID(uid int64, planID int64) (bool, error)
}

type Workout interface {
	CreateWorkout(planID int64, name string) (int64, error)
	FindWorkoutsByPlanID(planID int64) ([]*model.Workout, error)
	CheckWorkoutExistByUID(uid int64, workoutID int64) (bool, error)
}