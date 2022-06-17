package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/food"
)

type service struct {
	repository food.Repository
}

func (s service) WithTrx(gormTool orm.Tool) {
	//TODO implement me
	panic("implement me")
}

func New(repository food.Repository) Service {
	return &service{repository: repository}
}

func (s service) List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error) {
	input.IsDeleted = util.PointerInt(0)
	input.OrderField = "create_at"
	input.OrderType = "DESC"
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
