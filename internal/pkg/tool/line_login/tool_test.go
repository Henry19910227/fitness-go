package line_login

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTool_GetClientIDByAccessToken(t *testing.T) {
	accessToken := "eyJhbGciOiJIUzI1NiJ9.R6OpXX6cEtCl983HtZ7UlYf9zQU1yIosbVJNVY3rPfWM2GmDMskqEyyyX3YMe9arzb1VGOe1wQsSC6V6Y2V-Dt-8g2yLY3DKupvdkd82YUNmbF1_1LsxPh37ZipngO03Modt5YBpl9GiU2sUn0sQ2HNNdq32a1lZZcvF5edkoBM.XPIk9tR2wqSeqbXCsSscIB8c_dNqHQGoomelXs9Mrn8"
	tool := NewTool()
	uid, err := tool.GetUserID(accessToken, "")
	assert.NoError(t, err)
	assert.Equal(t, "1657326779", uid)
}
