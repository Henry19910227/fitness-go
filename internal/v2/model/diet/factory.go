package diet

import "github.com/Henry19910227/fitness-go/internal/pkg/util"

func NewMockTables() []*Table {
	tables := make([]*Table, 0)
	data1 := Table{}
	data1.ID = util.PointerInt64(1)
	data1.UserID = util.PointerInt64(10001)
	return tables
}
