package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
)

type review struct {
	Base
	reviewRepo           repository.Review
	reviewImageRepo      repository.ReviewImage
	reviewStatRepo       repository.ReviewStatistic
	courseRepo           repository.Course
	trainerStatisticRepo repository.TrainerStatistic
	transactionRepo      repository.Transaction
	uploader             handler.Uploader
	resHandler           handler.Resource
	errHandler           errcode.Handler
}

func NewReview(reviewRepo repository.Review, reviewImageRepo repository.ReviewImage,
	reviewStatRepo repository.ReviewStatistic, courseRepo repository.Course,
	trainerStatisticRepo repository.TrainerStatistic, transactionRepo repository.Transaction,
	uploader handler.Uploader, resHandler handler.Resource, errHandler errcode.Handler) Review {
	return &review{reviewRepo: reviewRepo, reviewImageRepo: reviewImageRepo,
		reviewStatRepo: reviewStatRepo, courseRepo: courseRepo,
		trainerStatisticRepo: trainerStatisticRepo, transactionRepo: transactionRepo, uploader: uploader, resHandler: resHandler, errHandler: errHandler}
}

func (r *review) CreateReview(c *gin.Context, param *dto.CreateReviewParam) (*dto.Review, errcode.Error) {
	//生成教練相簿照片名稱
	var reviewImageNames []string
	for _, file := range param.Images {
		reviewImageName, err := r.uploader.GenerateNewImageName(file.FileNamed)
		if err != nil {
			return nil, r.errHandler.Set(c, "uploader", err)
		}
		file.FileNamed = reviewImageName
		reviewImageNames = append(reviewImageNames, reviewImageName)
	}
	tx := r.transactionRepo.CreateTransaction()
	//創建評論
	reviewID, err := r.reviewRepo.CreateReview(tx, &model.CreateReviewParam{
		CourseID:   param.CourseID,
		UserID:     param.UserID,
		Score:      param.Score,
		Body:       param.Body,
		ImageNames: reviewImageNames,
	})
	if err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	//創建評論照片
	if err := r.reviewImageRepo.CreateReviewImages(tx, reviewID, reviewImageNames); err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "review image repo", err)
	}
	//計算評論統計
	reviewStat, err := r.reviewStatRepo.CalculateReviewStatistic(tx, param.CourseID)
	if err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "review statistic repo", err)
	}
	//儲存評論統計
	if err := r.reviewStatRepo.SaveReviewStatistic(tx, param.CourseID, &model.SaveReviewStatisticParam{
		ScoreTotal: reviewStat.ScoreTotal,
		Amount:     reviewStat.Amount,
		FiveTotal:  reviewStat.FiveTotal,
		FourTotal:  reviewStat.FourTotal,
		ThreeTotal: reviewStat.ThreeTotal,
		TwoTotal:   reviewStat.TwoTotal,
		OneTotal:   reviewStat.OneTotal,
	}); err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "review statistic repo", err)
	}
	//查詢該課表教練id
	course := struct {
		UserID int64 `gorm:"column:user_id"`
	}{}
	if err := r.courseRepo.FindCourseByID(tx, param.CourseID, &course); err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "course repo", err)
	}
	//計算評論平均分數
	reviewScore, err := r.trainerStatisticRepo.CalculateTrainerReviewScore(tx, course.UserID)
	if err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "trainer statistic repo", err)
	}
	//儲存教練評論平均分數
	if err := r.trainerStatisticRepo.SaveTrainerStatistic(tx, course.UserID, &model.SaveTrainerStatisticParam{
		ReviewScore: &reviewScore,
	}); err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "trainer statistic repo", err)
	}
	//查詢並回傳創建資料
	item, err := r.reviewRepo.FindReviewByID(tx, reviewID)
	if err != nil {
		tx.Rollback()
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	r.transactionRepo.FinishTransaction(tx)
	//儲存評論照片
	for _, file := range param.Images {
		if err := r.uploader.UploadReviewImage(file.Data, file.FileNamed); err != nil {
			r.errHandler.Set(c, "uploader", err)
		}
	}
	return parserReview(item), nil
}

