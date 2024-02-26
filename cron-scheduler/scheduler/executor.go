package scheduler

// Executor 用来执行具体的定时任务，各服务自己实现
type Executor interface {
	Execute()
}
