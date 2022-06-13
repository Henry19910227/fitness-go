package service

import (
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiet_CreateDiet(t *testing.T) {
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
	rdas := make([]*entity.RDA, 0)
	rda1 := entity.RDA{
		ID: 1,
		UserID: users[0].ID,
		TDEE: 2000,
		CreateAt: "2022-05-10 00:00:00",
	}
	rda2 := entity.RDA{
		ID: 2,
		UserID: users[0].ID,
		TDEE: 2500,
		CreateAt: "2022-06-10 00:00:00",
	}
	rdas = append(rdas, &rda1)
	rdas = append(rdas, &rda2)
	if err := gormTool.DB().Create(&rdas).Error; err != nil {
		t.Fatalf(err.Error())
	}
	//驗證 diet create
	dietService := NewMockDietService(gormTool)
	diet, errcode := dietService.CreateDiet(nil, users[0].ID, "2022-06-10")
	if errcode != nil {
		t.Fatalf(errcode.Msg())
	}
	//驗證排程日期是否正確
	assert.Equal(t, "2022-06-10 00:00:00", *diet.ScheduleAt)
	//驗證是否會選擇最新創建的 rda
	assert.Equal(t, 2500, diet.RDA.TDEE)
}
