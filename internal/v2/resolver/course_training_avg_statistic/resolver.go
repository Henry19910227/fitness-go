package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic/api_get_cms_statistic_monthly_course_training_avg"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_training_avg_statistic"
)

type resolver struct {
	courseService    course.Service
	statisticService course_training_avg_statistic.Service
}

func New(courseService course.Service, statisticService course_training_avg_statistic.Service) Resolver {
	return &resolver{courseService: courseService, statisticService: statisticService}
}

func (r *resolver) APIGetCMSCourseTrainingAvgStatistic(input *api_get_cms_statistic_monthly_course_training_avg.Input) (output api_get_cms_statistic_monthly_course_training_avg.Output) {
	listInput := courseModel.ListInput{}
	listInput.ID = input.Query.CourseID
	listInput.CourseStatus = input.Query.CourseStatus
	listInput.Name = input.Query.Name
	listInput.SaleType = input.Query.SaleType
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if input.Query.SaleType == nil {
		listInput.Wheres = []*whereModel.Where{
			{Query: "courses.sale_type IN (?)", Args: []interface{}{[]int{courseModel.SaleTypeFree, courseModel.SaleTypeSubscribe, courseModel.SaleTypeCharge}}},
		}
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
	data := api_get_cms_statistic_monthly_course_training_avg.Data{}
	if err := util.Parser(courseOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) Statistic() {
	_ = r.statisticService.Statistic()
}
