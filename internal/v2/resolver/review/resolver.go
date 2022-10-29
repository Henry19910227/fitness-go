package review

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	groupModel "github.com/Henry19910227/fitness-go/internal/v2/model/group"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderByModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	reviewImageModel "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	reviewStatisticModel "github.com/Henry19910227/fitness-go/internal/v2/model/review_statistic"
	trainerStatisticModel "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	reviewService "github.com/Henry19910227/fitness-go/internal/v2/service/review"
	reviewImageService "github.com/Henry19910227/fitness-go/internal/v2/service/review_image"
	"github.com/Henry19910227/fitness-go/internal/v2/service/review_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_statistic"
	"gorm.io/gorm"
)

type resolver struct {
	reviewService      reviewService.Service
	reviewImageService reviewImageService.Service
	reviewStatisticService review_statistic.Service
	trainerStatisticService trainer_statistic.Service
	courseService course.Service
	uploadTool         uploader.Tool
}

func New(reviewService reviewService.Service, reviewImageService reviewImageService.Service,
	reviewStatisticService review_statistic.Service, trainerStatisticService trainer_statistic.Service,
	courseService course.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{reviewService: reviewService, reviewImageService: reviewImageService,
		reviewStatisticService: reviewStatisticService, trainerStatisticService: trainerStatisticService,
		courseService: courseService, uploadTool: uploadTool}
}

