package food

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/gin-gonic/gin"
)

type Resolver interface {
	APIGetFoods(ctx *gin.Context, input *model.APIGetFoodsInput) interface{}
}
