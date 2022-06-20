package food

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"time"
)


func Generate(input *GenerateInput) []*Table {
	tables := make([]*Table, 0)
	for i := 1; i <= input.DataAmount; i++ {
		table := Table{}
		table.ID = util.PointerInt64(int64(i))
		table.UserID = nil
		table.FoodCategoryID = nil
		table.Source = util.PointerInt(1)
		table.Name = util.PointerString(fmt.Sprintf("food_%v", i))
		table.Calorie = util.PointerInt(i * 100)
		table.AmountDesc = util.PointerString(fmt.Sprintf("amount_desc_%v", i))
		table.IsDeleted = util.PointerInt(0)
		table.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		table.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		tables = append(tables, &table)
	}
	if input.UserID != nil {
		for _, item := range input.UserID {
			datas := tables[item.Start-1:item.End]
			for _, data := range datas {
				data.UserID = util.PointerInt64(item.Value.(int64))
			}
		}
	}
	if input.FoodCategoryID != nil {
		for _, item := range input.FoodCategoryID {
			datas := tables[item.Start-1:item.End]
			for _, data := range datas {
				data.FoodCategoryID = util.PointerInt64(item.Value.(int64))
			}
		}
	}
	if input.Source != nil {
		for _, item := range input.Source {
			datas := tables[item.Start-1:item.End]
			for _, data := range datas {
				data.Source = util.PointerInt(item.Value.(int))
			}
		}
	}
	if input.IsDeleted != nil {
		for _, item := range input.IsDeleted {
			datas := tables[item.Start-1:item.End]
			for _, data := range datas {
				data.IsDeleted = util.PointerInt(item.Value.(int))
			}
		}
	}
	return tables
}
