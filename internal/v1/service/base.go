package service

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"math"
)

// Base ...
type Base struct {

}

// MysqlAccessDenied ...
func (b *Base) MysqlAccessDenied(err error) bool {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
			return true
		}
		return false
	}
	return false
}

// MysqlDuplicateEntry ...
func (b *Base) MysqlDuplicateEntry(err error) bool {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
			return true
		}
		return false
	}
	return false
}

func (b *Base) GetPagingIndex(page int, size int) (int, int) {
	offset := (page - 1) * size
	return offset, size
}

func (b *Base) GetTotalPage(totalCount int, size int) int {
	totalPage := int(math.Ceil(float64(totalCount)/float64(size)))
	if totalPage < 0 {
		return 0
	}
	return totalPage
}
