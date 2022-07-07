package review

type IDRequired struct {
	ID int64 `json:"id" uri:"review_id" binding:"required" example:"1"` //評論id
}
