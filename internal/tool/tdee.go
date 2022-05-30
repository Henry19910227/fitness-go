package tool

import "github.com/Henry19910227/fitness-go/internal/global"

type tdee struct {
	activityLevelMap map[global.ActivityLevel]float64
	exerciseFeqLevel map[global.ExerciseFeqLevel]float64
}

func NewTDEE() TDEE {
	activityLevelMap := map[global.ActivityLevel]float64{
		global.ActivityLevel1:  1.0,
		global.ActivityLevel2:  1.05,
		global.ActivityLevel3:  1.1,
		global.ActivityLevel4:  1.16,
		global.ActivityLevel5:  1.2,
		global.ActivityLevel6:  1.375,
		global.ActivityLevel7:  1.425,
		global.ActivityLevel8:  1.55,
		global.ActivityLevel9:  1.75,
		global.ActivityLevel10: 1.9,
	}
	exerciseFeqLevel := map[global.ExerciseFeqLevel]float64{
		global.ExerciseFeqLevel1: 0,
		global.ExerciseFeqLevel2: 150,
		global.ExerciseFeqLevel3: 300,
		global.ExerciseFeqLevel4: 450,
	}
	return &tdee{activityLevelMap: activityLevelMap, exerciseFeqLevel: exerciseFeqLevel}
}

func (t *tdee) TDEE(bmr int, activityLevel global.ActivityLevel, exerciseFeqLevel global.ExerciseFeqLevel) float64 {
	if _, ok := t.activityLevelMap[activityLevel]; !ok {
		return 0
	}
	if _, ok := t.exerciseFeqLevel[exerciseFeqLevel]; !ok {
		return 0
	}
	return float64(bmr)*t.activityLevelMap[activityLevel] + t.exerciseFeqLevel[exerciseFeqLevel]
}
