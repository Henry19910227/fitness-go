package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestRepository_List(t *testing.T) {
	migTool := migrate.NewMockTool()
	err := migTool.Down(nil)
	if err != nil && err.Error() != "no change"{
		t.Fatalf(err.Error())
	}
	err = migTool.Up(nil)
	if err != nil && err.Error() != "no change" {
		t.Fatalf(err.Error())
	}
	gormTool := orm.NewMockTool()
	// 創建user
	users := user.NewMockTables()
	if err := gormTool.DB().Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet
	diets := make([]*diet.Table, 0)
	for i := 1; i <= 2; i++ {
		dietItem := diet.Table{}
		dietItem.ID = util.PointerInt64(int64(i))
		dietItem.UserID = users[0].ID
		dietItem.ScheduleAt = util.PointerString("2022-06" + "-0" + strconv.Itoa(i))
		diets = append(diets, &dietItem)
	}
	if err := gormTool.DB().Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建meal
	meals := make([]*meal.Table, 0)
	for i := 1; i <= 10; i++ {
		mealItem := meal.Table{}
		mealItem.ID = util.PointerInt64(int64(i))
		if i <= 5 {
			mealItem.DietID = diets[0].ID
		} else {
			mealItem.DietID = diets[1].ID
		}
		meals = append(meals, &mealItem)
	}
	if err := gormTool.DB().Create(&meals).Error; err != nil {
		t.Fatalf(err.Error())
	}
	repo := New(gormTool)
	input := meal.ListInput{}
	input.DietID = util.PointerInt64(1)
	input.Page = 3
	input.Size = 2
	outputs, amount, err := repo.List(&input)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, int64(5), amount)
	assert.Equal(t, 1, len(outputs))
}

func TestRepository_Create(t *testing.T) {
	migTool := migrate.NewMockTool()
	err := migTool.Down(nil)
	if err != nil && err.Error() != "no change"{
		assert.EqualError(t, err, "no change")
	}
	err = migTool.Up(nil)
	if err != nil && err.Error() != "no change" {
		assert.EqualError(t, err, "no change")
	}
	gormTool := orm.NewMockTool()
	// 創建user
	users := user.NewMockTables()
	if err := gormTool.DB().Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet
	diets := make([]*diet.Table, 0)
	for i := 1; i <= 2; i++ {
		dietItem := diet.Table{}
		dietItem.ID = util.PointerInt64(int64(i))
		dietItem.UserID = users[0].ID
		dietItem.ScheduleAt = util.PointerString("2022-06" + "-0" + strconv.Itoa(i))
		diets = append(diets, &dietItem)
	}
	if err := gormTool.DB().Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	//測試項目
	meals := make([]*meal.Table, 0)
	mealItem := meal.Table{}
	mealItem.DietID = diets[0].ID
	mealItem.Type = util.PointerInt(1)
	mealItem.Amount = util.PointerFloat64(0.5)
	mealItem.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	meals = append(meals, &mealItem)

	repo := New(gormTool)
	if err := repo.Create(meals); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRepository_Delete(t *testing.T) {
	migTool := migrate.NewMockTool()
	err := migTool.Down(nil)
	if err != nil && err.Error() != "no change"{
		t.Fatalf(err.Error())
	}
	err = migTool.Up(nil)
	if err != nil && err.Error() != "no change" {
		t.Fatalf(err.Error())
	}
	gormTool := orm.NewMockTool()
	// 創建user
	users := user.NewMockTables()
	if err := gormTool.DB().Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet
	diets := make([]*diet.Table, 0)
	for i := 1; i <= 2; i++ {
		dietItem := diet.Table{}
		dietItem.ID = util.PointerInt64(int64(i))
		dietItem.UserID = users[0].ID
		dietItem.ScheduleAt = util.PointerString("2022-06" + "-0" + strconv.Itoa(i))
		diets = append(diets, &dietItem)
	}
	if err := gormTool.DB().Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建meal
	meals := make([]*meal.Table, 0)
	for i := 1; i <= 10; i++ {
		mealItem := meal.Table{}
		mealItem.ID = util.PointerInt64(int64(i))
		if i <= 5 {
			mealItem.DietID = diets[0].ID
		} else {
			mealItem.DietID = diets[1].ID
		}
		meals = append(meals, &mealItem)
	}
	if err := gormTool.DB().Create(&meals).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 測試刪除meal
	repo := New(gormTool)
	input := meal.DeleteInput{}
	input.DietID = diets[0].ID
	if err := repo.Delete(&input); err != nil {
		t.Fatalf(err.Error())
	}
	var amount int64
	if err := gormTool.DB().Model(meal.Table{}).Count(&amount).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, int64(5), amount)
}

func TestRepository_Update(t *testing.T) {
	migTool := migrate.NewMockTool()
	err := migTool.Down(nil)
	if err != nil && err.Error() != "no change"{
		t.Fatalf(err.Error())
	}
	err = migTool.Up(nil)
	if err != nil && err.Error() != "no change" {
		t.Fatalf(err.Error())
	}
	gormTool := orm.NewMockTool()
	// 創建user
	users := user.NewMockTables()
	if err := gormTool.DB().Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet
	diets := make([]*diet.Table, 0)
	for i := 1; i <= 2; i++ {
		dietItem := diet.Table{}
		dietItem.ID = util.PointerInt64(int64(i))
		dietItem.UserID = users[0].ID
		dietItem.ScheduleAt = util.PointerString("2022-06" + "-0" + strconv.Itoa(i))
		diets = append(diets, &dietItem)
	}
	if err := gormTool.DB().Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建meal
	meals := make([]*meal.Table, 0)
	for i := 1; i <= 10; i++ {
		mealItem := meal.Table{}
		mealItem.ID = util.PointerInt64(int64(i))
		if i <= 5 {
			mealItem.DietID = diets[0].ID
		} else {
			mealItem.DietID = diets[1].ID
		}
		meals = append(meals, &mealItem)
	}
	if err := gormTool.DB().Create(&meals).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 測試更新meal
	m := meals[0]
	m.Amount = util.PointerFloat64(1000)
	repo := New(gormTool)
	if err := repo.Update([]*meal.Table{m}); err != nil {
		t.Fatalf(err.Error())
	}
	var mealItem meal.Table
	if err := gormTool.DB().Find(&mealItem, "id = ?", *m.ID).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, float64(1000), *mealItem.Amount)
}