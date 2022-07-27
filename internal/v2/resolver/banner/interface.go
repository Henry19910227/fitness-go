package banner

import model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"

type Resolver interface {
	APIGetBanners(input *model.APIGetBannersInput) (output model.APIGetBannersOutput)
	APICreateCMSBanner(input *model.APICreateCMSBannerInput) (output model.APICreateCMSBannerOutput)
	APIGetCMSBanners(input *model.APIGetCMSBannersInput) (output model.APIGetCMSBannersOutput)
	APIDeleteCMSBanner(input *model.APIDeleteCMSBannerInput) (output model.APIDeleteCMSBannerOutput)
}
