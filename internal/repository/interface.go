package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"gorm.io/gorm"
)

type Admin interface {
	GetAdminID(email string, password string) (int64, error)
	GetAdmin(uid int64, entity interface{}) error
}

type User interface {
	CreateUser(accountType int, account string, nickname string, password string) (int64, error)
	UpdateUserByUID(tx *gorm.DB, uid int64, param *model.UpdateUserParam) error
	FindUserByUID(uid int64, entity interface{}) error
	FindUserByAccountAndPassword(account string, password string, entity interface{}) error
	FindUserIDByNickname(nickname string) (int64, error)
	FindUserIDByEmail(email string) (int64, error)
}

type Trainer interface {
	CreateTrainer(uid int64, param *model.CreateTrainerParam) error
	FindTrainerByUID(uid int64, entity interface{}) error
	FindTrainers(entity interface{}, status *global.TrainerStatus, orderBy *model.OrderBy, paging *model.PagingParam) error
	FindTrainersCount(status *global.TrainerStatus) (int, error)
	UpdateTrainerByUID(uid int64, param *model.UpdateTrainerParam) error
}

type Course interface {
	CreateCourse(uid int64, param *model.CreateCourseParam) (int64, error)
	CreateSingleWorkoutCourse(uid int64, param *model.CreateCourseParam) (int64, error)
	UpdateCourseByID(courseID int64, param *model.UpdateCourseParam) error
	FindCourseSummaries(param *model.FindCourseSummariesParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.CourseSummary, error)
	FindCourseProductSummaries(param model.FindCourseProductSummariesParam, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.CourseProductSummary, error)
	FindCourseProductCount(param model.FindCourseProductCountParam) (int, error)
	FindCourseProduct(courseID int64) (*model.CourseProduct, error)
	FindProgressCourseAssetSummaries(userID int64, paging *model.PagingParam) ([]*model.CourseAssetSummary, error)
	FindChargeCourseAssetSummaries(userID int64, paging *model.PagingParam) ([]*model.CourseAssetSummary, error)
	FindProgressCourseAssetCount(userID int64) (int, error)
	FindChargeCourseAssetCount(userID int64) (int, error)
	FindCourseAsset(courseID int64, userID int64) (*model.CourseAsset, error)
	FindCourseByCourseID(courseID int64) (*model.Course, error)
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
	FindPlanAssets(userID int64, courseID int64) ([]*model.PlanAsset, error)
	UpdatePlanByID(planID int64, name string) error
	DeletePlanByID(planID int64) error
	FindPlanOwnerByID(planID int64) (int64, error)
}

