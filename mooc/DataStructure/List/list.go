package main

import (
	"fmt"
	"reflect"
)

// List  线性表数据结构
type List[T any] []T

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Append(v T) {
	*l = append(*l, v)
}

func (l *List[T]) Len() int {
	return len(*l)
}

func (l *List[T]) Get(index int) T {
	return (*l)[index]
}

func (l *List[T]) Set(index int, value T) {
	if index < 0 || index >= l.Len() {
		panic("Index out of range")
	}

	// 修改指定下标处的元素
	(*l)[index] = value
}

func (l *List[T]) IndexOf(value T) (index int) {
	for index, a := range *l {
		if reflect.DeepEqual(a, value) {
			return index
		}
	}
	return -1
}

func (l *List[T]) Insert(index int, value T) {
	if index < 0 || index >= l.Len() {
		panic("Index out of range")
	}

	l.Append(value)
	copy((*l)[index+1:], (*l)[index:])
	(*l)[index] = value
}

func (l *List[T]) RemoveAt(index int) {
	if index < 0 || index >= l.Len() {
		panic("Index out of range")
	}

	*l = append((*l)[:index], (*l)[index+1:]...)
}

func (l *List[T]) ListAll() {
	for i, _ := range *l {
		fmt.Print(l.Get(i))
		fmt.Print(" ")
	}
	fmt.Println()
}

func main() {
	l := NewList[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	fmt.Println(l.Len())

	fmt.Println(l.Get(1))

	fmt.Println(l.IndexOf(3))

	fmt.Println(l.Get(0))

	l.ListAll()

	l.Insert(0, 0)

	l.ListAll()

	l.Set(0, 4)

	l.ListAll()

	l.RemoveAt(0)

	l.ListAll()

	l2 := NewList[float64]()
	l2.Append(1.1)
	l2.Append(2.1)
	l2.Append(3.1)

	fmt.Println(l2.Get(1))

	fmt.Println(l2.IndexOf(3))

	l2.Insert(2, 4.5)

	l2.ListAll()

	l2.RemoveAt(0)

	l2.ListAll()
}
