package global

type CourseStatus int
const (
	Preparing CourseStatus = 1
	Reviewing = 2
	Sale = 3
	Reject = 4
	Remove = 5
)

type SaleType int
const (
	SaleTypeFree SaleType = 1
	SaleTypeVIP SaleType = 2
	SaleTypeCharge SaleType = 3
)
