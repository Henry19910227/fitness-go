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
	page.TotalPage = util.PointerInt(util.Pagination(int(amount), input.Size))
	page.Page = util.PointerInt(input.Page)
	page.Size = util.PointerInt(input.Size)
	return output, page, err
}

func (s *service) LatestList(input *model.LatestListInput) (outputs []*model.Output, err error) {
	outputs, err = s.repository.LatestList(input)
	return outputs, err
}

func (s *service) Update(item *model.Table) (err error) {
	input := model.FindInput{}
	input.ID = item.ID
	output, err := s.repository.Find(&input)
	if err != nil {
		return err
	}
	// 將output轉換為table
	var table model.Table
	err = util.Parser(output, &table)
	if err != nil {
		return err
	}
	// 將須更新的值映射到table
	err = util.Parser(item, &table)
	if err != nil {
		return err
	}
	// 設置當前修改時間
	table.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	// 更新資料
	err = s.repository.Update(&table)
	return err
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
}
