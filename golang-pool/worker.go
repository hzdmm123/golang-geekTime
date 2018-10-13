package worker

import (
	"fmt"
	"time"
)

var WorkMaxTask int
var WorkTaskPool chan taskWork
var WorkTaskReturn chan []interface{}

//run task
func (t *taskWork) start() {
	go func() {
		fmt.Println("启动线程...")
		for {
			select {
			case funcRun := <-WorkTaskPool:
				if funcRun.startBool {
					funcRun.Run(funcRun.params)
				} else {
					fmt.Printf("task stop!")
					return
				}

			case <-time.After(time.Millisecond * 1000):
				fmt.Printf("time out")
			}

		}
	}()
}

//stop task
func (t *taskWork) stop() {
	fmt.Printf("task stop")
	t.startBool = false
}

//启动协程池
func StartPool(maxTask int) {
	WorkMaxTask = maxTask
	WorkTaskPool = make(chan taskWork, maxTask)
	WorkTaskReturn = make(chan []interface{}, maxTask) //任务的返回值

	for i := 0; i < maxTask; i++ {
		var t = createTask()
		fmt.Printf("start task 第%d个任务\n", i)
		t.start()
	}
}

func createTask() taskWork {
	var funcJob Job
	var paramsSlice []interface{}
	return taskWork{true, paramsSlice, funcJob}
}

func StopPool() {
	var funcJob Job
	var paramSlice []interface{}

	for i := 0; i < WorkMaxTask; i++ {
		WorkTaskPool <- taskWork{false, paramSlice, funcJob}
	}
}

func Dispatch(funcJob Job, params ...interface{}) {
	WorkTaskPool <- taskWork{true, params, funcJob}
}
