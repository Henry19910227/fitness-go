package fb_login

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTool(t *testing.T) {
	tool := NewTool()
	fbUid, err := tool.GetUserIDByAccessToken("EAAucgU8qZCzMBAOZCy59TLD1aM2NAO1ITBpZC64imFp95CRuPv4ZAWepAMVISqmFseTq7WnAIA1VBnwdhzBUm29dxsiqlCGH76s2WFuYqzWUp76nGeEMp6CL3ahUFhpZCArfPKs4cQa6v2RUcjOPLHxbLE6tx6ZCdvwLYYNiGSD3bEN68KviPYsGWApv1DzUeNSLke9ZA4cpEvmgSBZALWroNV9iHD06p8GDzbo1dMlZA5QZDZD")
	assert.NoError(t, err)
	assert.Equal(t, "2316609181824364", fbUid)
}
