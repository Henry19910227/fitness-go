package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
)

type resolver struct {
	courseService course.Service
}

func New(courseService course.Service) Resolver {
	return &resolver{courseService: courseService}
}

func (r *resolver) APIGetCMSCourseTrainingAvgStatistic(input *model.APIGetCMSCourseTrainingAvgStatisticInput) (output model.APIGetCMSCourseTrainingAvgStatisticOutput) {
	listInput := courseModel.ListInput{}
	listInput.ID = input.Query.CourseID
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if input.Query.SaleType == nil {
		listInput.SaleTypes = []int{courseModel.SaleTypeFree, courseModel.SaleTypeSubscribe, courseModel.SaleTypeCharge}
	}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "SaleItem.ProductLabel"},
		{Field: "CourseTrainingAvgStatistic"},
	}
	courseOutputs, page, err := r.courseService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSCourseTrainingAvgStatisticData{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}
