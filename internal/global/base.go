package global

type OrderType string
const (
	ASC OrderType = "ASC"
	DESC OrderType = "DESC"
)

type OrderTypeField string
const (
	FieldUpdateAt OrderTypeField = "update_at"
)
