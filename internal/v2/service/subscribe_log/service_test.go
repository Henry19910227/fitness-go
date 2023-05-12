package subscribe_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	subscribeLogModel "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_log"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_CreateOrUpdate(t *testing.T) {
	subscribeLogService := NewService(orm.Shared().DB())

	subscribeLogTable := subscribeLogModel.Table{}
	subscribeLogTable.OriginalTransactionID = util.PointerString("2000000")
	subscribeLogTable.TransactionID = util.PointerString("2000001")
	subscribeLogTable.PurchaseDate = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	subscribeLogTable.ExpiresDate = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	subscribeLogTable.Type = util.PointerString("renew")
	subscribeLogTable.Msg = util.PointerString("DID_RENEW ")
	_, err := subscribeLogService.CreateOrUpdate(&subscribeLogTable)
	assert.NoError(t, err)
}
