package validator

type UpdateUserInfoBody struct {
	//Email       *string `json:"email" binding:"omitempty,email" example:"henry@gmail.com"`             // 信箱
	//Nickname    *string `json:"nickname" binding:"omitempty,min=1,max=16" example:"henry"`             // 暱稱 (1~16字元)
	Sex         *string `json:"sex" binding:"omitempty,oneof=m f" example:"m"`                         // 性別 (f:女/m:男)
	Birthday    *string `json:"birthday" binding:"omitempty,datetime=2006-01-02" example:"1991-02-27"` // 生日
	Height      *float64 `json:"height" binding:"omitempty,max=230" example:"176.5"`                    // 身高 (最大230)
	Weight      *float64 `json:"weight" binding:"omitempty,max=230" example:"70.5"`                     // 體重 (最大230)
	Experience  *int     `json:"experience" binding:"omitempty,max=4" example:"2"`                      // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
	Target      *int     `json:"target" binding:"omitempty,max=3" example:"3"`                          // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
}

type CreateTrainerBody struct {
	Name     string `json:"name" binding:"required,min=1,max=16" example:"王小明"`                // 本名 (1~16字元)
	Nickname string `json:"nickname" binding:"required,min=1,max=16" example:"jason"`            // 暱稱 (1~16字元)
	Phone    string `json:"phone" binding:"required" example:"0978820789"`                       // 聯絡電話
	Email    string `json:"email" binding:"required,email" example:"jason@gmail.com"`            // 信箱 (最大230)
}