package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	course2 "github.com/Henry19910227/fitness-go/internal/v2/entity/course"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course"
	"gorm.io/gorm"
	"time"
)

type service struct {
	repository course.Repository
}

func New(repository course.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Tx(tx *gorm.DB) Service {
	return NewService(tx)
}

func (s *service) Create(item *course2.Table) (id int64, err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	item.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	id, err = s.repository.Create(item)
	return id, err
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
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
	page.TotalPage = util.Pagination(int(amount), input.Size)
	page.Page = input.Page
	page.Size = input.Size
	return output, page, err
}

func (s *service) FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, page *paging.Output, err error) {
	output, amount, err := s.repository.FavoriteList(input)
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

func (s *service) ProgressList(input *model.ProgressListInput) (outputs []*model.Output, page *paging.Output, err error) {
	output, amount, err := s.repository.ProgressList(input)
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

func (s *service) ChargeList(input *model.ChargeListInput) (outputs []*model.Output, page *paging.Output, err error) {
	output, amount, err := s.repository.ChargeList(input)
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

func (s *service) Updates(items []*course2.Table) (err error) {
	// 查找須更新的資料
	itemMap := make(map[int64]*course2.Table)
	courseIDs := make([]int64, 0)
	for _, item := range items {
		courseIDs = append(courseIDs, *item.ID)
		itemMap[*item.ID] = item
	}
	listInput := model.ListInput{}
	listInput.IDs = courseIDs
	outputs, _, err := s.repository.List(&listInput)
	if err != nil {
		return err
	}
	// 將output轉換為table
	tables := make([]*course2.Table, 0)
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

func (s *service) Update(item *course2.Table) (err error) {
	input := model.FindInput{}
	input.ID = item.ID
	output, err := s.repository.Find(&input)
	if err != nil {
		return err
	}
	// 將output轉換為table
	var table course2.Table
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
