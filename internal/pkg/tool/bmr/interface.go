package bmr

type Tool interface {
	// MSJBMR Mifflin-St Jeor BMR 方程式
	MSJBMR(weight float64, height float64, age int, sex string) float64
	// KMABMR Katch-McArdle BMR 方程式
	KMABMR(weight float64, bodyFat int) float64
}
