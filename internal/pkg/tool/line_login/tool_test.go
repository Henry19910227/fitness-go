package line_login

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTool_GetClientIDByAccessToken(t *testing.T) {
	accessToken := "eyJhbGciOiJIUzI1NiJ9.nulxNAqmNAsu-mjO4ixndWDoOzCv4Rm12yg9M6o0up4DIik3xLfKQKp2SOUr3CtofE5iufzduKY9XosMgCxR_RC4rA_7g0gnbwWVr5DGom-5T4xvUtyAwvNRkhC5mVOtBYAs_sBGhAquWGRiTuwI-xiLz8ZptiVDAKy2tVrK7ow.m3AUNqIu1TMdN48pYnlqZJUaggGUJFM9ZUdiA4POwag"
	tool := NewTool()
	uid, err := tool.GetUserID(accessToken)
	assert.NoError(t, err)
	assert.Equal(t, "U4c93579de5ed302ef1fd0421ae4456e7", uid)
}
