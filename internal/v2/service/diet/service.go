package diet

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	repository "github.com/Henry19910227/fitness-go/internal/v2/repository/diet"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
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
	page.Page = input.Page
	page.Size = input.Size
	if input.Size != nil {
		page.TotalPage = util.PointerInt(util.Pagination(int(amount), *input.Size))
	}
	return output, page, err
}

func (s *service) Updates(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	// 查找須更新的資料
	itemMap := make(map[int64]*model.Table)
	dietIDs := make([]int64, 0)
	for _, item := range items {
		dietIDs = append(dietIDs, *item.ID)
		itemMap[*item.ID] = item
	}
	listInput := model.ListInput{}
	listInput.Wheres = []*whereModel.Where{
		{Query: "diets.id IN (?)", Args: []interface{}{dietIDs}},
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
	// 將須更新的值映射到table(過濾未查找到的course)
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
