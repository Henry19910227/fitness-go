package diet

import model "github.com/Henry19910227/fitness-go/internal/v2/model/diet"

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
}
