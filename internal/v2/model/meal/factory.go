package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"time"
)

func Generate(input *GenerateInput) []*Table {
	tables := make([]*Table, 0)
	nowDate, _ := time.Parse("2006-01-02 15:04:05", "2021-12-31 00:00:00")
	for i := 1; i <= input.DataAmount; i++ {
		nowDate = nowDate.AddDate(0, 0, 1)
		table := Table{}
		table.ID = util.PointerInt64(int64(i))
		table.Type = util.PointerInt(1)
		table.Amount = util.PointerFloat64(0.5)
		table.CreateAt = util.PointerString(nowDate.Format("2006-01-02 15:04:05"))
		tables = append(tables, &table)
	}
	if input.DietID != nil {
		for _, item := range input.DietID {
			datas := tables[item.Start-1 : item.End]
			for _, data := range datas {
				data.DietID = util.PointerInt64(item.Value.(int64))
			}
		}
	}
	if input.FoodID != nil {
		for _, item := range input.FoodID {
			datas := tables[item.Start-1 : item.End]
			for _, data := range datas {
				data.FoodID = util.PointerInt64(item.Value.(int64))
			}
		}
	}
	if input.Type != nil {
		for _, item := range input.Type {
			datas := tables[item.Start-1 : item.End]
			for _, data := range datas {
				data.Type = util.PointerInt(item.Value.(int))
			}
		}
	}
	return tables
}
