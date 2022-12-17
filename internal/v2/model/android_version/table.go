package android_version

import "github.com/Henry19910227/fitness-go/internal/v2/field/android_version/optional"

type Table struct {
	optional.IDField
	optional.VersionField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "android_versions"
}

