package entity

type Meal struct {
	ID       int64   `gorm:"column:id"`        //餐食id
	DietID   int64   `gorm:"column:diet_id"`   //飲食紀錄id
	FoodID   int64   `gorm:"column:food_id"`   //食物id
	Type     int     `gorm:"column:type"`      //類型(1:/早餐/2:午餐/3:晚餐/4:點心)
	Amount   float64 `gorm:"column:amount"`    //數量
	CreateAt string  `gorm:"column:create_at"` //創建日期
}

func (Meal) TableName() string {
	return "meals"
}
