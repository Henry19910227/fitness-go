package api_delete_cms_banner

import "github.com/Henry19910227/fitness-go/internal/v2/field/banner/required"

// Input /v2/cms/banner/{banner_id} [DELETE]
type Input struct {
	Uri Uri
}
type Uri struct {
	required.BannerIDField
}
