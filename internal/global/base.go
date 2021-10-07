package global

type OrderType string
const (
	ASC OrderType = "ASC"
	DESC OrderType = "DESC"
)

type OrderTypeField string
const (
	UpdateAt OrderTypeField = "update_at"
)
