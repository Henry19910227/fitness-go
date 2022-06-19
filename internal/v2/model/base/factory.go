package base

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
)

func SetRequest(status int, msg string) *Output {
	switch status {
	case code.Success:
		return Success()
	case code.InvalidToken:
		return InvalidToken()
	case code.PermissionDenied:
		return PermissionDenied()
	case code.BadRequest:
		return BadRequest(util.PointerString(msg))
	default:
		return BadRequest(nil)
	}
}

func BadRequest(msg *string) *Output {
	output := Output{}
	output.Set(code.BadRequest, "bad request")
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

func Success() *Output {
	output := Output{}
	output.Code = code.Success
	output.Msg = "success"
	return &output
}
