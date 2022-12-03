package optional

type BannerIDField struct {
	BannerID *int64 `json:"banner_id,omitempty" uri:"banner_id" form:"banner_id" binding:"omitempty" example:"1"` //id
}
