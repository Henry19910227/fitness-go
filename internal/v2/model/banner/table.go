package banner

import "github.com/Henry19910227/fitness-go/internal/v2/field/banner/optional"

type Table struct {
	optional.IDField
	optional.CourseIDField
	optional.UserIDField
	optional.UrlField
	optional.ImageField
	optional.TypeField
	optional.CreateAtField
	optional.UpdateAtField
}
func (Table) TableName() string {
	return "banners"
}
