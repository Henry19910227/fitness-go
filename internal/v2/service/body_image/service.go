package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	body "github.com/Henry19910227/fitness-go/internal/v2/repository/body_image"
)

type service struct {
	repository body.Repository
}

func New(repository body.Repository) Service {
	return &service{repository: repository}
}

func (s *service) List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error) {
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
