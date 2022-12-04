package user_subscribe_info

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user_subscribe_info"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository user_subscribe_info.Repository
}

func New(repository user_subscribe_info.Repository) Service {
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
	page.TotalPage = util.PointerInt(util.Pagination(int(amount), input.Size))
	page.Page = util.PointerInt(input.Page)
	page.Size = util.PointerInt(input.Size)
	return output, page, err
}

func (s *service) CreateOrUpdate(item *model.Table) (err error) {
	item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	err = s.repository.CreateOrUpdate(item)
	return err
}

func (s *service) Update(item *model.Table) (err error) {
	input := model.FindInput{}
	input.UserID = item.UserID
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

func (s *service) Updates(items []*model.Table) (err error) {
	if len(items) == 0 {
		return err
	}
	// 查找須更新的資料
	itemMap := make(map[int64]*model.Table)
	userIDs := make([]int64, 0)
	for _, item := range items {
		userIDs = append(userIDs, *item.UserID)
		itemMap[*item.UserID] = item
	}
	listInput := model.ListInput{}
	listInput.Wheres = []*whereModel.Where{
		{Query: "user_subscribe_infos.user_id IN (?)", Args: []interface{}{userIDs}},
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
		item := itemMap[*table.UserID]
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
