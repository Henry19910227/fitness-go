package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	bodyImageService "github.com/Henry19910227/fitness-go/internal/v2/service/body_image"
)

type resolver struct {
	bodyImageService bodyImageService.Service
}

func New(bodyImageService bodyImageService.Service) Resolver {
	return &resolver{bodyImageService: bodyImageService}
}

func (r *resolver) APIGetBodyImages(input *model.APIGetBodyImagesInput) (output model.APIGetBodyImagesOutput) {
	// parser input
	bodyInput := model.ListInput{}
	bodyInput.UserID = util.PointerInt64(input.UserID)
	bodyInput.OrderField = "create_at"
	bodyInput.OrderType = order_by.DESC
	if err := util.Parser(input.Query, &bodyInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// List
	datas, page, err := r.bodyImageService.List(&bodyInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetBodyImagesData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}
