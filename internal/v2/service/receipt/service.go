package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/receipt"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository receipt.Repository
}

func New(repository receipt.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
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

func (s *service) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.CreateOrUpdate(item)
	return id, err
}
