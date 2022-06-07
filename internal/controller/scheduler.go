package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	courseUsageStatisticService            service.CourseUsageStatistic
	userCourseUsageMonthlyStatisticService service.UserCourseUsageMonthlyStatistic
	userIncomeMonthlyStatisticService      service.UserIncomeMonthlyStatistic
}

func NewScheduler(schedulerTool *cron.Cron,
	courseUsageStatisticService service.CourseUsageStatistic,
	userCourseUsageMonthlyStatisticService service.UserCourseUsageMonthlyStatistic,
	userIncomeMonthlyStatisticService service.UserIncomeMonthlyStatistic) {
	scheduler := Scheduler{
		courseUsageStatisticService:            courseUsageStatisticService,
		userCourseUsageMonthlyStatisticService: userCourseUsageMonthlyStatisticService,
		userIncomeMonthlyStatisticService:      userIncomeMonthlyStatisticService,
	}
	// 參考 https://segmentfault.com/a/1190000039647260
	_, _ = schedulerTool.AddFunc("0 * * * * *", scheduler.MinuteTask) // 每分鐘零秒執行任務
}

func (s *Scheduler) MinuteTask() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " 執行每分鐘輪詢任務")
	//s.courseUsageStatisticService.Update()
	//s.userCourseUsageMonthlyStatisticService.Update()
	//s.userIncomeMonthlyStatisticService.Update()
}
