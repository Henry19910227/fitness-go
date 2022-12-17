package api_get_ios_version

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/ios_version [GET]
type Input struct {
	userRequired.UserIDField
}
