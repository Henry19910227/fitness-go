package banner

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_get_cms_banners"
)

type Resolver interface {
	APIGetBanners(input *model.APIGetBannersInput) (output model.APIGetBannersOutput)
	APICreateCMSBanner(input *model.APICreateCMSBannerInput) (output model.APICreateCMSBannerOutput)
	APIGetCMSBanners(input *api_get_cms_banners.Input) (output api_get_cms_banners.Output)
	APIDeleteCMSBanner(input *model.APIDeleteCMSBannerInput) (output model.APIDeleteCMSBannerOutput)
}
