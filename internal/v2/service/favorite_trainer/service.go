package favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/favorite_trainer"
	"time"
)

type service struct {
	repository favorite_trainer.Repository
}

func New(repository favorite_trainer.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(item *model.Table) (err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	err = s.repository.Create(item)
	return err
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
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
