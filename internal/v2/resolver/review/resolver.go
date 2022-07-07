package review

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	reviewService "github.com/Henry19910227/fitness-go/internal/v2/service/review"
)

type resolver struct {
	orderService reviewService.Service
}

func New(orderService reviewService.Service) Resolver {
	return &resolver{orderService: orderService}
}

func (r *resolver) APIGetCMSReviews(input *model.APIGetCMSReviewsInput) (output model.APIGetCMSReviewsOutput) {
	// parser input
	param := model.ListInput{}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Course"},
		{Field: "User"},
		{Field: "Images"},
	}
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// get list
	datas, page, err := r.orderService.List(&param)
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
