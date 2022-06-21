package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestF(t *testing.T) {
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
		DataAmount: 6,
		UserID: []*base.GenerateSetting{
			{Start: 1, End: 3, Value: *users[1].ID},
			{Start: 4, End: 6, Value: *users[0].ID},
		},
	})
	if err := prepareTx.Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建food
	foods := food.Generate(&food.GenerateInput{
		DataAmount: 6,
		UserID:     []*base.GenerateSetting{{Start: 4, End: 6, Value: *users[0].ID}},
		Source:     []*base.GenerateSetting{{Start: 1, End: 3, Value: 1}, {Start: 4, End: 6, Value: 2}},
	})
	if err := prepareTx.Create(&foods).Error; err != nil {
		t.Fatalf(err.Error())
	}
	prepareTx.Commit()

	// 建立測試meals input
	items := make([]*meal.APIPutMealsInputItem, 0)
	item := meal.APIPutMealsInputItem{}
	item.FoodID = foods[0].ID
	item.Type = util.PointerInt(1)
	item.Amount = util.PointerFloat64(0.5)
	items = append(items, &item)

	// 驗證加入一筆meal
	db1 := orm.NewMockTool().DB()
	resolver := NewResolver(db1)
	input := meal.APIPutMealsInput{}
	input.UserID = users[0].ID
	input.DietID = diets[5].ID
	input.Meals = items
	tx1 := db1.Begin()
	output := resolver.APIPutMeals(tx1, &input)
	assert.Equal(t, code.Success, output.Code)
	//驗證加入空meal
	db2 := orm.NewMockTool().DB()
	resolver = NewResolver(db2)
	input2 := meal.APIPutMealsInput{}
	input2.UserID = users[0].ID
	input2.DietID = diets[5].ID
	input2.Meals = make([]*meal.APIPutMealsInputItem, 0)
	tx2 := db2.Begin()
	output = resolver.APIPutMeals(tx2, &input2)
	assert.Equal(t, code.Success, output.Code)
	//驗證非本人編輯飲食紀錄
	db3 := orm.NewMockTool().DB()
	resolver = NewResolver(db3)
	input3 := meal.APIPutMealsInput{}
	input3.UserID = users[0].ID
	input3.DietID = diets[0].ID
	input3.Meals = items
	tx3 := db3.Begin()
	output = resolver.APIPutMeals(tx3, &input3)
	assert.Equal(t, code.PermissionDenied, output.Code)
}

func TestResolver_APIGetMeals(t *testing.T) {
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
	// 創建user(2個普通user)
	users := user.Generate(&user.GenerateInput{
		DataAmount: 2,
	})
	if err := prepareTx.Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet(1個user1飲食、1個user2飲食)
	diets := diet.Generate(&diet.GenerateInput{
		DataAmount: 2,
		UserID: []*base.GenerateSetting{
			{Start: 1, End: 1, Value: *users[0].ID},
			{Start: 2, End: 2, Value: *users[1].ID},
		},
	})
	if err := prepareTx.Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建food(4個系統食物、1個user1食物、1個user2食物)
	foods := food.Generate(&food.GenerateInput{
		DataAmount: 6,
		UserID: []*base.GenerateSetting{
			{Start: 5, End: 5, Value: *users[0].ID},
			{Start: 6, End: 6, Value: *users[1].ID},
		},
		Source: []*base.GenerateSetting{
			{Start: 1, End: 4, Value: 1},
			{Start: 5, End: 6, Value: 2},
		},
	})
	if err := prepareTx.Create(&foods).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建meal
	meals := meal.Generate(&meal.GenerateInput{
		DataAmount: 4,
		DietID: []*base.GenerateSetting{
			{Start: 1, End: 2, Value: *diets[0].ID},
			{Start: 3, End: 4, Value: *diets[1].ID},
		},
		FoodID: []*base.GenerateSetting{{Start: 1, End: 4, Value: *foods[0].ID}},
		Type: []*base.GenerateSetting{
			{Start: 1, End: 1, Value: 1},
			{Start: 2, End: 2, Value: 2},
			{Start: 3, End: 3, Value: 3},
			{Start: 4, End: 4, Value: 4},
		},
	})
	if err := prepareTx.Create(&meals).Error; err != nil {
		t.Fatalf(err.Error())
	}
	prepareTx.Commit()

	// 建立測試resolver
	db1 := orm.NewMockTool().DB()
	resolver := NewResolver(db1)
	input := meal.APIGetMealsInput{}
	input.UserID = users[0].ID
	output := resolver.APIGetMeals(&input)
	// 驗證狀態碼
	assert.Equal(t, code.Success, output.Code)
	// 驗證資料個數
	assert.Equal(t, 2, len(output.Data))
}
