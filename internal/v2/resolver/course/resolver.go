package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	courseService "github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/gin-gonic/gin"
)

type resolver struct {
	courseService courseService.Service
	uploadTool    uploader.Tool
}

func New(courseService courseService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{courseService: courseService, uploadTool: uploadTool}
}

func (r *resolver) APIGetCMSCourses(ctx *gin.Context, input *model.APIGetCMSCoursesInput) interface{} {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
	}
	// 調用 repo
	result, page, err := r.courseService.List(&param)
	if err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSCoursesData{}
	if err := util.Parser(result, &data); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSCoursesOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging = page
	return output
}

func (r *resolver) APIGetCMSCourse(ctx *gin.Context, input *model.APIGetCMSCourseInput) interface{} {
	param := model.FindInput{}
	if err := util.Parser(input, &param); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
	}
	// 調用 repo
	result, err := r.courseService.Find(&param)
	if err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSCourseData{}
	if err := util.Parser(result, &data); err != nil {
		logger.Shared().Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSCourseOutput{}
	output.Data = &data
	output.Code = code.Success
	output.Msg = "success!"
	return output
}

func (r *resolver) APIUpdateCMSCoursesStatus(input *model.APIUpdateCMSCoursesStatusInput) (output base.Output) {
	tables := make([]*model.Table, 0)
	for _, courseID := range input.IDs {
		table := model.Table{}
		table.ID = util.PointerInt64(courseID)
		table.CourseStatus = &input.CourseStatus
		tables = append(tables, &table)
	}
	if err := r.courseService.Updates(tables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIUpdateCMSCourseCover(input *model.APIUpdateCMSCourseCoverInput) (output model.APIUpdateCMSCourseCoverOutput) {
	fileNamed, err := r.uploadTool.Save(input.File, input.CoverNamed)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table := model.Table{}
	table.ID = util.PointerInt64(input.ID)
	table.Cover = util.PointerString(fileNamed)
	if err := r.courseService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = util.PointerString(fileNamed)
	return output
}
