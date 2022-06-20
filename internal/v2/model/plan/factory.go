package plan

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"time"
)

func Generate(input *GenerateInput) []*Table {
	tables := make([]*Table, 0)
	createDate, _ := time.Parse("2006-01-02 15:04:05", "2021-12-31 00:00:00")
	for i := 1; i <= input.DataAmount; i++ {
		createDate = createDate.AddDate(0, 0, 1)
		table := Table{}
		table.ID = util.PointerInt64(int64(i))
		table.Name = util.PointerString(fmt.Sprintf("plan_%v", i))
		table.WorkoutCount = util.PointerInt(1)
		table.CreateAt = util.PointerString(createDate.Format("2006-01-02 15:04:05"))
		table.UpdateAt = util.PointerString(createDate.Format("2006-01-02 15:04:05"))
		tables = append(tables, &table)
	}
	if input.CourseID != nil {
		for _, item := range input.CourseID {
			datas := tables[item.Start-1 : item.End]
			for _, data := range datas {
				data.CourseID = util.PointerInt64(item.Value.(int64))
			}
		}
	}
	return tables
}
