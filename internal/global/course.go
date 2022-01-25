package global

type CourseStatus int

const (
	Preparing CourseStatus = 1
	Reviewing CourseStatus = 2
	Sale      CourseStatus = 3
	Reject    CourseStatus = 4
	Remove    CourseStatus = 5
)

type SaleType int

const (
	SaleTypeNone      SaleType = 0 // 未指定
	SaleTypeFree      SaleType = 1 // 免費型課表
	SaleTypeSubscribe SaleType = 2 // 訂閱型課表
	SaleTypeCharge    SaleType = 3 // 付費型課表
)

type ScheduleType int

const (
	SingleScheduleType   ScheduleType = 1 // 單一訓練
	MultipleScheduleType ScheduleType = 2 // 多項計劃
)
