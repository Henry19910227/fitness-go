package ios_version

import "github.com/Henry19910227/fitness-go/internal/v2/model/ios_version/api_get_ios_version"

type Resolver interface {
	APIGetIOSVersion(input *api_get_ios_version.Input) (output api_get_ios_version.Output)
}
