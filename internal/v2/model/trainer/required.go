package trainer

type UserIDRequired struct {
	UserID int64 `json:"user_id,omitempty" uri:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` // 用戶id
}
type NameRequired struct {
	Name *string `json:"name,omitempty" gorm:"column:name" example:"亨利"` // 教練本名
}
type NicknameRequired struct {
	Nickname *string `json:"nickname,omitempty" gorm:"column:nickname" example:"Henry"` // 教練暱稱
}
type SkillRequired struct {
	Skill *string `json:"skill,omitempty" gorm:"column:skill" example:"1,4"` // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
}
type AvatarRequired struct {
	Avatar *string `json:"avatar,omitempty" gorm:"column:avatar" example:"abc.png"` // 教練大頭照
}
type TrainerStatusRequired struct {
	TrainerStatus *int `json:"trainer_status,omitempty" gorm:"column:trainer_status" example:"1"` // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
}
type TrainerLevelRequired struct {
	TrainerLevel *int `json:"trainer_level,omitempty" gorm:"column:trainer_level" example:"1"` // 教練評鑑等級
}
type EmailRequired struct {
	Email *string `json:"email,omitempty" gorm:"column:email" example:"henry@gmail.com"` // 信箱
}
type PhoneRequired struct {
	Phone *string `json:"phone,omitempty" gorm:"column:phone" example:"0955123321"` // 電話
}
type AddressRequired struct {
	Address *string `json:"address,omitempty" gorm:"column:address" example:"台北市信義區五段五號"` // 住址
}
type IntroRequired struct {
	Intro *string `json:"intro,omitempty" gorm:"column:intro" example:"Henry教練"` // 個人介紹
}
type ExperienceRequired struct {
	Experience *int `json:"experience,omitempty" gorm:"column:experience" example:"5"` // 年資
}
type MottoRequired struct {
	Motto *string `json:"motto,omitempty" gorm:"column:motto" example:"勞其筋骨"` // 座右銘
}
type FacebookURLRequired struct {
	FacebookURL *string `json:"facebook_url,omitempty" gorm:"column:facebook_url" example:"www.facebook.com"` // 臉書連結
}
type InstagramURLRequired struct {
	InstagramURL *string `json:"instagram_url,omitempty" gorm:"column:instagram_url" example:"www.ig.com"` // ig連結
}
type YoutubeURLRequired struct {
	YoutubeURL *string `json:"youtube_url,omitempty" gorm:"column:youtube_url" example:"www.youtube.com"` // youtube連結
}
type CreateAtRequired struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtRequired struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-12 00:00:00"` // 更新時間
}
