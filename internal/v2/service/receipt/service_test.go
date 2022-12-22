package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	"testing"
)

func TestService_CreateOrUpdate(t *testing.T) {
	svc := NewService(orm.Shared().DB())
	table := receipt.Table{}
	table.OrderID = util.PointerString("20221125112602870326")
	table.PaymentType = util.PointerInt(1)
	table.ReceiptToken = util.PointerString("ABBA")
	table.OriginalTransactionID = util.PointerString("GPA.000")
	table.TransactionID = util.PointerString("GPA.000")
	table.ProductID = util.PointerString("com.fitness.silver_course")
	table.Quantity = util.PointerInt(1)
	table.CreateAt = util.PointerString("2022-11-25 11:29:52")
	_, err := svc.CreateOrUpdate(&table)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
