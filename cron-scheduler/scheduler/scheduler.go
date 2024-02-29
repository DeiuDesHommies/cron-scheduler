package scheduler

import (
	"github.com/robfig/cron/v3"
	"time"
)

type TaskManager struct {
	Cron *cron.Cron
	Jobs map[int]cron.EntryID
}

var GlobalTaskManager *TaskManager

// NewTaskManager 新建调度器
func NewTaskManager() *TaskManager {
	timeZone, _ := time.LoadLocation("Asia/Shanghai")
	return &TaskManager{
		Cron: cron.New(cron.WithSeconds(), cron.WithLocation(timeZone)),
		Jobs: make(map[int]cron.EntryID),
	}
}

func (tm *TaskManager) AddJob(cronExp string, executor Executor, scheduleID int) error {
	id, err := tm.Cron.AddFunc(cronExp, func() {
		executor.Execute()
	})
	if err != nil {
		return err
	}
	tm.Jobs[scheduleID] = id
	return nil
}

func (tm *TaskManager) RemoveJob(scheduleID int) {
	if jobId, exists := tm.Jobs[scheduleID]; exists {
		tm.Cron.Remove(jobId)
		delete(tm.Jobs, scheduleID)
	}
}

// UpdateJob 更新定时策略逻辑，更新了数据库内容后要更新Job
func (tm *TaskManager) UpdateJob(cronExp string, executor Executor, scheduleID int) error {
	// 移除旧任务
	tm.RemoveJob(scheduleID)

	// 添加新任务
	if err := tm.AddJob(cronExp, executor, scheduleID); err != nil {
		return err
	}
	return nil
}
