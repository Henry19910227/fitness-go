package controller

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/robfig/cron/v3"
	"time"
)

type Scheduler struct {
	courseUsageStatisticService service.CourseUsageStatistic
}

func NewScheduler(schedulerTool *cron.Cron, courseUsageStatisticService service.CourseUsageStatistic)  {
	scheduler := Scheduler{
		courseUsageStatisticService: courseUsageStatisticService,
	}
	_, _ = schedulerTool.AddFunc("0 * * * * *", scheduler.MinuteTask) // 每分鐘零秒執行任務
}

func (s *Scheduler) MinuteTask() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " 執行每分鐘輪詢任務")
	s.courseUsageStatisticService.UpdateCourseUsageStatistic()
}