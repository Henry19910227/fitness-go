package workout_set_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/workout_set_log"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository workout_set_log.Repository
}

func New(repository workout_set_log.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(items []*model.Table) (ids []int64, err error) {
	if len(items) == 0 {
		return ids, err
	}
	for _, item := range items {
		item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	ids, err = s.repository.Create(items)
	return ids, err
}

func (s *service) List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error) {
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
