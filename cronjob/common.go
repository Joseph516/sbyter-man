package cronjob

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

// SkipIfStillRunningChain 产生一个chain，在这条chain上的任务只要有一个在新的周期未执行完就不会开启新的一轮执行
func SkipIfStillRunningChain() cron.Chain {
	return cron.NewChain(cron.SkipIfStillRunning(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
}

// DelayIfStillRunningChain 产生一个chain，在这条chain上的任务只要有一个在新的周期未执行完会被推迟执行
func DelayIfStillRunningChain() cron.Chain {
	return cron.NewChain(cron.DelayIfStillRunning(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
}

func GenerateJob(chain *cron.Chain, cmd func()) cron.Job {
	return chain.Then(cron.FuncJob(cmd))
}
