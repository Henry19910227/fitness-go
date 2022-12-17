package api_get_android_version

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/android_version [GET]
type Input struct {
	userRequired.UserIDField
}
