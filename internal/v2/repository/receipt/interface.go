package receipt

import model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"

type Repository interface {
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
