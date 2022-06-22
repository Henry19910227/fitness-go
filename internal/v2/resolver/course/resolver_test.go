package course

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/migrate"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ids := []int64{1, 2, 3}
	items := make([]*int64, 0)
	for _, id := range ids {
		items = append(items, &id)
	}
	for _, item := range items {
		fmt.Println(*item)
	}
}

func Test2(t *testing.T) {
	ids := []int64{1, 2, 3}
	items := make([]*int64, 0)
	for _, id := range ids {
		i := id
		items = append(items, &i)
	}
	for _, item := range items {
		fmt.Println(*item)
	}
}

func Test3(t *testing.T) {
	ids := []int64{1, 2, 3}
	items := make([]*int64, 0)
	items = append(items, &ids[0])
	items = append(items, &ids[1])
	items = append(items, &ids[2])
	for _, item := range items {
		fmt.Println(*item)
	}
}

func Test4(t *testing.T) {
	ids := []int64{1, 2, 3}
	items := make([]*int64, 0)
	for i := 0; i < len(ids); i++ {
		items = append(items, &ids[i])
	}
	for _, item := range items {
		fmt.Println(*item)
	}
}

func TestResolver_APIUpdateCMSCoursesStatus(t *testing.T) {
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
	// 創建course
	courses := course.Generate(&course.GenerateInput{
		DataAmount: 3,
		UserID: []*base.GenerateSetting{
			{Start: 1, End: 1, Value: *users[0].ID},
			{Start: 2, End: 2, Value: *users[1].ID},
			{Start: 3, End: 3, Value: *users[2].ID},
		},
	})
	if err := prepareTx.Create(&courses).Error; err != nil {
		t.Fatalf(err.Error())
	}
	prepareTx.Commit()
	// 測試 APIUpdateCMSCoursesStatus
	db1 := orm.NewMockTool().DB()
	resolver := NewResolver(db1)
	input := course.APIUpdateCMSCoursesStatusInput{}
	input.IDs = []int64{*courses[0].ID, *courses[1].ID, 5}
	input.CourseStatus = 3
	output := resolver.APIUpdateCMSCoursesStatus(&input)

	// 驗證是否修改成功
	assert.Equal(t, code.Success, output.Code)

	// 驗證course[0]狀態
	var status int
	if err := db1.Table("courses").
		Select("course_status").
		Take(&status, "id = ?", *courses[0].ID).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, 3, status)

	// 驗證course[1]狀態
	if err := db1.Table("courses").
		Select("course_status").
		Take(&status, "id = ?", *courses[1].ID).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, 3, status)

	// 驗證course[2]狀態
	if err := db1.Table("courses").
		Select("course_status").
		Take(&status, "id = ?", *courses[2].ID).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, 1, status)

	// 驗證空陣列
	input = course.APIUpdateCMSCoursesStatusInput{}
	input.IDs = []int64{}
	input.CourseStatus = 3
	output = resolver.APIUpdateCMSCoursesStatus(&input)
	assert.Equal(t, code.Success, output.Code)
}
