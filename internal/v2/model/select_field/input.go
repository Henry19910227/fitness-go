package select_field

type Select struct {
	Query string // query 必須塞一個欄位才會是正確結果
	Args  []interface{}
}

type Input struct {
	Selects []*Select
}
