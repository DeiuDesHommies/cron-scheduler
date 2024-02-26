package service

import (
	"cron-scheduler/model"
	"cron-scheduler/scheduler"
	"fmt"
)

type ExampleScheduler struct{}

func (bs *ExampleScheduler) InitCronScheduler(taskManager *scheduler.TaskManager) {
	fmt.Println("Start init scheduler.")
	schedules, err := model.ScheduleDB.GetSchedules()
	if err != nil {
		fmt.Println("Failed to get schedules.")
	}

	// 将定时计划任务注册到调度器
	for _, schedule := range schedules {
		executor, err := GetExecutor(schedule)
		if err != nil {
			fmt.Println("Failed to get executor type.")
		}
		if executor != nil {
			err := taskManager.AddJob(schedule.Cron, executor, schedule.ID)
			if err != nil {
				fmt.Println("Failed to add schedule.")
			}
		}
	}
	fmt.Println("finished init scheduler.")
}

func GetExecutor(schedule *model.Schedule) (scheduler.Executor, error) {
	if schedule.Type == "task" {
		return &Task{
			TaskID: schedule.TaskID,
			Type:   schedule.Type,
		}, nil
	}
	var err error
	err = fmt.Errorf("Task type error")
	return nil, err
}
