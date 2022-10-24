package review_image

import "github.com/Henry19910227/fitness-go/internal/v2/field/review_image/optional"

type Table struct {
	optional.IDField
	optional.ReviewIDField
	optional.ImageField
	optional.CreateAtField
}

func (Table) TableName() string {
	return "review_images"
}
