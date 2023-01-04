package api_cms_logout

import "github.com/Henry19910227/fitness-go/internal/v2/field/admin/required"

// Input /v2/cms/logout [POST]
type Input struct {
	required.IDField
}
