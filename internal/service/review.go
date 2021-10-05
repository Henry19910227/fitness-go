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
	err := r.reviewRepo.CreateReview(&model.CreateReviewParam{
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
	var review dto.Review
	if err := r.reviewRepo.FindReviewByCourseIDAndUserID(param.CourseID, param.UserID, &review); err != nil {
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	if err := r.reviewRepo.FindReviewImagesByReviewID(param.CourseID, param.UserID, &review.Images); err != nil {
		return nil, r.errHandler.Set(c, "review repo", err)
	}
	return &review, nil
}
