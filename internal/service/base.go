package service

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
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
