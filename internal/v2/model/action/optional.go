package action

type IDOptional struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //動作id
}
type SourceOptional struct {
	Source *int `json:"source,omitempty" gorm:"column:source" example:"2"` //動作來源(1:系統動作/2:教練自創動作)
}
type VideoOptional struct {
	Video *string `json:"video,omitempty" example:"1234.mp4"` //影片名
}
