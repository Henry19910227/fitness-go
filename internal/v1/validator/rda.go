package validator

type CalculateRDABody struct {
	Sex              string  `json:"sex" binding:"required,oneof=m f" example:"m"`                             // 性別 (f:女/m:男)
	Birthday         string  `json:"birthday" binding:"required,datetime=2006-01-02" example:"1992-02-02"`     // 生日
	Height           float64 `json:"height" binding:"required,max=999" example:"178"`                          // 身高 (最大999)
	Weight           float64 `json:"weight" binding:"required,max=999" example:"70"`                           // 體重 (最大999)
	BodyFat          *int    `json:"body_fat" binding:"omitempty,max=99" example:"20"`                         // 體脂肪率 (最大99)
	ActivityLevel    int     `json:"activity_level" binding:"required,oneof=1 2 3 4 5 6 7 8 9 10" example:"6"` // 活動量 (1:麻痺、昏迷、無法活動/2:臥床不動，僅手臂移動/3:幾乎坐著或躺著/4:大部分坐著，少許步行/5:久坐、辦公室性質工作/6:每週輕度步行3-4天/7:每週輕度步行5-7天/8:體力勞動工作性質/9:沉重的體力勞動工作性質/10:極重度的勞動或職業運動員)
	ExerciseFeqLevel int     `json:"exercise_feq_level" binding:"required,oneof=1 2 3 4" example:"3"`          // 運動頻率 (1:無運動/2:一週2-3次，一次30-45分鐘/3:一週3-5次，一次45-60分鐘/4:一週5次以上，一次60分鐘)
	DietTarget       int     `json:"diet_target" binding:"required,oneof=1 2 3 4 5 6" example:"2"`             // 飲食目標 (1:減脂/2:增肌/3:維持健康生活/4:提升體能與力量/5:哺乳者/6:懷孕者)
	DietType         int     `json:"diet_type" binding:"required,oneof=1 2 3 4 5 " example:"1"`                // 飲食型態 (1:標準飲食/2:全素食/3:蛋奶素食/4:蛋素食/5:奶素食)
}

type UpdateRDABody struct {
	TDEE      int `json:"tdee" binding:"required" example:"2000"`    // TDEE
	Calorie   int `json:"calorie" binding:"required" example:"1800"` // 目標熱量
	Protein   int `json:"protein" binding:"required" example:"70"`   // 蛋白質(克)
	Fat       int `json:"fat" binding:"required" example:"20"`       // 脂肪(克)
	Carbs     int `json:"carbs" binding:"required" example:"50"`     // 碳水化合物(克)
	Grain     int `json:"grain" binding:"required" example:"3"`      // 穀物類(份)
	Vegetable int `json:"vegetable" binding:"required" example:"5"`  // 蔬菜類(份)
	Fruit     int `json:"fruit" binding:"required" example:"2"`      // 水果類(份)
	Meat      int `json:"meat" binding:"required" example:"6"`       // 蛋豆魚肉類(份)
	Dairy     int `json:"dairy" binding:"omitempty" example:"3"`      // 乳製品類(份)
	Nut       int `json:"nut" binding:"required" example:"1"`        // 堅果類(份)
}
