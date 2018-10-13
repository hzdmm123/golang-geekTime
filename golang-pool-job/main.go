package job

import (
	"golang-geekTime/golang-pool"
	"fmt"
)

var poolOne worker.WorkPool

func init() {
	poolOne.InitPool()
	worker.StartPool(6)


}

var JobCReturn chan int

func Tain() {
	//调用协程池进行处理
	worker.Dispatch(Run, 1, 2)
	worker.Dispatch(Run, 3, 4)
	worker.Dispatch(Run, 1, 8)
	worker.Dispatch(RunA, 3, 4)
	//获取协程池结果
	fmt.Println("======\n")
	fmt.Println(<-worker.WorkTaskReturn)
	fmt.Println("======\n")
	worker.StopPool()

	/*	JobCReturn = make(chan int, 3)
		poolOne.Run(Run, 5, 6)
		//var runcReturn worker.ReturnType
		//利用map 传递地址的特性 来拿回结果
		var resultChan = make(chan interface{}, 200)
		fmt.Println(resultChan)
		for i := 0; i < 2000; i++ {
			var paramMap = make(map[string]interface{})
			paramMap["a"] = 7 + i
			paramMap["b"] = 8
			poolOne.Run(RunC, paramMap)
			//runcReturn =<-resultChan
			//fmt.Println(runcReturn.(int))
		}
		//<-resultChan
		poolOne.Stop()*/
}
