package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user"
	"time"
)

type service struct {
	repository user.Repository
}

func New(repository user.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Find(input *model.FindInput) (output *model.Output, err error) {
	output, err = s.repository.Find(input)
	if err != nil {
		return output, err
	}
	return output, err
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
