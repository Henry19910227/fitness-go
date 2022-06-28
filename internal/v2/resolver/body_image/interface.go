package body_image

import model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"

type Resolver interface {
	APIGetBodyImages(input *model.APIGetBodyImagesInput) (output model.APIGetBodyImagesOutput)
	APICreateBodyImage(input *model.APICreateBodyImageInput) (output model.APICreateBodyImageOutput)
}
