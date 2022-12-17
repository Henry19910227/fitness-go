package api_get_ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/ios_version/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/ios_version [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.VersionField
}
