package android_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/android_version/api_get_android_version"
)

type Resolver interface {
	APIGetAndroidVersion(input *api_get_android_version.Input) (output api_get_android_version.Output)
}