type Workout interface {
	CreateWorkout(planID int64, name string) (int64, error)
	FindWorkoutsByPlanID(planID int64) ([]*model.Workout, error)
	FindWorkoutAssets(userID int64, planID int64) ([]*model.WorkoutAsset, error)
	FindWorkoutByID(workoutID int64, obj interface{}) error
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
	FindWorkoutSetIDsByWorkoutID(workoutID int64) ([]int64, error)
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

type ActionPR interface {
	FindActionPRs(userID int64, actionIDs []int64) ([]*model.ActionPR, error)
	SaveActionPRs(tx *gorm.DB, userID int64, params []*model.CreateActionPRParam) error
}

type Sale interface {
	FindSaleItems(saleType *int) ([]*model.SaleItem, error)
	FindSaleItemByID(saleID int64) (*model.SaleItem, error)
}

type SubscribePlan interface {
	FindSubscribePlans() ([]*model.SubscribePlan, error)
	FinsSubscribePlanByID(subscribePlanID int64) (*model.SubscribePlan, error)
	FinsSubscribePlanByProductID(productID string) (*model.SubscribePlan, error)
	FindSubscribePlansByPeriod(period global.PeriodType) ([]*model.SubscribePlan, error)
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
	FindReviews(uid int64, param *model.FindReviewsParam, paging *model.PagingParam) ([]*model.Review, error)
	FindReviewsCount(param *model.FindReviewsParam) (int, error)
}

type Order interface {
	CreateCourseOrder(param *model.CreateOrderParam) (string, error)
	CreateSubscribeOrder(param *model.CreateSubscribeOrderParam) (string, error)
	UpdateOrderStatus(tx *gorm.DB, orderID string, orderStatus global.OrderStatus) error
	UpdateOrderSubscribePlan(tx *gorm.DB, orderID string, subscribePlanID int64) error
	FindOrder(orderID string) (*model.Order, error)
	FindOrderByOriginalTransactionID(originalTransactionID string) (*model.Order, error)
	FindOrderByCourseID(userID int64, courseID int64) (*model.Order, error)
	FindOrdersByUserID(userID int64, paymentOrderType global.PaymentOrderType, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.Order, error)
}

type Receipt interface {
	SaveReceipt(tx *gorm.DB, param *model.CreateReceiptParam) (int64, error)
	FindReceiptsByOrderID(orderID string, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.Receipt, error)
}

type UserCourseAsset interface {
	CreateUserCourseAsset(tx *gorm.DB, param *model.CreateUserCourseAssetParam) (int64, error)
	FindUserCourseAsset(param *model.FindUserCourseAssetParam) (*model.UserCourseAsset, error)
}

type PurchaseLog interface {
	CreatePurchaseLog(tx *gorm.DB, param *model.CreatePurchaseLogParam) (int64, error)
}

type SubscribeLog interface {
	SaveSubscribeLog(tx *gorm.DB, param *model.CreateSubscribeLogParam) (int64, error)
}

type UserSubscribeInfo interface {
	SaveSubscribeInfo(tx *gorm.DB, param *model.SaveUserSubscribeInfoParam) (int64, error)
	FindSubscribeInfo(uid int64) (*model.UserSubscribeInfo, error)
	FindSubscribeInfoByOriginalTransactionID(originalTransactionID string) (*model.UserSubscribeInfo, error)
}

type Transaction interface {
	CreateTransaction() *gorm.DB
	FinishTransaction(tx *gorm.DB)
}

type WorkoutLog interface {
	FindWorkoutLog(workoutLogID int64) (*model.WorkoutLog, error)
	FindWorkoutLogsByDate(userID int64, startDate string, endDate string) ([]*model.WorkoutLog, error)
	FindWorkoutLogsByPlanID(planID int64) ([]*model.WorkoutLog, error)
	CalculateUserCourseStatistic(userID int64, workoutID int64) (*model.WorkoutLogCourseStatistic, error)
	CalculateUserPlanStatistic(userID int64, workoutID int64) (*model.WorkoutLogPlanStatistic, error)
	CreateWorkoutLog(tx *gorm.DB, param *model.CreateWorkoutLogParam) (int64, error)
}

type WorkoutSetLog interface {
	FindWorkoutSetLogsByWorkoutLogID(workoutLogID int64) ([]*model.WorkoutSetLog, error)
	FindWorkoutSetLogsByWorkoutSetIDs(userID int64, workoutSetIDs []int64) ([]*model.WorkoutSetLog, error)
	CreateWorkoutSetLogs(tx *gorm.DB, params []*model.WorkoutSetLogParam) error
	CalculateBestWorkoutSetLog(userID int64, actionIDs []int64) ([]*model.BestActionSetLog, error)
}

type UserCourseStatistic interface {
	FindUserCourseStatistic(userID int64, courseID int64) (*model.UserCourseStatistic, error)
	FindUserCourseStatisticByWorkoutID(workoutID int64, userID int64) (*model.UserCourseStatistic, error)
	SaveUserCourseStatistic(tx *gorm.DB, param *model.SaveUserCourseStatisticParam) (int64, error)
}

type UserPlanStatistic interface {
	FindUserPlanStatistics(userID int64, planID int64) ([]*model.UserPlanStatistic, error)
	SaveUserPlanStatistic(tx *gorm.DB, param *model.SaveUserPlanStatisticParam) (int64, error)
}

type Favorite interface {
	CreateFavoriteCourse(userID int64, courseID int64) error
	CreateFavoriteTrainer(userID int64, trainerID int64) error
	CreateFavoriteAction(userID int64, actionID int64) error
	FindFavoriteCourse(userID int64, courseID int64) (*model.FavoriteCourse, error)
	FindFavoriteTrainer(userID int64, trainerID int64) (*model.FavoriteTrainer, error)
	DeleteFavoriteCourse(userID int64, courseID int64) error
	DeleteFavoriteTrainer(userID int64, trainerID int64) error
}
