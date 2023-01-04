package api_cms_login

import "github.com/Henry19910227/fitness-go/internal/v2/field/admin/required"

// Input /v2/cms/login [POST]
type Input struct {
	Body Body
}
type Body struct {
	required.EmailField
	required.PasswordField
}
