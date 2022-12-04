package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner_order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/banner_order"
	"gorm.io/gorm"
)

type service struct {
	repository banner_order.Repository
}

func New(repository banner_order.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Creates(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	err = s.repository.Creates(items)
	return err
}

func (s *service) List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error) {
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

func (s *service) DeleteAll() (err error) {
	err = s.repository.DeleteAll()
	return err
}
