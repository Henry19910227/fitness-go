package fcm

type Output struct {
	Message struct{
		Token string `json:"token"`
		Notification struct{
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"notification"`
		Data struct{
			Title string `json:"title"`
			Message  string `json:"message"`
		} `json:"data"`
	} `json:"message"`
}
