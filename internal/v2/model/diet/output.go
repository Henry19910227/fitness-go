package diet

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda"
)

type Table struct {
	IDField
	UserIDField
	RdaIDField
	ScheduleAtField
	CreateAtField
	UpdateAtField
	Rda   *rda.Table  `json:"rda,omitempty" gorm:"foreignKey:rda_id;references:id"`    // rda
	Meals *meal.Table `json:"meals,omitempty" gorm:"foreignKey:diet_id;references:id"` // 餐食
}

func (Table) TableName() string {
	return "diets"
}
