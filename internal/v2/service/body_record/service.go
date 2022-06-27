package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	body "github.com/Henry19910227/fitness-go/internal/v2/repository/body_record"
	"time"
)

type service struct {
	repository body.Repository
}

func New(repository body.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(item *model.Table) (output *model.Output, err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err := s.repository.Create(item)
	if err != nil {
		return nil, err
	}
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(id)
	output, err = s.repository.Find(&findInput)
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
