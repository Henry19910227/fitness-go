package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/repository/favorite_course"
	"time"
)

type service struct {
	repository favorite_course.Repository
}

func New(repository favorite_course.Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(item *model.Table) (err error) {
	item.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	err = s.repository.Create(item)
	return err
}

func (s *service) Delete(input *model.DeleteInput) (err error) {
	err = s.repository.Delete(input)
	return err
}
