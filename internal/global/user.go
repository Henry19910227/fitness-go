package global

type TrainerStatus int
const (
	TrainerActivity TrainerStatus = 1 //正常
	TrainerReviewing = 2 //審核中
	TrainerRevoke = 3  //停權
	TrainerDraft = 4 //編輯中
)

// UserStatus 用戶狀態(1:正常/2:違規)
type UserStatus int
const (
	UserActivity UserStatus = 1
	UserIllegal = 2
)

// UserType 用戶類型 (1:一般用戶/2:訂閱用戶)
type UserType int
const (
	NormalUserType UserType = 1
	SubscribeUserType UserType = 2
)

// Role 角色(1:用戶/2:管理員)
type Role int
const (
	UserRole Role = 1
	AdminRole = 2
)