package article_05

import (
	"math"
	"errors"
	"fmt"
)

var (
	ArrayOutOfRange = errors.New("数组越界")
	ArrayFULL       = errors.New("数组满了")
)

//定义一个数组的数据结构

type Array struct {
	data   []int
	length uint
}

//make一个数组
func New(capacity uint) *Array {
	if capacity <= 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

//根据index返回一个值
func (a *Array) Find(index uint) (int, error) {
	//确保index的值没有越界
	if a.OutOfRange(index) {
		return math.MaxInt32, ArrayOutOfRange
	}
	return a.data[index], nil
}

func (a *Array) OutOfRange(index uint) bool {
	if index < 0 || index >= a.length {
		return true
	}
	return false
}

//插入到一个数组中，是需要移动的
func (a *Array) Insert(index uint, val int) error {
	if a.length == uint(cap(a.data)) {
		return ArrayFULL
	}

	if index != a.length && a.OutOfRange(index) {
		return ArrayOutOfRange
	}

	//移动  从后面开始移动是因为会产生覆盖
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = val
	a.length++
	return nil
}

//删除index位置的数
func (a *Array) Del(index uint) error {
	if a.OutOfRange(index) {
		return ArrayOutOfRange
	}
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return nil
}

func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

func (a *Array) Len() uint {
	return a.length
}

func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|Array[%v]=[%v]\n", i, a.data[i])
	}
	fmt.Printf(format)
}
