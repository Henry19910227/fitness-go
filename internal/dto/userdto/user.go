package userdto

type Update struct {
	AccountType *int    `gorm:"column:account_type"`
	Account     *string `gorm:"column:account"`
	Password    *string `gorm:"column:password"`
	DeviceToken *string `gorm:"column:device_token"`
	UserStatus  *string `gorm:"column:user_status"`
	UserType    *string `gorm:"column:user_type"`
	CreateAt    *string `gorm:"column:create_at"`
	UpdateAt    *string `gorm:"column:update_at"`
	Email       *string `gorm:"column:email"`
	Nickname    *string `gorm:"column:nickname"`
	Sex         *string `gorm:"column:sex"`
	Birthday    *string `gorm:"column:birthday"`
	Height      *string `gorm:"column:height"`
	Weight      *string `gorm:"column:weight"`
	Experience  *string `gorm:"column:experience"`
	Target      *string `gorm:"column:target"`
}