package main

import (
	"cron-scheduler/model"
	"cron-scheduler/scheduler"
	"cron-scheduler/service"
)

func init() {
	// 初始化数据库
	model.Initdb()
}

func main() {
	// 初始化定时器，启动定时器
	scheduler.GlobalTaskManager = scheduler.NewTaskManager()
	service.ExampleSchedulerService.InitCronScheduler(scheduler.GlobalTaskManager)
	scheduler.GlobalTaskManager.Cron.Start()

	// 阻塞主线程，保持服务运行
	select {}
}
