package course

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/entity/course"
	"time"
)

func Generate(input *GenerateInput) []*course.Table {
	tables := make([]*course.Table, 0)
	createDate, _ := time.Parse("2006-01-02 15:04:05", "2021-12-31 00:00:00")
	for i := 1; i <= input.DataAmount; i++ {
		createDate = createDate.AddDate(0, 0, 1)
		table := course.Table{}
		table.ID = util.PointerInt64(int64(i))
		table.SaleType = util.PointerInt(1)
		table.CourseStatus = util.PointerInt(1)
		table.Category = util.PointerInt(1)
		table.ScheduleType = util.PointerInt(1)
		table.Name = util.PointerString(fmt.Sprintf("course_%v", i))
		table.Cover = util.PointerString(fmt.Sprintf("%v.jpg", i))
		table.Intro = util.PointerString(fmt.Sprintf("intro_%v", i))
		table.Food = util.PointerString(fmt.Sprintf("food_%v", i))
		table.Level = util.PointerInt(1)
		table.Suit = util.PointerString(fmt.Sprintf("1"))
		table.Equipment = util.PointerString(fmt.Sprintf("1"))
		table.Place = util.PointerString(fmt.Sprintf("1"))
		table.TrainTarget = util.PointerString(fmt.Sprintf("1"))
		table.BodyTarget = util.PointerString(fmt.Sprintf("1"))
		table.Notice = util.PointerString(fmt.Sprintf("notice_%v", i))
		table.PlanCount = util.PointerInt(1)
		table.WorkoutCount = util.PointerInt(1)
		table.CreateAt = util.PointerString(createDate.Format("2006-01-02 15:04:05"))
		table.UpdateAt = util.PointerString(createDate.Format("2006-01-02 15:04:05"))
		tables = append(tables, &table)
	}
	if input.UserID != nil {
		for _, item := range input.UserID {
			datas := tables[item.Start-1 : item.End]
			for _, data := range datas {
				data.UserID = util.PointerInt64(item.Value.(int64))
			}
		}
	}
	return tables
}
