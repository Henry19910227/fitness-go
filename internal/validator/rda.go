package validator

type CalculateRDABody struct {
	Sex string `json:"sex" binding:"required,oneof=m f" example:"m"`                         // 性別 (f:女/m:男)
	Birthday string `json:"birthday" binding:"required,datetime=2006-01-02" example:"1991-02-27"` // 生日
	Height   float64 `json:"height" binding:"required,max=999" example:"176.5"`                    // 身高 (最大999)
	Weight   float64 `json:"weight" binding:"required,max=999" example:"70.5"`                     // 體重 (最大999)
	BodyFat  *int `json:"body_fat" binding:"omitempty,max=99" example:"15"`                     // 體脂肪率 (最大99)
	ActivityLevel int `json:"activity_level" binding:"required,oneof=1 2 3 4 5 6 7 8 9 10" example:"6"`  // 活動量 (1:麻痺、昏迷、無法活動/2:臥床不動，僅手臂移動/3:幾乎坐著或躺著/4:大部分坐著，少許步行/5:久坐、辦公室性質工作/6:每週輕度步行3-4天/7:每週輕度步行5-7天/8:體力勞動工作性質/9:沉重的體力勞動工作性質/10:極重度的勞動或職業運動員)
	ExerciseFeq int `json:"exercise_feq" binding:"required,oneof=1 2 3 4" example:"3"`  // 運動頻率 (1:無運動/2:一週2-3次，一次30-45分鐘/3:一週3-5次，一次45-60分鐘/4:一週5次以上，一次60分鐘)
	Target int `json:"target" binding:"required,oneof=1 2 3 4 5 6" example:"4"` // 目標 (1:減脂/2:增肌/3:維持健康生活/4:提升體能與力量/5:哺乳者/6:懷孕者)
	DietType int `json:"diet_type" binding:"required,oneof=1 2 3 4 5 " example:"2"`  // 飲食型態 (1:標準飲食/2:全素食/3:蛋奶素食/4:蛋素食/5:奶素食)
}
