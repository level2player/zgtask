package zgtask

import (
	"github.com/benbjohnson/clock"
	"log"
	"runtime/debug"
	"time"
)

var (
	TaskContainer map[string]ITask
)

func InjectionTask(id string, task ITask) {
	if TaskContainer == nil {
		TaskContainer = map[string]ITask{}
	}
	if _, ok := TaskContainer[id]; ok {
		panic("id repeat,please check")
	}
	TaskContainer[id] = task
}

func Run() {

	for key, v := range TaskContainer {
		if v != nil {
			config := v.InitConfig()
			if config == nil {
				log.Println("task init error,", key)
				continue
			}
			if isFixedTask(config) {
				fixedTaskConfig := config.(*FixedTaskConfig)
				go func(k string, task ITask) {

					defer func() {
						if err := recover(); err != nil {
							log.Println(err)
							log.Println(string(debug.Stack()))
						}
					}()
					scheduler := task.GetScheduler()
					if fixedTaskConfig.isStartImmediately() {
						task.Execute()
					}
					tick := clock.New()
					for t := range tick.Tick(time.Second) {
						select {
						case <-scheduler:
							log.Println(k, "safe exit")
							return
						default:
						}
						for _, v := range fixedTaskConfig.getTaskExecuteTime() {
							if t.Format("15:04:05") == v {
								task.Execute()
							}
						}
					}
				}(key, v)
			}

			if isTimerTask(config) {
				timerConfig := config.(*TimerConfig)
				go func(k string, task ITask) {
					defer func() {
						if err := recover(); err != nil {
							log.Println(err)
							log.Println(string(debug.Stack()))
						}
					}()
					scheduler := task.GetScheduler()
					if timerConfig.isStartImmediately() {
						task.Execute()
					}
					tick := clock.New()
					for range tick.Tick(timerConfig.getTaskExecuteIntervalTime()) {
						select {
						case <-scheduler:
							log.Println(k, "safe exit")
							return
						default:
						}
						task.Execute()
					}
				}(key, v)
			}
		}
	}
}

func Stop() {
	for _, task := range TaskContainer {
		c := task.GetScheduler()
		c <- true
	}
}
