package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository_List(t *testing.T) {
	migTool := migrate.NewMockTool()
	err := migTool.Down(nil)
	if err != nil && err.Error() != "no change" {
		t.Fatalf(err.Error())
	}
	err = migTool.Up(nil)
	if err != nil && err.Error() != "no change" {
		t.Fatalf(err.Error())
	}
	// 創建user
	users := user.NewMockTables()
	if err := orm.Mock().DB().Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建food
	foods := make([]*food.Table, 0)
	for i := 1; i <= 10; i++ {
		foodItem := food.Table{}
		foodItem.ID = util.PointerInt64(int64(i))
		foodItem.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		foodItem.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		if i <= 5 {
			//用戶創建食物
			foodItem.UserID = users[0].ID
			foodItem.Source = util.PointerInt(2)
			foodItem.FoodCategoryID = util.PointerInt64(1)
		} else {
			//系統創建食物
			foodItem.Source = util.PointerInt(1)
			foodItem.FoodCategoryID = util.PointerInt64(4)
		}
		foods = append(foods, &foodItem)
	}
	if err := orm.Mock().DB().Create(&foods).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 驗證測試項目
	repo := New(orm.Mock().DB())
	input := food.ListInput{}
	input.UserID = users[0].ID
	input.Tag = util.PointerInt(1)
	input.Preloads = []*preload.Preload{{Field: "FoodCategory"}}
	input.Page = 3
	input.Size = 2
	input.OrderField = "create_at"
	input.OrderType = "DESC"
	input.IsDisabled = util.PointerInt(0)
	input.IsDeleted = util.PointerInt(0)
	outputs, amount, err := repo.List(&input)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, int64(5), amount)
	assert.Equal(t, 1, len(outputs))
}
