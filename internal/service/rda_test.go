package service

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//飲食計算機條件：
//標準飲食
//性別：男性
//年齡：30（生日：1992/2/2）
//身高：178cm
//體重：70kg
//體脂：20%
//活動量：每週輕度步行3-4天（1.375）
//運動量：中度 一週3-5次，一次45-60分鐘 （+300）
//飲食目標：增肌
//
//
//計算結果：
//TDEE=2533
//建議每日攝取熱量=2913
//蛋白質146克、脂肪65克、醣類437克
//乳製品1份、蔬菜5份、水果2份、全穀雜糧25份、蛋豆魚肉12份、油脂堅果5份

//—————————————以下為計算公式—————————————
//男性（有體脂）BMR公式={[10*體重(kg)+6.25*身高(cm)-5*年齡+5] + [370+21.6*(100% - 體脂率) * 體重]}/2
//BMR= {[10*70+6.25*178-5*30+5] + [370+21.6*(0.8) * 70]}/2= 1624 (1623.55 四捨五入）
//TDEE公式 = BMR * 活動量因子 + 運動量
//1624 * 1.375 + 300 = 2533
//飲食目標增肌建議每日攝取熱量 = TDEE *1.15
//2533*1.15 = 2913 (2912.95 四捨五入）
//增肌-三大營養素克數: 蛋白質20%、脂肪20%、醣類60%
//2913*0.2/4 = 146
//2913*0.2/9 = 65
//2913*0.6/4 = 437
//蛋白質146克、脂肪65克、醣類437克
//a乳製品 1
//b蔬菜 5
//c水果 2
//d全穀雜糧 = (437 – (1 * 12 + 5 * 5 + 2 * 15)) / 15 = 25
//e蛋豆魚肉 = (146 – (1 * 8 + 25 * 2 + 5 * 1)) / 7 = 12
//f油脂堅果 = (65 – (1 * 4 + 12 * 3)) / 5 = 5

func TestRda_CalculateRDA(t *testing.T) {
	gormTool, err := tool.NewGorm(setting.NewMockMysql())
	if err != nil {
		t.Fatalf(err.Error())
	}
	rdaService := NewMockRdaService(gormTool)
	rda := rdaService.CalculateRDA(&dto.CalculateRDAParam{
		DietType:         1,
		Sex:              "m",
		Birthday:         "1992-02-02",
		Height:           178,
		Weight:           70,
		BodyFat:          util.PointerInt(20),
		ActivityLevel:    6,
		ExerciseFeqLevel: 3,
		DietTarget:       2,
	})
	assert.Equal(t, 2533, rda.TDEE)
	assert.Equal(t, 2913, rda.Calorie)
	assert.Equal(t, 146, rda.Protein)
	assert.Equal(t, 65, rda.Fat)
	assert.Equal(t, 437, rda.Carbs)
	assert.Equal(t, 1, rda.Dairy)
	assert.Equal(t, 5, rda.Vegetable)
	assert.Equal(t, 2, rda.Fruit)
	assert.Equal(t, 25, rda.Grain)
	assert.Equal(t, 12, rda.Meat)
	assert.Equal(t, 5, rda.Nut)
}

// 驗證 CreateRDA 創建時，是否會修改今天之後的 diet rdaID
func TestRda_CreateRDA(t *testing.T) {
	if err := PrepareMigrate(); err != nil {
		t.Fatalf(err.Error())
	}
	gormTool, sqlDB, err := PrepareGorm()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer sqlDB.Close()
	// 創建user
	users := entity.NewMockUsers()
	if err := gormTool.DB().Create(users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建rda
	rda := entity.RDA{
		ID: 1,
		UserID: users[0].ID,
		TDEE: 2000,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := gormTool.DB().Create(&rda).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet
	diets := make([]*entity.Diet, 0)
	diet1 := entity.Diet{
		ID: 1,
		UserID: users[0].ID,
		RdaID: rda.ID,
		ScheduleAt: "2020-01-01",
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	diet2 := entity.Diet{
		ID: 2,
		UserID: users[0].ID,
		RdaID: rda.ID,
		ScheduleAt: "2030-01-01",
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	diets = append(diets, &diet1)
	diets = append(diets, &diet2)
	if err := gormTool.DB().Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	//測試CreateRDA
	rdaService := NewMockRdaService(gormTool)
	if e := rdaService.CreateRDA(nil, 10001, &dto.RDA{
		TDEE: 2500,
	}); e != nil {
		t.Fatalf(e.Msg())
	}
	var tdee int
	//驗證今天以前的 diet rda tdee 不會被修改
	if err := gormTool.DB().
		Table("diets").
		Select("rdas.tdee").
		Joins("INNER JOIN rdas ON diets.rda_id = rdas.id").
		Where("diets.id = ?", 1).
		Scan(&tdee).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, 2000, tdee)
	//驗證未來的 diet rda tdee 不會被修改
	if err := gormTool.DB().
		Table("diets").
		Select("rdas.tdee").
		Joins("INNER JOIN rdas ON diets.rda_id = rdas.id").
		Where("diets.id = ?", 2).
		Scan(&tdee).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, 2500, tdee)
}