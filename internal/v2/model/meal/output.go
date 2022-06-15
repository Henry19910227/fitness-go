package meal

import "github.com/Henry19910227/fitness-go/internal/v2/model/food"

type Output struct {
	Table
	Food *food.Table `json:"food,omitempty" gorm:"foreignKey:id;references:food_id"` // 食物
}

func (Output) TableName() string {
	return "meals"
}
