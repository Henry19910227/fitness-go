package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/dto/registerdto"
	"github.com/Henry19910227/fitness-go/internal/global"
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
	GetCMSUsers(c *gin.Context, param *dto.FinsCMSUsersParam, orderByParam *dto.OrderByParam, pagingParam *dto.PagingParam) ([]*dto.CMSUserSummary, *dto.Paging, errcode.Error)
	GetCMSUser(c *gin.Context, userID int64) (*dto.CMSUser, errcode.Error)
	UploadUserAvatarByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.UserAvatar, errcode.Error)
	UploadUserAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*dto.UserAvatar, errcode.Error)
}

type Trainer interface {
	CreateTrainer(c *gin.Context, uid int64, param *dto.CreateTrainerParam) (*dto.Trainer, errcode.Error)
	UpdateTrainer(c *gin.Context, uid int64, param *dto.UpdateTrainerParam) (*dto.Trainer, errcode.Error)
	GetTrainer(c *gin.Context, uid *int64, trainerID int64) (*dto.Trainer, errcode.Error)
	GetTrainerSummaries(c *gin.Context, param dto.GetTrainerSummariesParam, page, size int) ([]*dto.TrainerSummary, *dto.Paging, errcode.Error)
	UploadAlbumPhoto(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerAlbumPhotoResult, errcode.Error)
	DeleteAlbumPhoto(c *gin.Context, photoID int64) errcode.Error
	CreateCertificate(c *gin.Context, uid int64, name, imageNamed string, imageFile multipart.File) (*dto.Certificate, errcode.Error)
	UpdateCertificate(c *gin.Context, cerID int64, name *string, file *dto.File) (*dto.Certificate, errcode.Error)
	DeleteCertificate(c *gin.Context, cerID int64) errcode.Error
	GetTrainerAlbumPhotoCount(c *gin.Context, uid int64) (int, errcode.Error)
	GetCertificateCount(c *gin.Context, uid int64) (int, errcode.Error)
	GetCMSTrainers(c *gin.Context, param *dto.FinsCMSTrainersParam, orderByParam *dto.OrderByParam, pagingParam *dto.PagingParam) ([]*dto.CMSTrainerSummary, *dto.Paging, errcode.Error)
	GetCMSTrainer(c *gin.Context, userID int64) (*dto.CMSTrainer, errcode.Error)
}

type Course interface {
	CreateCourseByToken(c *gin.Context, token string, param *dto.CreateCourseParam) (*dto.Course, errcode.Error)
	CreateCourse(c *gin.Context, uid int64, param *dto.CreateCourseParam) (*dto.Course, errcode.Error)
	UpdateCourse(c *gin.Context, courseID int64, param *dto.UpdateCourseParam) (*dto.Course, errcode.Error)
	UpdateCourseSaleType(c *gin.Context, courseID int64, saleType int, saleID *int64) (*dto.Course, errcode.Error)
	DeleteCourse(c *gin.Context, courseID int64) (*dto.CourseID, errcode.Error)
	GetCourseSummariesByUID(c *gin.Context, uid int64, status []int, orderByParam *dto.OrderByParam, pagingParam *dto.PagingParam) ([]*dto.CourseSummary, *dto.Paging, errcode.Error)
	GetCourseDetailByCourseID(c *gin.Context, courseID int64) (*dto.Course, errcode.Error)
	GetCourseProductByCourseID(c *gin.Context, userID int64, courseID int64) (*dto.CourseProduct, errcode.Error)
	GetCourseOverviewByCourseID(c *gin.Context, courseID int64) (*dto.CourseProduct, errcode.Error)
	GetCourseProductSummaries(c *gin.Context, param *dto.GetCourseProductSummariesParam, page, size int) ([]*dto.CourseProductSummary, *dto.Paging, errcode.Error)
	GetProgressCourseAssetSummaries(c *gin.Context, userID int64, page int, size int) ([]*dto.CourseAssetSummary, *dto.Paging, errcode.Error)
	GetChargeCourseAssetSummaries(c *gin.Context, userID int64, page int, size int) ([]*dto.CourseAssetSummary, *dto.Paging, errcode.Error)
	GetCourseAsset(c *gin.Context, userID int64, courseID int64) (*dto.CourseAsset, errcode.Error)
	GetCourseAssetStructure(c *gin.Context, userID int64, courseID int64) (*dto.CourseAssetStructure, errcode.Error)
	GetCourseProductStructure(c *gin.Context, userID int64, courseID int64) (*dto.CourseProductStructure, errcode.Error)
	UploadCourseCoverByID(c *gin.Context, courseID int64, param *dto.UploadCourseCoverParam) (*dto.CourseCover, errcode.Error)
	CourseSubmit(c *gin.Context, courseID int64) errcode.Error
	GetCourseStatus(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error)
}

