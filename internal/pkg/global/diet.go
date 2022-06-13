package global

type ActivityLevel int

const (
	ActivityLevel1  ActivityLevel = 1  // 麻痺、昏迷、無法活動
	ActivityLevel2  ActivityLevel = 2  // 臥床不動，僅手臂移動
	ActivityLevel3  ActivityLevel = 3  // 幾乎坐著或躺著
	ActivityLevel4  ActivityLevel = 4  // 大部分坐著，少許步行
	ActivityLevel5  ActivityLevel = 5  // 久坐、辦公室性質工作
	ActivityLevel6  ActivityLevel = 6  // 每週輕度步行3-4天
	ActivityLevel7  ActivityLevel = 7  // 每週輕度步行5-7天
	ActivityLevel8  ActivityLevel = 8  // 體力勞動工作性質
	ActivityLevel9  ActivityLevel = 9  // 沉重的體力勞動工作性質
	ActivityLevel10 ActivityLevel = 10 // 極重度的勞動或職業運動員
)

type ExerciseFeqLevel int

const (
	ExerciseFeqLevel1 ExerciseFeqLevel = 1 // 無運動
	ExerciseFeqLevel2 ExerciseFeqLevel = 2 // 一週2-3次，一次30-45分鐘
	ExerciseFeqLevel3 ExerciseFeqLevel = 3 // 一週3-5次，一次45-60分鐘
	ExerciseFeqLevel4 ExerciseFeqLevel = 4 // 一週5次以上，一次60分鐘
)

type DietTarget int

const (
	DietTargetLoseFat     DietTarget = 1 // 減脂
	DietTargetBuildMuscle DietTarget = 2 // 增肌
	DietTargetKeep        DietTarget = 3 // 維持健康生活
	DietTargetPowerUp     DietTarget = 4 // 提升體能與力量
	DietTargetFeed        DietTarget = 5 // 哺乳者
	DietTargetPregnant    DietTarget = 6 // 懷孕者
)

type DietType int

const (
	DietTypeGeneral       DietType = 1 // 標準飲食
	DietTypeAllVegan      DietType = 2 // 全素食
	DietTypeOvoLactoVegan DietType = 3 // 蛋奶素食
	DietTypeOvoVegan      DietType = 4 // 蛋素食
	DietTypeLactoVegan    DietType = 5 // 奶素食
)

type FoodCategoryTag int

const (
	GrainTag     FoodCategoryTag = 1 // 穀物類
	MeatTag      FoodCategoryTag = 2 // 蛋豆魚肉類
	FruitTag     FoodCategoryTag = 3 // 水果類
	VegetableTag FoodCategoryTag = 4 // 蔬菜類
	DairyTag     FoodCategoryTag = 5 // 乳製品類
	NutTag       FoodCategoryTag = 6 // 堅果類
)
