package max_rm_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/max_rm_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/max_rm_record"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository max_rm_record.Repository
}

func New(repository max_rm_record.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	if err != nil {
		return output, err
	}
	return output, err
}

func (s *service) List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error) {
	output, amount, err := s.repository.List(input)
	if err != nil {
		return output, page, err
	}
	page = &paging.Output{}
	page.TotalCount = int(amount)
	page.Page = input.Page
	page.Size = input.Size
	if input.Size != nil {
		page.TotalPage = util.PointerInt(util.Pagination(int(amount), *input.Size))
	}
	return output, page, err
}

func (s *service) Create(item *model.Table) (id int64, err error) {
	item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.Create(item)
	return id, err
}

func (s *service) Update(item *model.Table) (err error) {
	input := model.FindInput{}
	input.UserID = item.UserID
	input.ActionID = item.ActionID
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

func (s *service) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	listInput := model.ListInput{}
	listInput.UserID = item.UserID
	listInput.ActionID = item.ActionID
	outputs, _, err := s.List(&listInput)
	if err != nil {
		return nil, err
	}
	if len(outputs) == 0 {
		id, err := s.Create(item)
		return &id, err
	}
	err = s.Update(item)
	return id, err
}
