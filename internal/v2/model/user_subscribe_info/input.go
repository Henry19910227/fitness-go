package user_subscribe_info

type FindInput struct {
	UserIDOptional
}

// APIGetUserSubscribeInfoInput /v2/user/subscribe_info [GET]
type APIGetUserSubscribeInfoInput struct {
	UserIDRequired
}

