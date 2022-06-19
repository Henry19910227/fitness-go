package base

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
)

type Output struct {
	CodeField
	MsgField
}

func (o *Output) SetStatus(status int) {
	switch status {
	case code.Success:
		o.Msg = "success"
	case code.InvalidToken:
		o.Msg = "JWT rejected"
	case code.PermissionDenied:
		o.Msg = "bad Permission denied"
	case code.BadRequest:
		o.Msg = "bad request"
	default:
		o.Msg = "bad request"
	}
	o.Code = status
}

func (o *Output) Set(status int, msg string) {
	o.Code = status
	o.Msg = msg
}
