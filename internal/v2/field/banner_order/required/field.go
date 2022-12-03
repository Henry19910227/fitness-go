package required

type BannerIDField struct {
	BannerID int64 `json:"banner_id" gorm:"column:banner_id" binding:"required" example:"1"` //banner id
}
type SeqField struct {
	Seq int `json:"seq" gorm:"column:seq" binding:"required" example:"1"` //排序號
}
