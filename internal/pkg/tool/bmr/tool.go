package bmr

type tool struct{}

func New() Tool {
	return &tool{}
}

func (t *tool) MSJBMR(weight float64, height float64, age int, sex string) float64 {
	baseBMR := weight*10 + height*6.25 - float64(age)*5
	if sex == "m" {
		return baseBMR + 5
	}
	return baseBMR - 161
}

func (t *tool) KMABMR(weight float64, bodyFat int) float64 {
	return 370 + 21.6*(weight*(100-float64(bodyFat))/100)
}
