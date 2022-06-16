package diet

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda"
)

type Output struct {
	Table
	Rda   *rda.Table   `json:"rda,omitempty" gorm:"foreignKey:rda_id;references:id"`    // rda
	Meals *meal.Output `json:"meals,omitempty" gorm:"foreignKey:diet_id;references:id"` // 餐食
}

func (Output) TableName() string {
	return "diets"
}
