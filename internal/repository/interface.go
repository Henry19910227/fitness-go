package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
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
	FindTrainers(entity interface{}, status *global.TrainerStatus, orderBy *model.OrderBy, paging *model.PagingParam) error
	UpdateTrainerByUID(uid int64, param *model.UpdateTrainerParam) error
}

type Course interface {
	CreateCourse(uid int64, param *model.CreateCourseParam) (int64, error)
	CreateSingleWorkoutCourse(uid int64, param *model.CreateCourseParam) (int64, error)
	UpdateCourseByID(courseID int64, param *model.UpdateCourseParam) error
	FindCourseSummaries(param *model.FindCourseSummariesParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.CourseSummaryEntity, error)
	FindCourseProductSummaries(param model.FindCourseProductSummariesParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.CourseProductSummary, error)
	FindCourseProduct(courseID int64) (*model.CourseProduct, error)
	FindCourseDetailByCourseID(courseID int64) (*model.CourseDetailEntity, error)
	FindCourseAmountByUserID(uid int64) (int, error)
	FindCourseByID(courseID int64, entity interface{}) error
	FindCourseByPlanID(planID int64, entity interface{}) error
	FindCourseByWorkoutID(workoutID int64, entity interface{}) error
	FindCourseByWorkoutSetID(setID int64, entity interface{}) error
	FindCourseByActionID(actionID int64, entity interface{}) error
	DeleteCourseByID(courseID int64) error
}

type Plan interface {
	CreatePlan(courseID int64, name string) (int64, error)
	FindPlanByID(planID int64, entity interface{}) error
	FindPlansByCourseID(courseID int64) ([]*model.Plan, error)
	UpdatePlanByID(planID int64, name string) error
	DeletePlanByID(planID int64) error
	FindPlanOwnerByID(planID int64) (int64, error)
}

type Workout interface {
	CreateWorkout(planID int64, name string) (int64, error)
	FindWorkoutsByPlanID(planID int64) ([]*model.Workout, error)
	FindWorkoutByID(workoutID int64, entity interface{}) error
	FindStartAudioCountByAudioName(audioName string) (int, error)
	FindEndAudioCountByAudioName(audioName string) (int, error)
	UpdateWorkoutByID(workoutID int64, param *model.UpdateWorkoutParam) error
	DeleteWorkoutByID(workoutID int64) error
	FindWorkoutOwnerByID(workoutID int64) (int64, error)
}

type WorkoutSet interface {
	CreateWorkoutSetsByWorkoutID(workoutID int64, actionIDs []int64) ([]int64, error)
	CreateWorkoutSetsByWorkoutIDAndSets(workoutID int64, sets []*entity.WorkoutSet) ([]int64, error)
	CreateRestSetByWorkoutID(workoutID int64) (int64, error)
	FindWorkoutSetByID(setID int64) (*model.WorkoutSet, error)
	FindWorkoutSetsByIDs(setIDs []int64) ([]*model.WorkoutSet, error)
	FindWorkoutSetsByWorkoutID(workoutID int64) ([]*model.WorkoutSet, error)
	FindWorkoutSetsByCourseID(courseID int64) ([]*model.WorkoutSet, error)
	FindStartAudioCountByAudioName(audioName string) (int, error)
	FindProgressAudioCountByAudioName(audioName string) (int, error)
	UpdateWorkoutSetByID(setID int64, param *model.UpdateWorkoutSetParam) error
	DeleteWorkoutSetByID(setID int64) error
	UpdateWorkoutSetOrdersByWorkoutID(workoutID int64, params []*model.WorkoutSetOrder) error
}

type Action interface {
	CreateAction(courseID int64, param *model.CreateActionParam) (int64, error)
	FindActionByID(actionID int64, entity interface{}) error
	FindActionsByParam(courseID int64, param *model.FindActionsParam, entity interface{}) error
	UpdateActionByID(actionID int64, param *model.UpdateActionParam) error
	DeleteActionByID(actionID int64) error
}

type Sale interface {
	FinsSaleItems() ([]*model.SaleItemEntity, error)
}

type TrainerAlbum interface {
	CreateAlbumPhoto(uid int64, imageNamed string) error
	FindAlbumPhotoByUID(uid int64) ([]*model.TrainerAlbumPhotoEntity, error)
	FindAlbumPhotosByUID(uid int64, entity interface{}) error
	FindAlbumPhotoByID(photoID int64, entity interface{}) error
	FindAlbumPhotosByIDs(photoIDs []int64, entity interface{}) error
	DeleteAlbumPhotoByID(photoID int64) error
}

type Certificate interface {
	CreateCertificate(uid int64, name string, imageNamed string) (int64, error)
	FindCertificatesByUID(uid int64, entity interface{}) error
	UpdateCertificate(cerID int64, name *string, imageNamed *string) error
	DeleteCertificateByID(cerID int64) error
	FindCertificate(cerID int64, entity interface{}) error
	FindCertificatesByIDs(cerIDs []int64, entity interface{}) error
}

type Review interface {
	CreateReview(param *model.CreateReviewParam) (int64, error)
	DeleteReview(reviewID int64) error
	FindReviewByID(reviewID int64) (*model.Review, error)
	FindReviewsByCourseID(courseID int64, uid int64, paging *model.PagingParam) ([]*model.Review, error)
	FindReviewImages(courseID int64, userID int64) ([]*model.ReviewImageItem, error)
}