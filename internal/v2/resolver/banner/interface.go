package banner

import model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"

type Resolver interface {
	APICreateCMSBanner(input *model.APICreateCMSBannerInput) (output model.APICreateCMSBannerOutput)
}
