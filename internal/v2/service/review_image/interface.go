package review_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
)

type Service interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	Delete(input *model.DeleteInput) (err error)
}
