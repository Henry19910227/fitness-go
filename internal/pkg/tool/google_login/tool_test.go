package google_login

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTool_GetGoogleUidByAccessToken(t *testing.T) {
	accessToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjA3NGI5MjhlZGY2NWE2ZjQ3MGM3MWIwYTI0N2JkMGY3YTRjOWNjYmMiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIzNzY1MjczNDMxNjUtdHU3ZnI3ZXBxZ3FjMWMxbWYyZ2JxcHBzbjA1NmFsZjguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIzNzY1MjczNDMxNjUtdHU3ZnI3ZXBxZ3FjMWMxbWYyZ2JxcHBzbjA1NmFsZjguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTI0MzYyNDQwMTY0NzgzMDYxNTMiLCJlbWFpbCI6InRveW9rb3lvMTk5QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhdF9oYXNoIjoicno2UkU2R0EtTWQ1RDNBUnRmcTNpZyIsIm5vbmNlIjoiYV8zWkFYR3NsNWtKdFhTNXVoV3E1MF9YTFRtQWl5aXVFSkU3NXZQU1FaMCIsIm5hbWUiOiLlu5blhqDnv7AiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EtL0FGZFp1Y3BjWVpIUVFBNGczWHAwS1hTRTdfdGktTC1YRmRybEVyTWJlUEtRPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IuWGoOe_sCIsImZhbWlseV9uYW1lIjoi5buWIiwibG9jYWxlIjoiemgtVFciLCJpYXQiOjE2NTg1NDkwNjUsImV4cCI6MTY1ODU1MjY2NX0.WeaJdo4oJLeSeDUqlQAc5JoIjAL9DJwHqQu_Ri75SJJZ5-glSwzroHQ7L1WQDmgYxEYu2TSQSUOGGhMGY4BSVgvXkqcolEUd9z0Gtd5bhfkzTNdgClhF6SK6nIKgPTaOUl5AG3uVjWhY_QCYNP0KyBGdrbBSeyhIg-sLfGUKdrZC1IFw1tXbAGMmg6dU_xGNINH3Iw_2Wq4bd_RCaL6oL_U1av1JWos-s8_3_jpzUBitF2zDMUVRCzb2VVJ7y3txjxtk7zfAY1VVKJP3EokaSFhVBLOsQG2cbFRLbaUWGKaqaMf_fqpyLPq5N_ZduE3fDPQVi__NtaPRgZ1EsBxk1w"
	tool := NewTool()
	uid, err := tool.GetGoogleUidByAccessToken(accessToken)
	assert.NoError(t, err)
	assert.Equal(t, "112436244016478306153", uid)
}

func TestUnix(t *testing.T) {
	fmt.Println(time.Now().Unix())
}
