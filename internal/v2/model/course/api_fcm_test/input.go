package api_fcm_test

// Input /v2/fcm_test [POST]
type Input struct {
	Body Body
}
type Body struct {
	Payload interface{} `json:"payload" binding:"required"`
}