func (r *resolver) APIGetCMSReviews(input *model.APIGetCMSReviewsInput) (output model.APIGetCMSReviewsOutput) {
	// parser input
	param := model.ListInput{}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Course"},
		{Field: "User"},
		{Field: "Images"},
	}
	if err := util.Parser(input.Query, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// get list
	datas, page, err := r.reviewService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSReviewsData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIUpdateCMSReview(input *model.APIUpdateCMSReviewInput) (output model.APIUpdateCMSReviewOutput) {
	//parser input
	table := model.Table{}
	if err := util.Parser(input.Uri, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//執行更新
	if err := r.reviewService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteCMSReview(input *model.APIDeleteCMSReviewInput) (output model.APIDeleteCMSReviewOutput) {
	//查找評論
	findInput := model.FindInput{}
	if err := util.Parser(input.Uri, &findInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Images"},
	}
	findOutput, err := r.reviewService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser delete input
	deleteInput := model.DeleteInput{}
	if err := util.Parser(input.Uri, &deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//刪除評論
	if err := r.reviewService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//刪除評論照片
	for _, imageOutput := range findOutput.Images {
		_ = r.uploadTool.Delete(util.OnNilJustReturnString(imageOutput.Image, ""))
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetStoreCourseReviews(input *model.APIGetStoreCourseReviewsInput) (output model.APIGetStoreCourseReviewsOutput) {
	joins := make([]*joinModel.Join, 0)
	groups := make([]*groupModel.Group, 0)
	if util.OnNilJustReturnInt(input.Query.FilterType, 0) == 2 {
		joins = append(joins, &joinModel.Join{Query: "INNER JOIN review_images ON review_images.review_id = reviews.id"})
		groups = append(groups, &groupModel.Group{Name: "reviews.id"})
	}
	listInput := model.ListInput{}
	listInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	listInput.Joins = joins
	listInput.Groups = groups
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "User"},
		{Field: "Images"},
	}
	listInput.Orders = []*orderByModel.Order{
		{Value: fmt.Sprintf("reviews.user_id <> %v %s", input.UserID, orderByModel.ASC)},
	}
	listInput.Page = input.Query.Page
	listInput.Size = input.Query.Size
	reviewOutputs, page, err := r.reviewService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	data := model.APIGetStoreCourseReviewsData{}
	if err := util.Parser(reviewOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreCourseReview(input *model.APIGetStoreCourseReviewInput) (output model.APIGetStoreCourseReviewOutput) {
	findInput := model.FindInput{}
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "User"},
		{Field: "Images"},
	}
	findInput.ID = util.PointerInt64(input.Uri.ID)
	reviewOutput, err := r.reviewService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser Output
	data := model.APIGetStoreCourseReviewData{}
	if err := util.Parser(reviewOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APICreateStoreCourseReview(tx *gorm.DB, input *model.APICreateStoreCourseReviewInput) (output model.APICreateStoreCourseReviewOutput) {
	defer tx.Rollback()
	// 查詢課表資訊
	findCourseInput := courseModel.FindInput{}
	findCourseInput.ID = util.PointerInt64(input.Uri.CourseID)
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證是否有評分過此課表
	reviewListInput := model.ListInput{}
	reviewListInput.CourseID = util.PointerInt64(input.Uri.CourseID)
	reviewListInput.UserID = util.PointerInt64(input.UserID)
	reviewOutputs, _, err := r.reviewService.Tx(tx).List(&reviewListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(reviewOutputs) > 0 {
		output.Set(code.BadRequest, "已評分過此課表")
		return output
	}
	// 新增 review
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.CourseID = util.PointerInt64(input.Uri.CourseID)
	table.Score = util.PointerInt(input.Form.Score)
	table.Body = input.Form.Body
	reviewID, err := r.reviewService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 新增 review image
	reviewImageTables := make([]*reviewImageModel.Table, 0)
	for _, file := range input.Files {
		imageNamed, _ := r.uploadTool.Save(file.Data, file.Named)
		if len(imageNamed) > 0 {
			reviewImageTable := reviewImageModel.Table{}
			reviewImageTable.ReviewID = util.PointerInt64(reviewID)
			reviewImageTable.Image = util.PointerString(imageNamed)
			reviewImageTables = append(reviewImageTables, &reviewImageTable)
		}
	}
	if err := r.reviewImageService.Tx(tx).Create(reviewImageTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新評論統計
	reviewStatisticInput := reviewStatisticModel.StatisticInput{}
	reviewStatisticInput.CourseID = util.OnNilJustReturnInt64(courseOutput.ID, 0)
	if err := r.reviewStatisticService.Tx(tx).Statistic(&reviewStatisticInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新教練統計
	trainerStatisticInput := trainerStatisticModel.StatisticReviewScoreInput{}
	trainerStatisticInput.UserID = util.OnNilJustReturnInt64(courseOutput.UserID, 0)
	if err := r.trainerStatisticService.Tx(tx).StatisticReviewScore(&trainerStatisticInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteStoreCourseReview(tx *gorm.DB, input *model.APIDeleteStoreCourseReviewInput) (output model.APIDeleteStoreCourseReviewOutput) {
	defer tx.Rollback()
	// 查詢評論資訊
	findReviewInput := model.FindInput{}
	findReviewInput.ID = util.PointerInt64(input.Uri.ID)
	findReviewInput.Preloads = []*preloadModel.Preload{
		{Field: "Images"},
	}
	reviewOutput, err := r.reviewService.Tx(tx).Find(&findReviewInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt64(reviewOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非評論發布者，無法刪除該評論")
		return output
	}
	// 查詢課表資訊
	findCourseInput := courseModel.FindInput{}
	findCourseInput.ID = reviewOutput.CourseID
	courseOutput, err := r.courseService.Tx(tx).Find(&findCourseInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除評論
	deleteReviewInput := model.DeleteInput{}
	deleteReviewInput.ID = util.PointerInt64(input.Uri.ID)
	if err := r.reviewService.Tx(tx).Delete(&deleteReviewInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新評論統計
	reviewStatisticInput := reviewStatisticModel.StatisticInput{}
	reviewStatisticInput.CourseID = util.OnNilJustReturnInt64(courseOutput.ID, 0)
	if err := r.reviewStatisticService.Tx(tx).Statistic(&reviewStatisticInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新教練統計
	trainerStatisticInput := trainerStatisticModel.StatisticReviewScoreInput{}
	trainerStatisticInput.UserID = util.OnNilJustReturnInt64(courseOutput.UserID, 0)
	if err := r.trainerStatisticService.Tx(tx).StatisticReviewScore(&trainerStatisticInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// 刪除圖片檔
	for _, image := range reviewOutput.Images{
		_ = r.uploadTool.Delete(util.OnNilJustReturnString(image.Image, ""))
	}
	// Parser output
	output.Set(code.Success, "success")
	return output
}