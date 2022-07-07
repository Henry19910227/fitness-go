package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/receipt"
)

type service struct {
	repository receipt.Repository
}

func New(repository receipt.Repository) Service {
	return &service{repository: repository}
}

func (s *service) List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error) {
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
