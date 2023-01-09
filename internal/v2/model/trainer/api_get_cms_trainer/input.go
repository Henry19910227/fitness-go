package api_get_cms_trainer

import "github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"

// Input /v2/cms/trainer/{user_id} [GET]
type Input struct {
	Uri Uri
}
type Uri struct {
	optional.UserIDField
}
