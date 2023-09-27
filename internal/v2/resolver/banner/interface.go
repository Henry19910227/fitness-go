package banner

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_create_cms_banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_delete_cms_banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_get_cms_banners"
)

type Resolver interface {
	APIGetBanners(input *model.APIGetBannersInput) (output model.APIGetBannersOutput)
	APICreateCMSBanner(input *api_create_cms_banner.Input) (output api_create_cms_banner.Output)
	APIGetCMSBanners(input *api_get_cms_banners.Input) (output api_get_cms_banners.Output)
	APIDeleteCMSBanner(input *api_delete_cms_banner.Input) (output api_delete_cms_banner.Output)
}
