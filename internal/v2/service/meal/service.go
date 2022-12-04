package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/meal"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository meal.Repository
}

func New(repository meal.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	for _, item := range items {
		item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	err = s.repository.Create(items)
	return err
}

func (s *service) List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error) {
	input.OrderField = "create_at"
	input.OrderType = orderBy.DESC
	output, amount, err := s.repository.List(input)
	if err != nil {
		return output, page, err
	}
	page = &paging.Output{}
	page.TotalCount = int(amount)
	page.TotalPage = util.PointerInt(util.Pagination(int(amount), input.Size))
	page.Page = util.PointerInt(input.Page)
	page.Size = util.PointerInt(input.Size)
	return output, page, err
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
}
