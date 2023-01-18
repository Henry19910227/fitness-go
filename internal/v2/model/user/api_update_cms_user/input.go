package api_update_cms_user

import (
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/cms/user/{user_id} [PATCH]
type Input struct {
	Uri  Uri
	Body Body
}
type Uri struct {
	userRequired.UserIDField
}
type Body struct {
	userOptional.PasswordField
	userOptional.UserStatusField
}
