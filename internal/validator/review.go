package validator

type CreateReviewForm struct {
	Score int `form:"score" binding:"required,min=1,max=5" example:"5"` // 評分 (1~5分)
	Body string `form:"body" binding:"required,min=1,max=400" example:"非常棒的課表"` // 內文 (最大兩百字元)
}
