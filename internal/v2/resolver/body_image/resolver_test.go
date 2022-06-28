package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"
	"github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolver_APICreateBodyImage(t *testing.T) {
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
	// 創建body_record
	records := body_record.Generate(&body_record.GenerateInput{
		DataAmount: 6,
		UserID: []*base.GenerateSetting{
			{Start: 1, End: 3, Value: *users[0].ID},
			{Start: 4, End: 6, Value: *users[1].ID},
		},
		Value: []*base.GenerateSetting{
			{Start: 3, End: 3, Value: float64(50)},
		},
	})
	if err := prepareTx.Create(&records).Error; err != nil {
		t.Fatalf(err.Error())
	}
	//建立 resolver
	db1 := orm.NewMockTool().DB()
	resolver := NewResolver(db1)
	//測試 APICreateBodyImage
	input := model.APICreateBodyImageInput{}
	input.UserID = *users[0].ID
	output := resolver.APICreateBodyImage(&input)
	assert.Equal(t, code.Success, output.Code)
}
