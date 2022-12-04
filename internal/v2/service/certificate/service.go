package certificate

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/certificate"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository certificate.Repository
}

func New(repository certificate.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(item *model.Table) (id int64, err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.Create(item)
	return id, err
}

func (s *service) Creates(items []*model.Table) (err error) {
	for _, item := range items {
		item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	err = s.repository.Creates(items)
	return err
}

func (s *service) Updates(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	// 查找須更新的資料
	itemMap := make(map[int64]*model.Table)
	certIDs := make([]int64, 0)
	for _, item := range items {
		certIDs = append(certIDs, *item.ID)
		itemMap[*item.ID] = item
	}
	listInput := model.ListInput{}
	listInput.Wheres = []*whereModel.Where{
		{Query: "certificates.id IN (?)", Args: []interface{}{certIDs}},
	}
	outputs, _, err := s.repository.List(&listInput)
	if err != nil {
		return err
	}
	// 將output轉換為table
	tables := make([]*model.Table, 0)
	err = util.Parser(outputs, &tables)
	if err != nil {
		return err
	}
	// 將須更新的值映射到table
	for _, table := range tables {
		item := itemMap[*table.ID]
		err = util.Parser(item, &table)
		if err != nil {
			return err
		}
	}
	// 設置當前修改時間
	for _, table := range tables {
		table.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	}
	// 更新資料
	err = s.repository.Updates(tables)
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

func (s *service) Deletes(inputs []*model.DeleteInput) (err error) {
	err = s.repository.Deletes(inputs)
	return err
}
