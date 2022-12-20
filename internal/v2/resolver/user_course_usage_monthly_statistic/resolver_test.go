package user_course_usage_monthly_statistic

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	timeValue, _ := time.Parse("2006-01-02 15:04:05", "2022-12-06 12:00:00")
	fmt.Println(timeValue.Format("2006"))
	fmt.Println(timeValue.Format("1"))
}
