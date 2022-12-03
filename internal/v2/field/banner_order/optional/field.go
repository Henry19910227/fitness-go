package optional

type BannerIDField struct {
	BannerID *int64 `json:"banner_id,omitempty" gorm:"column:banner_id" binding:"omitempty" example:"1"` //banner id
}
type SeqField struct {
	Seq *int `json:"seq,omitempty" gorm:"column:seq" binding:"omitempty" example:"1"` //排序號
}
