package api_get_food_categories

import (
	foodCategoryOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food_category/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/food_categories [GET] 獲取食物分類
type Output struct {
	base.Output
	Data   *Data          `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	foodCategoryOptional.IDField
	foodCategoryOptional.TagField
	foodCategoryOptional.TitleField
	foodCategoryOptional.CreateAtField
	foodCategoryOptional.UpdateAtField
}
