package cronjob

import (
	"github.com/robfig/cron/v3"
	"log"
)

func CronFunc() {

	c := cron.New(cron.WithSeconds())
	//以下测试能够发现任务1和任务3，4是可以同时运行的，任务2也是，但任务1和任务2，任务3和任务4之间是不能同时运行的
	chain1 := SkipIfStillRunningChain()
	chain2 := SkipIfStillRunningChain()

	cmd1 := func() {
		log.Println("任务1执行")
		log.Println("任务1结束")
	}
	cmd2 := func() {
		log.Println("任务2执行")
		log.Println("任务2结束")
	}
	cmd3 := func() {
		log.Println("任务3执行")
		log.Println("任务3结束")
	}
	cmd4 := func() {
		log.Println("任务4执行")
		log.Println("任务4结束")
	}

	job1 := GenerateJob(&chain1, cmd1)
	job2 := GenerateJob(&chain1, cmd2)
	job3 := GenerateJob(&chain2, cmd3)
	job4 := GenerateJob(&chain2, cmd4)

	c.AddJob("0/2 * * * * ? ", job1)
	c.AddJob("0/2 * * * * ? ", job2)
	c.AddJob("0/2 * * * * ? ", job3)
	c.AddJob("0/2 * * * * ? ", job4)
	c.Start()
	select {}
}
