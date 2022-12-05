package api_create_cms_banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/banner/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/banner [POST]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.IDField
	optional.CourseIDField
	optional.UserIDField
	optional.UrlField
	optional.ImageField
	optional.TypeField
	optional.CreateAtField
	optional.UpdateAtField
}
