package food_category

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"time"
)

func Generate(input *GenerateInput) []*Output {
	tables := make([]*Output, 0)
	for i := 1; i <= input.DataAmount; i++ {
		table := Output{}
		table.ID = util.PointerInt64(int64(i))
		table.Tag = util.PointerInt(1)
		table.Title = util.PointerString(fmt.Sprintf("category_%v", i))
		table.IsDeleted = util.PointerInt(0)
		table.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		table.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		tables = append(tables, &table)
	}
	if input.Tag != nil {
		for _, item := range input.Tag {
			datas := tables[item.Start-1:item.End]
			for _, data := range datas {
				data.Tag = util.PointerInt(item.Value.(int))
			}
		}
	}
	return tables
}
