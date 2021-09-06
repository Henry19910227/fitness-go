package global

type TrainerStatus int
const (
	TrainerActivity TrainerStatus = 1
	TrainerReviewing = 2
	TrainerRevoke = 3
)

// UserStatus 用戶狀態(1:正常/2:違規)
type UserStatus int
const (
	UserActivity UserStatus = 1
	UserIllegal = 2
)

// Role 角色(1:用戶/2:管理員)
type Role int
const (
	UserRole Role = 1
	AdminRole = 2
)