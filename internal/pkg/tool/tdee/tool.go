package tdee

const (
	ActivityLevel1  = 1  // 麻痺、昏迷、無法活動
	ActivityLevel2  = 2  // 臥床不動，僅手臂移動
	ActivityLevel3  = 3  // 幾乎坐著或躺著
	ActivityLevel4  = 4  // 大部分坐著，少許步行
	ActivityLevel5  = 5  // 久坐、辦公室性質工作
	ActivityLevel6  = 6  // 每週輕度步行3-4天
	ActivityLevel7  = 7  // 每週輕度步行5-7天
	ActivityLevel8  = 8  // 體力勞動工作性質
	ActivityLevel9  = 9  // 沉重的體力勞動工作性質
	ActivityLevel10 = 10 // 極重度的勞動或職業運動員
)

const (
	ExerciseFeqLevel1 = 1 // 無運動
	ExerciseFeqLevel2 = 2 // 一週2-3次，一次30-45分鐘
	ExerciseFeqLevel3 = 3 // 一週3-5次，一次45-60分鐘
	ExerciseFeqLevel4 = 4 // 一週5次以上，一次60分鐘
)

type tool struct {
	activityLevelMap map[int]float64
	exerciseFeqLevel map[int]float64
}

func New() Tool {
	activityLevelMap := map[int]float64{
		ActivityLevel1:  1.0,
		ActivityLevel2:  1.05,
		ActivityLevel3:  1.1,
		ActivityLevel4:  1.16,
		ActivityLevel5:  1.2,
		ActivityLevel6:  1.375,
		ActivityLevel7:  1.425,
		ActivityLevel8:  1.55,
		ActivityLevel9:  1.75,
		ActivityLevel10: 1.9,
	}
	exerciseFeqLevel := map[int]float64{
		ExerciseFeqLevel1: 0,
		ExerciseFeqLevel2: 150,
		ExerciseFeqLevel3: 300,
		ExerciseFeqLevel4: 450,
	}
	return &tool{activityLevelMap: activityLevelMap, exerciseFeqLevel: exerciseFeqLevel}
}

func (t *tool) TDEE(bmr int, activityLevel int, exerciseFeqLevel int) float64 {
	if _, ok := t.activityLevelMap[activityLevel]; !ok {
		return 0
	}
	if _, ok := t.exerciseFeqLevel[exerciseFeqLevel]; !ok {
		return 0
	}
	return float64(bmr)*t.activityLevelMap[activityLevel] + t.exerciseFeqLevel[exerciseFeqLevel]
}
