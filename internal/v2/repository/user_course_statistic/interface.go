package user_course_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_statistic"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Statistic(input *model.Statistic) (err error)
	//Find(input *model.FindInput) (output *model.Output, err error)
	//List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	//Create(item *model.Table) (id int64, err error)
	//Update(item *model.Table) (err error)
}
