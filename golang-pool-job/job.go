package job

import (
	"fmt"
	"golang-geekTime/golang-pool"
)

func jobTest(param []int) {
	fmt.Printf("this is Job test!\n")
	fmt.Printf("this is test params[0]=[%v]  params[1]=[%v]\n", param[0], param[1])
	var returnParam []interface{}
	returnParam = append(returnParam, param[0]+param[1])
	returnParam = append(returnParam, param[0]*param[1])
	worker.WorkTaskReturn <- returnParam
}

func Run(param []interface{}) {
	var paramJob []int
	for _, p := range param {
		switch v := p.(type) {
		case int:
			var s int
			s = v
			paramJob = append(paramJob, s)
		}
	}
	jobTest(paramJob)
}
