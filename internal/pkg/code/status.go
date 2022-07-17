package code

const (
	Success int = 0 // Success Request
	BadRequest int = 9000 // Bad Request
	DataAlreadyExists int = 9003 // 資料已存在
	InvalidToken int = 9005 // JWT rejected
	PermissionDenied int = 9006 // Permission denied
)
