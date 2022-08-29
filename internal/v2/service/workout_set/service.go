package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/repository/workout_set"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository workoutSet.Repository
}

func New(repository workoutSet.Repository) Service {
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
		item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	ids, err = s.repository.Create(items)
	return ids, err
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	return output, err
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

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
}

