package api_get_cms_course_users

import (
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	courseAssetOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user_course_asset/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/cms/course/{course_id}/users [GET]
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	userOptional.IDField
	userOptional.NicknameField
	UserCourseAsset *struct {
		courseAssetOptional.SourceField
		courseAssetOptional.CreateAtField
		courseAssetOptional.UpdateAtField
	} `json:"user_course_asset,omitempty"`
}
