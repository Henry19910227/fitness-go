package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food"
	mealModel "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestF(t *testing.T)  {
	migrate.Mock().Down(nil)
}

func TestResolver_APIPutMeals(t *testing.T) {
	//設定 migrate
	if err := migrate.Mock().Up(nil); err != nil {
		t.Fatalf(err.Error())
	}
	defer func() {
		if err := migrate.Mock().Down(nil); err != nil {
			t.Fatalf(err.Error())
		}
	}()
	//準備資料
	prepareDB := orm.NewMockTool().DB()
	prepareTx := prepareDB.Begin()
	defer prepareTx.Rollback()
	// 創建user
	users := user.Generate(&user.GenerateInput{
		DataAmount: 3,
	})
	if err := prepareTx.Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet
	diets := diet.Generate(&diet.GenerateInput{
		DataAmount: 5,
		UserID: []*base.GenerateSetting{{Start: 1, End: 5, Value: *users[0].ID}},
	})
	if err := prepareTx.Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建food
	foods := food.Generate(&food.GenerateInput{
		DataAmount: 6,
		UserID: []*base.GenerateSetting{{Start: 4, End: 6, Value: *users[0].ID}},
		Source: []*base.GenerateSetting{{Start: 1, End: 3, Value: 1}, {Start: 4, End: 6, Value: 2}},
	})
	if err := prepareTx.Create(&foods).Error; err != nil {
		t.Fatalf(err.Error())
	}
	prepareTx.Commit()

	// 建立測試meals input
	items := make([]*mealModel.APIPutMealsInputItem, 0)
	item := mealModel.APIPutMealsInputItem{}
	item.FoodID = foods[0].ID
	item.Type = util.PointerInt(1)
	item.Amount = util.PointerFloat64(0.5)
	items = append(items, &item)

	// 驗證加入一筆meal
	db1 := orm.NewMockTool().DB()
	resolver := NewResolver(db1)
	input := mealModel.APIPutMealsInput{}
	input.UserID = util.OnNilJustReturnInt64(users[0].ID, 0)
	input.DietID = diets[0].ID
	input.Meals = items
	tx1 := db1.Begin()
	output := resolver.APIPutMeals(tx1, &input)
	assert.Equal(t, code.Success, output.Code)
	//驗證加入空meal
	db2 := orm.NewMockTool().DB()
	resolver = NewResolver(db2)
	input2 := mealModel.APIPutMealsInput{}
	input2.UserID = util.OnNilJustReturnInt64(users[0].ID, 0)
	input2.DietID = diets[0].ID
	input2.Meals = make([]*mealModel.APIPutMealsInputItem, 0)
	tx2 := db2.Begin()
	output = resolver.APIPutMeals(tx2, &input2)
	assert.Equal(t, code.Success, output.Code)
	//驗證非本人編輯飲食紀錄
	db3 := orm.NewMockTool().DB()
	resolver = NewResolver(db3)
	input3 := mealModel.APIPutMealsInput{}
	input3.UserID = util.OnNilJustReturnInt64(users[1].ID, 0)
	input3.DietID = diets[0].ID
	input3.Meals = items
	tx3 := db3.Begin()
	output = resolver.APIPutMeals(tx3, &input3)
	assert.Equal(t, code.PermissionDenied, output.Code)
}