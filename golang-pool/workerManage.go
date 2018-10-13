package worker

import (
	"time"
	"fmt"
)

type Job func([]interface{})

const (
	workerNumMax     = 10000
	workerNumDefault = 20
)

type taskWork struct {
	startBool bool //
	params    []interface{}
	Run       Job
}

type WorkPool struct {
	taskPool   chan taskWork
	workNum    int
	maxNum     int
	defaultNum int
	stopTopic  bool
	taskQueue  chan taskWork
}

func (p *WorkPool) InitPool() {
	*p = WorkPool{
		defaultNum: workerNumDefault,
		maxNum:     workerNumMax,
		stopTopic:  false,
		taskPool:   make(chan taskWork, workerNumDefault*2),
		taskQueue:  nil,
	}
	p.start()
}
func (p *WorkPool) start() {
	for i := 0; i < p.defaultNum; i++ {
		p.workInit(i)
	}
}
func (p *WorkPool) workInit(id int) {
	p.workNum++
	f := func(idNum int) {
		for {
			select {
			case task := <-p.taskPool:
				if task.startBool == true && task.Run != nil {
					task.Run(task.params)
				}

				if !task.startBool {
					return
				}
			case <-time.After(time.Second * 1000):
				if p.stopTopic == true && len(p.taskPool) == 0 {
					fmt.Printf("topic=%+v", p.stopTopic)
					p.workNum--
					return
				}
			case queueTask := <-p.taskQueue:
				if queueTask.startBool == true && queueTask.Run != nil {
					queueTask.Run(queueTask.params)
				}
			case <-func() chan int {
				a := make(chan int)
				fmt.Printf(`run ....`)
				return a
			}():
				fmt.Printf(`.....`)
			}

		}

	}
	go f(id)
}

func (p *WorkPool) Stop() {
	p.stopTopic = false
}

func (p *WorkPool) Run(funcJob Job, params ...interface{}) {
	p.taskQueue <- taskWork{true, params, funcJob}
}

func (p *WorkPool) RunAuto(funcJob Job, params ...interface{}) {
	task := taskWork{true, params, funcJob}
	select {
	case p.taskPool <- task:
	case <-time.After(time.Millisecond * 1000):
		p.taskQueueInit()
		p.worderAddConf()
		p.taskQueue <- task
	}
}
func (p *WorkPool) taskQueueInit() {
	if p.taskQueue == nil {
		p.taskQueue = make(chan taskWork, p.maxNum*2)
	}
}
func (p *WorkPool) worderAddConf() {
	if p.workNum < 1000 {
		p.workerAdd(p.workNum)

	} else if p.workNum < p.maxNum {
		tmpNum := p.maxNum - p.workNum
		tmpNum = tmpNum / 10
		if tmpNum == 0 {
			tmpNum = 1
		}
		p.workerAdd(tmpNum)
	}
}

func (p *WorkPool) workerRemoveConf() {
	for {
		select {
		case <-time.After(time.Millisecond * 1000 * 600):
			if p.workNum > p.defaultNum && len(p.taskPool) == 0 && len(p.taskQueue) == 0 {
				rmNum := (p.workNum - p.defaultNum) / 5
				if rmNum == 0 {
					rmNum = 1
				}
				p.workerRemove(rmNum)
			}
		}
	}
}
func (p *WorkPool) workerAdd(num int) {
	for i := 0; i < num; i++ {
		p.workNum++
		p.workInit(p.workNum)
	}
}
func (p *WorkPool) workerRemove(num int) {
	for i := 0; i < num; i++ {
		task := taskWork{startBool: false}
		p.taskPool <- task
		p.workNum--
	}
}
