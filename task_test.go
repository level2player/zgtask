package task_controller

import (
	"fmt"
	"testing"
	"time"
)

type TimedTask struct {
	Task
}

func (*TimedTask) Execute() {
	fmt.Println(time.Now().Format("15:04:05"))
}

func (*TimedTask) InitConfig() interface{} {
	f := TimerConfig{}
	f.StartImmediately = false
	f.TaskExecuteIntervalTime = time.Second * 2
	return &f
}

func Test_TimedTask(t *testing.T) {
	InjectionTask("TimedTask", &TimedTask{})
	Run()
	time.Sleep(time.Second * 10)
	Stop()
	time.Sleep(time.Minute * 5)
}

type FixedTask struct {
	Task
}

func (*FixedTask) Execute() {
	fmt.Println(time.Now().Format("15:04:05"))
}

func (*FixedTask) InitConfig() interface{} {
	f := FixedTaskConfig{}
	f.StartImmediately = false
	f.TaskExecuteTime = []string{"14:15:00", "14:15:05"}
	return &f
}

func Test_FixedTask(t *testing.T) {
	InjectionTask("FixedTask", &FixedTask{})
	Run()
	time.Sleep(time.Minute)
}

func Test_Tasks(t *testing.T) {
	InjectionTask("TimedTask", &TimedTask{})
	InjectionTask("FixedTask", &FixedTask{})
	Run()
	time.Sleep(time.Second * 10)
	Stop()
	time.Sleep(time.Second * 10)
}
