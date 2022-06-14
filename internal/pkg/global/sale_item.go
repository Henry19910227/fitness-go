package global

type PeriodType int
const (
	OneMonthPeriodType   PeriodType = 1  // 一個月
	TwoMonthPeriodType   PeriodType = 2  // 二個月
	ThreeMonthPeriodType PeriodType = 3  // 三個月
	SixMonthPeriodType   PeriodType = 6  // 六個月
	OneYearPeriodType    PeriodType = 12 // 一年
	ForeverPeriodType    PeriodType = 99 // 永久週期
)
