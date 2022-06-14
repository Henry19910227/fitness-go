package base

import "github.com/Henry19910227/fitness-go/internal/pkg/code"

func BadRequest(msg *string) *Output {
	output := Output{}
	output.Code = code.BadRequest
	output.Msg = "bad request"
	if msg != nil {
		output.Msg = *msg
	}
	return &output
}

func InvalidToken() *Output {
	output := Output{}
	output.Code = code.InvalidToken
	output.Msg = "invalid token"
	return &output
}

func PermissionDenied() *Output {
	output := Output{}
	output.Code = code.PermissionDenied
	output.Msg = "permission denied"
	return &output
}
