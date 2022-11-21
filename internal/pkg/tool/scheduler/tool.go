package scheduler

import (
	"github.com/robfig/cron/v3"
)

type tool struct {
	cron *cron.Cron
}

func New() Tool {
	return &tool{cron: cron.New(cron.WithSeconds())}
}

func (t *tool) Cron() *cron.Cron {
	return t.cron
}