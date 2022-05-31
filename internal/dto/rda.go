package dto

type RDA struct {
	TDEE      int `json:"tdee" example:"2533"`    // TDEE
	Calorie   int `json:"calorie" example:"2913"` // 目標熱量
	Protein   int `json:"protein" example:"146"`  // 蛋白質(克)
	Fat       int `json:"fat" example:"65"`       // 脂肪(克)
	Carbs     int `json:"carbs" example:"437"`    // 碳水化合物(克)
	Grain     int `json:"grain" example:"25"`     // 穀物類(份)
	Vegetable int `json:"vegetable" example:"5"`  // 蔬菜類(份)
	Fruit     int `json:"fruit" example:"2"`      // 水果類(份)
	Meat      int `json:"meat" example:"12"`      // 蛋豆魚肉類(份)
	Dairy     int `json:"dairy" example:"1"`      // 乳製品類(份)
	Nut       int `json:"nut" example:"5"`        // 堅果類(份)
}

type CalculateRDAParam struct {
	Sex              string  // 性別 (f:女/m:男)
	Birthday         string  // 生日
	Height           float64 // 身高 (最大999)
	Weight           float64 // 體重 (最大999)
	BodyFat          *int    // 體脂肪率 (最大99)
	ActivityLevel    int     // 活動量 (1:麻痺、昏迷、無法活動/2:臥床不動，僅手臂移動/3:幾乎坐著或躺著/4:大部分坐著，少許步行/5:久坐、辦公室性質工作/6:每週輕度步行3-4天/7:每週輕度步行5-7天/8:體力勞動工作性質/9:沉重的體力勞動工作性質/10:極重度的勞動或職業運動員)
	ExerciseFeqLevel int     // 運動頻率 (1:無運動/2:一週2-3次，一次30-45分鐘/3:一週3-5次，一次45-60分鐘/4:一週5次以上，一次60分鐘)
	DietTarget       int     // 飲食目標 (1:減脂/2:增肌/3:維持健康生活/4:提升體能與力量/5:哺乳者/6:懷孕者)
	DietType         int     // 飲食型態 (1:標準飲食/2:全素食/3:蛋奶素食/4:蛋素食/5:奶素食)
}
