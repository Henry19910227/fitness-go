package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course"
	"time"
)

type service struct {
	repository course.Repository
}

func New(repository course.Repository) Service {
	return &service{repository: repository}
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

func (s *service) Update(items []*model.Table) (err error) {
	// 查找須更新的資料
	itemMap := make(map[int64]*model.Table)
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
	err = s.repository.Update(tables)
	return err
}
