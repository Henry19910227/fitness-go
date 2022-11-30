package optional

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" uri:"user_id" gorm:"column:user_id" binding:"omitempty" example:"10001"` // 用戶id
}
type NameField struct {
	Name *string `json:"name,omitempty" form:"name" gorm:"column:name" binding:"omitempty,max=20" example:"亨利"` // 教練本名
}
type NicknameField struct {
	Nickname *string `json:"nickname,omitempty" form:"nickname" gorm:"column:nickname" binding:"omitempty,max=20" example:"Henry"` // 教練暱稱
}
type SkillField struct {
	Skill *string `json:"skill,omitempty" form:"skill" gorm:"column:skill" binding:"omitempty,trainer_skill" example:"1,4"` // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
}
type AvatarField struct {
	Avatar *string `json:"avatar,omitempty" form:"avatar" gorm:"column:avatar" binding:"omitempty" example:"abc.png"` // 教練大頭照
}
type TrainerStatusField struct {
	TrainerStatus *int `json:"trainer_status,omitempty" form:"trainer_status" gorm:"column:trainer_status" binding:"omitempty,oneof=1 2 3" example:"1"` // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
}
type TrainerLevelField struct {
	TrainerLevel *int `json:"trainer_level,omitempty" form:"trainer_level" gorm:"column:trainer_level" binding:"omitempty,min=1,max=5" example:"1"` // 教練評鑑等級
}
type EmailField struct {
	Email *string `json:"email,omitempty" form:"email" gorm:"column:email" binding:"omitempty,email,max=255" example:"henry@gmail.com"` // 信箱
}
type PhoneField struct {
	Phone *string `json:"phone,omitempty" form:"phone" gorm:"column:phone" binding:"omitempty,startswith=09,len=10" example:"0955123321"` // 電話
}
type AddressField struct {
	Address *string `json:"address,omitempty" form:"address" gorm:"column:address" binding:"omitempty,max=200" example:"台北市信義區五段五號"` // 住址
}
type IntroField struct {
	Intro *string `json:"intro,omitempty" form:"intro" gorm:"column:intro" binding:"omitempty,max=800" example:"Henry教練"` // 個人介紹
}
type ExperienceField struct {
	Experience *int `json:"experience,omitempty" form:"experience" gorm:"column:experience" binding:"omitempty,max=40" example:"5"` // 年資
}
type MottoField struct {
	Motto *string `json:"motto,omitempty" form:"motto" gorm:"column:motto" binding:"omitempty,max=200" example:"勞其筋骨"` // 座右銘
}
type FacebookURLField struct {
	FacebookURL *string `json:"facebook_url,omitempty" form:"facebook_url" gorm:"column:facebook_url" binding:"omitempty,max=100" example:"www.facebook.com"` // 臉書連結
}
type InstagramURLField struct {
	InstagramURL *string `json:"instagram_url,omitempty" form:"instagram_url" gorm:"column:instagram_url" binding:"omitempty,max=100" example:"www.ig.com"` // ig連結
}
type YoutubeURLField struct {
	YoutubeURL *string `json:"youtube_url,omitempty" form:"youtube_url" gorm:"column:youtube_url" binding:"omitempty,max=100" example:"www.youtube.com"` // youtube連結
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 更新時間
}
