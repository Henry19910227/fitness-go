package tdee

type Tool interface {
	TDEE(bmr int, activityLevel int, exerciseFeqLevel int) float64
}
