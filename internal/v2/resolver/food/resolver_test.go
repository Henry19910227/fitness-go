package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestF(t *testing.T) {
	if err := migrate.Mock().Down(nil); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestResolver_APIGetFoods(t *testing.T) {
	// 設定 migrate
	if err := migrate.Mock().Up(nil); err != nil {
		t.Fatalf(err.Error())
	}
	defer func() {
		if err := migrate.Mock().Down(nil); err != nil {
			t.Fatalf(err.Error())
		}
	}()
	// 準備資料
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
	// 創建food
	foods := food.Generate(&food.GenerateInput{
		DataAmount: 4,
		UserID:     []*base.GenerateSetting{{Start: 4, End: 4, Value: *users[0].ID}},
		Source:     []*base.GenerateSetting{{Start: 1, End: 3, Value: 1}, {Start: 4, End: 4, Value: 2}},
		FoodCategoryID: []*base.GenerateSetting{
			{Start: 1, End: 1, Value: int64(1)},
			{Start: 2, End: 2, Value: int64(4)},
			{Start: 3, End: 3, Value: int64(10)},
			{Start: 4, End: 4, Value: int64(7)},
		},
	})
	if err := prepareTx.Create(&foods).Error; err != nil {
		t.Fatalf(err.Error())
	}
	prepareTx.Commit()
	// 測試 APIGetFoods
	db1 := orm.NewMockTool().DB()
	resolver := NewResolver(db1)
	input := food.APIGetFoodsInput{}
	input.UserID = users[0].ID
	input.Tag = util.PointerInt(2)
	output := resolver.APIGetFoods(&input)
	assert.Equal(t, 2, len(output.Data))
	assert.Equal(t, int64(4), *output.Data[0].ID)
	assert.Equal(t, int64(10001), *output.Data[0].UserID)
	assert.Equal(t, "amount_desc_4", *output.Data[0].AmountDesc)
}
