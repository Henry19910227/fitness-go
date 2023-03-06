package user_course_statistic

import "gorm.io/gorm"

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	//Find(input *model.FindInput) (output *model.Output, err error)
	//List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	//Create(item *model.Table) (id int64, err error)
	//Update(item *model.Table) (err error)
}