type Plan interface {
	CreatePlan(c *gin.Context, courseID int64, name string) (*dto.Plan, errcode.Error)
	UpdatePlan(c *gin.Context, planID int64, name string) (*dto.Plan, errcode.Error)
	DeletePlan(c *gin.Context, planID int64) (*dto.PlanID, errcode.Error)
	GetPlansByCourseID(c *gin.Context, courseID int64) ([]*dto.Plan, errcode.Error)
	GetPlanAssets(c *gin.Context, userID int64, courseID int64) ([]*dto.PlanAsset, errcode.Error)
	GetPlanStatus(c *gin.Context, planID int64) (global.CourseStatus, errcode.Error)
}

type Workout interface {
	CreateWorkout(c *gin.Context, planID int64, name string) (*dto.Workout, errcode.Error)
	GetWorkouts(c *gin.Context, planID int64) ([]*dto.Workout, errcode.Error)
	GetWorkoutAssets(c *gin.Context, userID int64, planID int64) ([]*dto.WorkoutAsset, errcode.Error)
	UpdateWorkout(c *gin.Context, workoutID int64, param *dto.UpdateWorkoutParam) (*dto.Workout, errcode.Error)
	DeleteWorkout(c *gin.Context, workoutID int64) (*dto.WorkoutID, errcode.Error)
	UploadWorkoutStartAudio(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	UploadWorkoutEndAudio(c *gin.Context, workoutID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	CreateWorkoutByTemplate(c *gin.Context, planID int64, name string, workoutTemplateID int64) (*dto.Workout, errcode.Error)
	DeleteWorkoutStartAudio(c *gin.Context, workoutID int64) errcode.Error
	DeleteWorkoutEndAudio(c *gin.Context, workoutID int64) errcode.Error
	GetWorkoutStatus(c *gin.Context, workoutID int64) (global.CourseStatus, errcode.Error)
}

type WorkoutSet interface {
	CreateRestSet(c *gin.Context, workoutID int64) (*dto.WorkoutSet, errcode.Error)
	CreateWorkoutSets(c *gin.Context, workoutID int64, actionIDs []int64) ([]*dto.WorkoutSet, errcode.Error)
	DuplicateWorkoutSets(c *gin.Context, setID int64, count int) ([]*dto.WorkoutSet, errcode.Error)
	GetWorkoutSets(c *gin.Context, workoutID int64, userID *int64) ([]*dto.WorkoutSet, errcode.Error)
	UpdateWorkoutSet(c *gin.Context, setID int64, param *dto.UpdateWorkoutSetParam) (*dto.WorkoutSet, errcode.Error)
	DeleteWorkoutSet(c *gin.Context, setID int64) (*dto.WorkoutSetID, errcode.Error)
	UpdateWorkoutSetOrders(c *gin.Context, workoutID int64, params []*dto.WorkoutSetOrder) errcode.Error
	UploadWorkoutSetStartAudio(c *gin.Context, setID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	UploadWorkoutSetProgressAudio(c *gin.Context, setID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error)
	DeleteWorkoutSetStartAudio(c *gin.Context, setID int64) errcode.Error
	DeleteWorkoutSetProgressAudio(c *gin.Context, setID int64) errcode.Error
}

type WorkoutLog interface {
	CreateWorkoutLog(c *gin.Context, userID int64, workoutID int64, param *dto.WorkoutLogParam) ([]*dto.WorkoutSetLogTag, errcode.Error)
	GetWorkoutLogSummaries(c *gin.Context, userID int64, startDate string, endDate string) ([]*dto.WorkoutLogSummary, errcode.Error)
	GetWorkoutLog(c *gin.Context, workoutLogID int64) (*dto.WorkoutLog, errcode.Error)
	DeleteWorkoutLog(c *gin.Context, userID int64, workoutLogID int64) errcode.Error
}

type WorkoutSetLog interface {
	GetWorkoutSetLogSummaries(c *gin.Context, userID int64, actionID int64, startDate string, endDate string) ([]*dto.WorkoutSetLogSummary, errcode.Error)
}

type Action interface {
	CreateAction(c *gin.Context, courseID int64, param *dto.CreateActionParam) (*dto.Action, errcode.Error)
	UpdateAction(c *gin.Context, actionID int64, param *dto.UpdateActionParam) (*dto.Action, errcode.Error)
	SearchActions(c *gin.Context, userID int64, param *dto.FindActionsParam) ([]*dto.Action, errcode.Error)
	FindActionBestPR(c *gin.Context, userID int64, actionID int64) (*dto.ActionBestPR, errcode.Error)
	DeleteAction(c *gin.Context, actionID int64) (*dto.ActionID, errcode.Error)
	DeleteActionVideo(c *gin.Context, actionID int64) errcode.Error
}

type Sale interface {
	GetSaleItems(c *gin.Context) ([]*dto.SaleItem, errcode.Error)
	GetSubscribePlans(c *gin.Context) ([]*dto.SubscribePlan, errcode.Error)
}

type Store interface {
	GetHomePage(c *gin.Context) (*dto.StoreHomePage, errcode.Error)
}

type Review interface {
	CreateReview(c *gin.Context, param *dto.CreateReviewParam) (*dto.Review, errcode.Error)
	GetReview(c *gin.Context, reviewID int64) (*dto.Review, errcode.Error)
	GetReviews(c *gin.Context, uid int64, param *dto.GetReviewsParam, page int, size int) ([]*dto.Review, *dto.Paging, errcode.Error)
	DeleteReview(c *gin.Context, reviewID int64) errcode.Error
	GetReviewOwner(c *gin.Context, reviewID int64) (int64, errcode.Error)
}

type Payment interface {
	CreateCourseOrder(c *gin.Context, uid int64, courseID int64) (*dto.CourseOrder, errcode.Error)
	CreateSubscribeOrder(c *gin.Context, uid int64, subscribePlanID int64) (*dto.SubscribeOrder, errcode.Error)
	CheckUserSubscribeStatus(c *gin.Context, uid int64) (global.SubscribeStatus, errcode.Error)
	VerifyFreeCourseOrder(c *gin.Context, uid int64, orderID string) errcode.Error
	VerifyAppleReceipt(c *gin.Context, uid int64, orderID string, receiptData string) errcode.Error
	VerifyGoogleReceipt(c *gin.Context, uid int64, orderID string, productID string, receiptData string) errcode.Error
	HandleAppStoreNotification(c *gin.Context, base64PayloadString string) errcode.Error
	HandleGooglePlayNotification(c *gin.Context, base64PayloadString string) errcode.Error
	GetAppStoreAPISubscriptions(c *gin.Context, originalTransactionID string) (*dto.IAPSubscribeAPIResponse, errcode.Error)
	GetAppStoreAPIHistory(c *gin.Context, originalTransactionID string) (*dto.IAPHistoryAPIResponse, errcode.Error)
	GetGooglePlayAPIProduct(c *gin.Context, productID string, purchaseToken string) (*dto.IABProductAPIResponse, errcode.Error)
	GetGooglePlayAPISubscription(c *gin.Context, productID string, purchaseToken string) (*dto.IABSubscriptionAPIResponse, errcode.Error)
	GetGooglePlayApiAccessToken(c *gin.Context) (string, errcode.Error)
	GetAppleStoreApiAccessToken(c *gin.Context) (string, errcode.Error)
}

type Favorite interface {
	CreateFavoriteCourse(c *gin.Context, userID int64, courseID int64) errcode.Error
	CreateFavoriteTrainer(c *gin.Context, userID int64, trainerID int64) errcode.Error
	CreateFavoriteAction(c *gin.Context, userID int64, actionID int64) errcode.Error
	DeleteFavoriteCourse(c *gin.Context, userID int64, courseID int64) errcode.Error
	DeleteFavoriteTrainer(c *gin.Context, userID int64, trainerID int64) errcode.Error
	DeleteFavoriteAction(c *gin.Context, userID int64, actionID int64) errcode.Error
}
