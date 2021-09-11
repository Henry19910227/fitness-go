package validator

type CreateTrainerBody struct {
	Name     string `json:"name" binding:"required,min=1,max=20" example:"王小明"`                // 本名 (1~20字元)
	Phone    string `json:"phone" binding:"required" example:"0978820789"`                       // 手機
	Email    string `json:"email" binding:"required,email,max=255" example:"jason@gmail.com"`    // 信箱 (最大255字元)
	Address  string `json:"address" binding:"required,min=1,max=200" example:"台北市信義區松智路五段五號"`  // 地址 (最大100字元)
}

type UpdateTrainerBody struct {
	Name             *string  `json:"name" binding:"omitempty,max=20" example:"史考特"`                // 本名 (1~20字元)
	Nickname         *string  `json:"nickname" binding:"omitempty,max=20" example:"戰車老師"`                // 暱稱 (1~20字元)
	Email            *string  `json:"email" binding:"omitempty,email,max=255" example:"jason@gmail.com"`    // 信箱 (最大255字元)
	Phone            *string  `json:"phone" binding:"omitempty,startswith=09,len=10" example:"0922244123"`  // 手機
	Address          *string  `json:"address" binding:"omitempty,max=200" example:"台北市信義區松智路五段五號"`  // 地址 (最大100字元)
	Intro            *string  `json:"intro" binding:"omitempty,max=800" example:"我叫戰車老師"`                // 教練介紹 (1~400字元)
	Experience       *int     `json:"experience" binding:"omitempty,max=40" example:"5"`                      // 年資
	Motto            *string  `json:"motto" binding:"omitempty,max=200" example:"戰車老師"`                    // 座右銘 (1~100字元)
	FacebookURL      *string  `json:"facebook_url" binding:"omitempty,max=100" example:"www.facebook.com"`    // 臉書連結
	InstagramURL     *string  `json:"instagram_url" binding:"omitempty,max=100" example:"www.instagram.com"`  // instagram連結
	YoutubeURL       *string  `json:"youtube_url" binding:"omitempty,max=100" example:"www.youtube.com"`      // youtube連結
}

type TrainerAlbumPhotoIDUri struct {
	PhotoID int64 `uri:"photo_id" binding:"required" example:"1"`
}
