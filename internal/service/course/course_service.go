package course

import (
	"github.com/Henry19910227/fitness-go/code"
	"github.com/Henry19910227/fitness-go/internal/model/base"
	model "github.com/Henry19910227/fitness-go/internal/model/course"
	"github.com/Henry19910227/fitness-go/internal/model/paging"
	preloadModel "github.com/Henry19910227/fitness-go/internal/model/preload"
	"github.com/Henry19910227/fitness-go/internal/repository/course"
	"github.com/Henry19910227/fitness-go/internal/util"
)

type service struct {
	repository course.Repository
}

func New(repository course.Repository) Service {
	return &service{repository: repository}
}

func (s *service) List(input *model.ListParam) (output []*model.Table, page *paging.Output, err error) {
	output, amount, err := s.repository.List(input)
	if err != nil {
		return output, page, err
	}
	page = &paging.Output{}
	page.TotalCount = int(amount)
	page.TotalPage = util.Pagination(int(amount), input.Size)
	page.Page = input.Page
	page.Size = input.Size
	return output, page, err
}

func (s *service) APIGetCMSCourses(input *model.APIGetCMSCoursesInput) interface{} {
	// parser input
	param := model.ListParam{}
	if err := util.Parser(input, &param); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "SaleItem"},
		{Field: "SaleItem.ProductLabel"},
	}
	// 調用 repo
	result, amount, err := s.repository.List(&param)
	if err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSCoursesData{}
	if err := util.Parser(result, &data); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSCoursesOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging.TotalCount = int(amount)
	output.Paging.TotalPage = util.Pagination(int(amount), input.Size)
	output.Paging.Page = input.Page
	output.Paging.Size = input.Size
	return output
}
