package fcm

type Input struct {
	Message Message `json:"message"`
}
type Message struct {
	Token        string       `json:"token"`
	Notification Notification `json:"notification"`
}
type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
