package diet

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRepository_Find(t *testing.T) {
	if err := migrate.Mock().Up(nil); err != nil {
		t.Fatalf(err.Error())
	}
	defer func() {
		if err := migrate.Mock().Down(nil); err != nil {
			t.Fatalf(err.Error())
		}
	}()
	prepare(t)
	repo := New(orm.Mock().DB())
	input := diet.FindInput{}
	input.ID = util.PointerInt64(1)
	output, err := repo.Find(&input)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, int64(1), *output.ID)
}

func prepare(t *testing.T)  {
	tx := orm.Mock().DB().Begin()
	defer tx.Rollback()
	// 創建user
	users := user.NewMockTables()
	if err := tx.Create(&users).Error; err != nil {
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
	if err := tx.Create(&diets).Error; err != nil {
		t.Fatalf(err.Error())
	}
	tx.Commit()
}