func (r *review) GetReview(c *gin.Context, reviewID int64) (*dto.Review, errcode.Error) {
	item, err := r.reviewRepo.FindReviewByID(nil, reviewID)
	if err != nil {
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	return parserReview(item), nil
}

func (r *review) GetReviews(c *gin.Context, uid int64, param *dto.GetReviewsParam, page int, size int) ([]*dto.Review, *dto.Paging, errcode.Error) {
	//查詢並回傳創建資料
	offset, limit := r.GetPagingIndex(page, size)
	findReviewsParam := model.FindReviewsParam{
		CourseID:   param.CourseID,
		FilterType: param.FilterType,
	}
	items, err := r.reviewRepo.FindReviews(uid, &findReviewsParam, &model.PagingParam{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, r.errHandler.Set(c, "review repo", err)
	}
	reviews := make([]*dto.Review, 0)
	for _, item := range items {
		reviews = append(reviews, parserReview(item))
	}
	totalCount, err := r.reviewRepo.FindReviewsCount(&findReviewsParam)
	if err != nil {
		return nil, nil, r.errHandler.Set(c, "review repo", err)
	}
	paging := dto.Paging{
		TotalCount: totalCount,
		TotalPage:  r.GetTotalPage(totalCount, size),
		Page:       page,
		Size:       size,
	}
	return reviews, &paging, nil
}

func (r *review) DeleteReview(c *gin.Context, reviewID int64) errcode.Error {
	tx := r.transactionRepo.CreateTransaction()
	// 查詢當前 review 狀態
	review, err := r.reviewRepo.FindReviewByID(tx, reviewID)
	if err != nil {
		tx.Rollback()
		return r.errHandler.Set(c, "review repo", err)
	}
	// 刪除 review
	if err := r.reviewRepo.DeleteReview(tx, review.ID); err != nil {
		tx.Rollback()
		return r.errHandler.Set(c, "review repo", err)
	}
	//計算評論統計
	reviewStat, err := r.reviewStatRepo.CalculateReviewStatistic(tx, review.CourseID)
	if err != nil {
		tx.Rollback()
		return r.errHandler.Set(c, "review statistic repo", err)
	}
	//儲存評論統計
	if err := r.reviewStatRepo.SaveReviewStatistic(tx, review.CourseID, &model.SaveReviewStatisticParam{
		ScoreTotal: reviewStat.ScoreTotal,
		Amount:     reviewStat.Amount,
		FiveTotal:  reviewStat.FiveTotal,
		FourTotal:  reviewStat.FourTotal,
		ThreeTotal: reviewStat.ThreeTotal,
		TwoTotal:   reviewStat.TwoTotal,
		OneTotal:   reviewStat.OneTotal,
	}); err != nil {
		tx.Rollback()
		return r.errHandler.Set(c, "review statistic repo", err)
	}
	//查詢該課表教練id
	course := struct {
		UserID int64 `gorm:"column:user_id"`
	}{}
	if err := r.courseRepo.FindCourseByID(tx, review.CourseID, &course); err != nil {
		tx.Rollback()
		return r.errHandler.Set(c, "course repo", err)
	}
	//計算教練評論平均分數
	reviewScore, err := r.trainerStatisticRepo.CalculateTrainerReviewScore(tx, course.UserID)
	if err != nil {
		tx.Rollback()
		return r.errHandler.Set(c, "trainer statistic repo", err)
	}
	//儲存教練評論平均分數
	if err := r.trainerStatisticRepo.SaveTrainerStatistic(tx, course.UserID, &model.SaveTrainerStatisticParam{
		ReviewScore: &reviewScore,
	}); err != nil {
		tx.Rollback()
		return r.errHandler.Set(c, "trainer statistic repo", err)
	}
	r.transactionRepo.FinishTransaction(tx)
	// 刪除該 review 底下的圖片檔
	for _, image := range review.Images {
		if err := r.resHandler.DeleteReviewImage(image.Image); err != nil {
			r.errHandler.Set(c, "resource handler", err)
		}
	}
	return nil
}

func (r *review) GetReviewOwner(c *gin.Context, reviewID int64) (int64, errcode.Error) {
	review, err := r.reviewRepo.FindReviewByID(nil, reviewID)
	if err != nil {
		return 0, r.errHandler.Set(c, "review repo", err)
	}
	return review.UserID, nil
}

func parserReview(item *model.Review) *dto.Review {
	review := dto.Review{
		ID: item.ID,
		User: &dto.UserSummary{
			ID:       item.User.ID,
			Nickname: item.User.Nickname,
			Avatar:   item.User.Avatar,
		},
		CourseID: item.CourseID,
		Score:    item.Score,
		Body:     item.Body,
		CreateAt: item.CreateAt,
	}
	images := make([]*dto.ReviewImage, 0)
	for _, imageItem := range item.Images {
		image := dto.ReviewImage{
			ID:    imageItem.ID,
			Image: imageItem.Image,
		}
		images = append(images, &image)
	}
	review.Images = images
	return &review
}
