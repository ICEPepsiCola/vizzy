package bootstrap

import (
	"github.com/robfig/cron"
	"sync"
)

var (
	cronOnce     sync.Once
	cronInstance *cron.Cron
)

func initCron(registerFn func(cron *cron.Cron)) {
	cronOnce.Do(func() {
		cronInstance = cron.New()
	})
	registerFn(cronInstance)
	cronInstance.Start()
}
