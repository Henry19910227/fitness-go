package action

type SourceOptional struct {
	Source *int `json:"source,omitempty" gorm:"column:source" example:"2"` //動作來源(1:系統動作/2:教練自創動作)
}
