package body_record

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"time"
)

func Generate(input *GenerateInput) []*Table {
	tables := make([]*Table, 0)
	nowDate := time.Now()
	for i := 1; i <= input.DataAmount; i++ {
		nowDate = nowDate.AddDate(0, 0, 1)
		table := Table{}
		table.ID = util.PointerInt64(int64(i))
		table.UserID = nil
		table.RecordType = util.PointerInt(1)
		table.Value = util.PointerFloat64(float64(0))
		table.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
		table.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
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
	if input.RecordType != nil {
		for _, item := range input.RecordType {
			datas := tables[item.Start-1 : item.End]
			for _, data := range datas {
				data.RecordType = util.PointerInt(item.Value.(int))
			}
		}
	}
	if input.Value != nil {
		for _, item := range input.Value {
			datas := tables[item.Start-1 : item.End]
			for _, data := range datas {
				data.Value = util.PointerFloat64(item.Value.(float64))
			}
		}
	}
	return tables
}
