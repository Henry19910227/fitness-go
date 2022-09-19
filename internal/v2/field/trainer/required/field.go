package required

type UserIDField struct {
	UserID int64 `json:"user_id" uri:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type NameField struct {
	Name string `json:"name" gorm:"column:name" binding:"required" example:"亨利"` // 教練本名
}
type NicknameField struct {
	Nickname string `json:"nickname" gorm:"column:nickname" binding:"required" example:"Henry"` // 教練暱稱
}
type SkillField struct {
	Skill string `json:"skill" gorm:"column:skill" binding:"required" example:"1,4"` // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
}
type AvatarField struct {
	Avatar string `json:"avatar" gorm:"column:avatar" binding:"required" example:"abc.png"` // 教練大頭照
}
type TrainerStatusField struct {
	TrainerStatus int `json:"trainer_status" gorm:"column:trainer_status" binding:"required" example:"1"` // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
}
type TrainerLevelField struct {
	TrainerLevel int `json:"trainer_level" gorm:"column:trainer_level" binding:"required" example:"1"` // 教練評鑑等級
}
type EmailField struct {
	Email string `json:"email" gorm:"column:email" binding:"required" example:"henry@gmail.com"` // 信箱
}
type PhoneField struct {
	Phone string `json:"phone" gorm:"column:phone" binding:"required" example:"0955123321"` // 電話
}
type AddressField struct {
	Address string `json:"address" gorm:"column:address" binding:"required" example:"台北市信義區五段五號"` // 住址
}
type IntroField struct {
	Intro string `json:"intro" gorm:"column:intro" binding:"required" example:"Henry教練"` // 個人介紹
}
type ExperienceField struct {
	Experience int `json:"experience" gorm:"column:experience" binding:"required" example:"5"` // 年資
}
type MottoField struct {
	Motto string `json:"motto" gorm:"column:motto" binding:"required" example:"勞其筋骨"` // 座右銘
}
type FacebookURLField struct {
	FacebookURL string `json:"facebook_url" gorm:"column:facebook_url" binding:"required" example:"www.facebook.com"` // 臉書連結
}
type InstagramURLField struct {
	InstagramURL string `json:"instagram_url" gorm:"column:instagram_url" binding:"required" example:"www.ig.com"` // ig連結
}
type YoutubeURLField struct {
	YoutubeURL string `json:"youtube_url" gorm:"column:youtube_url" binding:"required" example:"www.youtube.com"` // youtube連結
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-12 00:00:00"` // 更新時間
}
