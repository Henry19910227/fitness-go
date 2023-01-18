package api_get_cms_user

import (
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/cms/user/{user_id} [GET]
type Input struct {
	Uri Uri
}
type Uri struct {
	userRequired.UserIDField
}
