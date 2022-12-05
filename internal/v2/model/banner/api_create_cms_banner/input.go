package api_create_cms_banner

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/banner/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/banner/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
)

// Input /v2/cms/banner [POST]
type Input struct {
	ImageFile *file.Input
	Form      Form
}
type Form struct {
	optional.CourseIDField
	optional.UserIDField
	optional.UrlField
	required.TypeField
}

