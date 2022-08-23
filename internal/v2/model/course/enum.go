package course

// CourseStatus Enum
const (
	Preparing int = 1 //準備中
	Reviewing int = 2 //審核中
	Sale      int = 3 //銷售中
	Reject    int = 4 //退審
	Remove    int = 5 //下架
)

// SaleType Enum
const (
	SaleTypeNone      int = 0 // 未指定
	SaleTypeFree      int = 1 // 免費型課表
	SaleTypeSubscribe int = 2 // 訂閱型課表
	SaleTypeCharge    int = 3 // 付費型課表
	SaleTypePersonal  int = 4 // 個人課表
)

// ScheduleType Enum
const (
	SingleWorkout int = 1 // 單一訓練
	MultiplePlan  int = 2 // 多項計畫
)
