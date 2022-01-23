package global

type CourseStatus int
const (
	Preparing CourseStatus = 1
	Reviewing CourseStatus = 2
	Sale CourseStatus = 3
	Reject CourseStatus = 4
	Remove CourseStatus = 5
)

type SaleType int
const (
	SaleTypeFree SaleType = 1
	SaleTypeSubscribe SaleType = 2
	SaleTypeCharge SaleType = 3
)
