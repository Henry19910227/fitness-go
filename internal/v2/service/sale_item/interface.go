package sale_item

import model "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"

type Service interface {
	Find(input *model.FindInput) (output *model.Output, err error)
}
