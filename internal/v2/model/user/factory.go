package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"time"
)


func NewMockTables() []*Table {
	tables := make([]*Table, 0)
	data1 := Table{}
	data1.ID = util.PointerInt64(10010)
	data1.AccountType = util.PointerInt(1)
	data1.Nickname = util.PointerString("Henry")
	data1.Sex = util.PointerString("m")
	data1.Account = util.PointerString("henry@gmail.com")
	data1.Email = util.PointerString("henry@gmail.com")
	data1.Height = util.PointerFloat64(176)
	data1.Weight = util.PointerFloat64(70)
	data1.Birthday = util.PointerString(time.Now().Format("2006-01-02"))
	data1.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	data1.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))

	data2 := Table{}
	data2.ID = util.PointerInt64(10011)
	data2.AccountType = util.PointerInt(1)
	data2.Nickname = util.PointerString("Jeff")
	data2.Sex = util.PointerString("m")
	data2.Account = util.PointerString("jeff@gmail.com")
	data2.Email = util.PointerString("jeff@gmail.com")
	data2.Height = util.PointerFloat64(172)
	data2.Weight = util.PointerFloat64(65)
	data2.Birthday = util.PointerString(time.Now().Format("2006-01-02"))
	data2.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	data2.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))

	tables = append(tables, &data1)
	tables = append(tables, &data2)
	return tables
}
