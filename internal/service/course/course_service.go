package course

import (
	model "github.com/Henry19910227/fitness-go/internal/model/course"
	"github.com/Henry19910227/fitness-go/internal/model/paging"
	"github.com/Henry19910227/fitness-go/internal/repository/course"
	"github.com/Henry19910227/fitness-go/internal/util"
)

type service struct {
	repository course.Repository
}

func New(repository course.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Find(input *model.FindInput) (output *model.Table, err error) {
	output, err = s.repository.Find(input)
	if err != nil {
		return output, err
	}
	return output, err
}

func (s *service) List(input *model.ListInput) (output []*model.Table, page *paging.Output, err error) {
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
