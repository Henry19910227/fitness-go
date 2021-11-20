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
	reviewRepo repository.Review
	uploader  handler.Uploader
	resHandler handler.Resource
	errHandler errcode.Handler
}

func NewReview(reviewRepo repository.Review, uploader  handler.Uploader, resHandler handler.Resource, errHandler errcode.Handler) Review {
	return &review{reviewRepo: reviewRepo, uploader: uploader, resHandler: resHandler, errHandler: errHandler}
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
	//創建評論
	reviewID, err := r.reviewRepo.CreateReview(&model.CreateReviewParam{
		CourseID: param.CourseID,
		UserID: param.UserID,
		Score: param.Score,
		Body: param.Body,
		ImageNames: reviewImageNames,
	})
	if err != nil {
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	//儲存評論照片
	for _, file := range param.Images {
		if err := r.uploader.UploadReviewImage(file.Data, file.FileNamed); err != nil {
			r.errHandler.Set(c, "uploader", err)
		}
	}
	//查詢並回傳創建資料
	item, err := r.reviewRepo.FindReviewByID(reviewID)
	if err != nil {
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	return parserReview(item), nil
}

func (r *review) GetReview(c *gin.Context, reviewID int64) (*dto.Review, errcode.Error) {
	item, err := r.reviewRepo.FindReviewByID(reviewID)
	if err != nil {
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	return parserReview(item), nil
}

func (r *review) GetReviews(c *gin.Context, courseID int64, uid int64, page int, size int) ([]*dto.Review, errcode.Error) {
	//查詢並回傳創建資料
	offset, limit := r.GetPagingIndex(page, size)
	items, err := r.reviewRepo.FindReviewsByCourseID(courseID, uid, &model.PagingParam{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	reviews := make([]*dto.Review, 0)
	for _, item := range items{
		reviews = append(reviews, parserReview(item))
	}
	return reviews, nil
}

func (r *review) DeleteReview(c *gin.Context, reviewID int64) errcode.Error {
	// 查詢當前 review 狀態
	review, err := r.reviewRepo.FindReviewByID(reviewID)
	if err != nil {
		return r.errHandler.Set(c, "review repo", err)
	}
	// 刪除 review
	if err := r.reviewRepo.DeleteReview(review.ID); err != nil {
		return r.errHandler.Set(c, "review repo", err)
	}
	// 刪除該 review 底下的圖片檔
	for _, image := range review.Images {
		if err := r.resHandler.DeleteReviewImage(image.Image); err != nil {
			r.errHandler.Set(c, "resource handler", err)
		}
	}
	return nil
}

func (r *review) GetReviewOwner(c *gin.Context, reviewID int64) (int64, errcode.Error) {
	review, err := r.reviewRepo.FindReviewByID(reviewID)
	if err != nil {
		return 0, r.errHandler.Set(c, "review repo", err)
	}
	return review.UserID, nil
}

func parserReview(item *model.Review) *dto.Review {
	review := dto.Review{
		ID: item.ID,
		User: &dto.UserSummary{
			ID:     item.User.ID,
			Nickname: item.User.Nickname,
			Avatar: item.User.Avatar,
		},
		CourseID: item.CourseID,
		Score: item.Score,
		Body: item.Body,
		CreateAt: item.CreateAt,
	}
	images := make([]*dto.ReviewImage, 0)
	for _, imageItem := range item.Images{
		image := dto.ReviewImage{
			ID: imageItem.ID,
			Image: imageItem.Image,
		}
		images = append(images, &image)
	}
	review.Images = images
	return &review
}

