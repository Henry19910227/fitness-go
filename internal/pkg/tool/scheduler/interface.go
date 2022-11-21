package scheduler

import (
	"github.com/robfig/cron/v3"
)

type Tool interface {
	Cron() *cron.Cron
}
