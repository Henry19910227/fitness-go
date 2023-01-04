package admin

import "github.com/Henry19910227/fitness-go/internal/v2/field/admin/optional"

type Table struct {
	optional.IDField
	optional.EmailField
	optional.PasswordField
	optional.NicknameField
	optional.LvField
	optional.CreateAtField
	optional.UpdateAtField
	optional.LastLoginField
	optional.StatusField
}

func (Table) TableName() string {
	return "admins"
}
