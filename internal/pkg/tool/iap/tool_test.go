package iap

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTool_GenerateAppleStoreAPIToken(t *testing.T) {
	tool := NewTool()
	token, err := tool.GenerateAppleStoreAPIToken(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(token) > 0)
}

func TestTool_VerifyAppleReceiptAPI(t *testing.T) {
	tool := NewTool()
	receiptData := "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"
	response, err := tool.VerifyAppleReceiptAPI(receiptData)
	assert.NoError(t, err)
	assert.Equal(t, 0, response.Status)
}

func TestTool_GetSubscribeAPI(t *testing.T) {
	tool := NewTool()
	//產出token
	token, err := tool.GenerateAppleStoreAPIToken(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(token) > 0)
	//以 original_transactionId 取得訂閱資訊
	response, err := tool.GetSubscribeAPI("2000000100820920", token)
	assert.NoError(t, err)
	assert.Equal(t, true, len(response.Data) > 0)
}
