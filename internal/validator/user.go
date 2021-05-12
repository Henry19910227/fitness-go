package validator

type UpdateMyUserInfoBody struct {
	Email       *string `json:"email" binding:"omitempty,email" example:"henry@gmail.com"`             // 信箱
	Nickname    *string `json:"nickname" binding:"omitempty,min=1,max=16" example:"henry"`             // 暱稱 (1~16字元)
	Sex         *string `json:"sex" binding:"omitempty,oneof=m f" example:"m"`                         // 性別 (f:女/m:男)
	Birthday    *string `json:"birthday" binding:"omitempty,datetime=2006-01-02" example:"1991-02-27"` // 生日
	Height      *string `json:"height" binding:"omitempty,max=230" example:"176.5"`                    // 身高 (最大230)
	Weight      *string `json:"weight" binding:"omitempty,max=230" example:"70.5"`                     // 體重 (最大230)
	Experience  *string `json:"experience" binding:"omitempty,max=4" example:"2"`                      // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
	Target      *string `json:"target" binding:"omitempty,max=3" example:"3"`                          // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
}