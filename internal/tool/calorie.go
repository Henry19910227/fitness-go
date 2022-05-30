package tool

import "github.com/Henry19910227/fitness-go/internal/global"

type calorie struct {
	dietTargetMap map[global.DietTarget]float64
}

func NewCalorie() Calorie {
	dietTargetMap := map[global.DietTarget]float64{
		global.DietTargetLoseFat:     0.85,
		global.DietTargetBuildMuscle: 1.15,
		global.DietTargetKeep:        1,
		global.DietTargetPowerUp:     1.15,
	}
	return &calorie{dietTargetMap: dietTargetMap}
}

func (c calorie) TargetCalorie(tdee int, target global.DietTarget) float64 {
	if target == global.DietTargetFeed {
		return float64(tdee) + 600
	}
	if target == global.DietTargetPregnant {
		return float64(tdee) + 300
	}
	return float64(tdee) * c.dietTargetMap[target]
}
