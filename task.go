package task_controller

import (
	"log"
	"time"
)

type ITask interface {
	Execute()
	InitConfig() interface{}
	GetScheduler() chan interface{}
}

type Config struct {
	//Whether to perform a task immediately after startup
	StartImmediately bool
}

type TimerConfig struct {
	Config
	TaskExecuteIntervalTime time.Duration
}

type FixedTaskConfig struct {
	Config
	TaskExecuteTime []string
}

func defaultConfig(interval time.Duration) interface{} {

	t := TimerConfig{
		TaskExecuteIntervalTime: interval,
	}
	t.StartImmediately = true

	return &t
}

func isFixedTask(dest interface{}) bool {
	_, b := dest.(*FixedTaskConfig)
	return b
}
func isTimerTask(dest interface{}) bool {
	_, b := dest.(*TimerConfig)
	return b
}

func (t *FixedTaskConfig) isStartImmediately() bool {
	return t.StartImmediately
}

func (t *FixedTaskConfig) getTaskExecuteTime() []string {
	return t.TaskExecuteTime
}

func (t *TimerConfig) isStartImmediately() bool {
	return t.StartImmediately
}

func (t *TimerConfig) getTaskExecuteIntervalTime() time.Duration {
	return t.TaskExecuteIntervalTime
}

type Task struct {
	scheduler chan interface{}
}

func (t *Task) GetScheduler() chan interface{} {
	if t.scheduler == nil {
		t.scheduler = make(chan interface{})
	}
	return t.scheduler
}

func (*Task) Execute() {
	log.Println("Not implementation execute")
}

func (*Task) InitConfig() interface{} {
	log.Println("The task type is not configured, so the default configuration will be started.")
	return defaultConfig(time.Second)
}
