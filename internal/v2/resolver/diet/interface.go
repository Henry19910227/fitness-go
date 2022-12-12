package diet

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet/api_create_diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet/api_get_diet"
)

type Resolver interface {
	APICreateDiet(input *api_create_diet.Input) (output api_create_diet.Output)
	APIGetDiet(input *api_get_diet.Input) (output api_get_diet.Output)
}
