package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/bank_account"
)

type service struct {
	repository bank_account.Repository
}

func New(repository bank_account.Repository) Service {
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
