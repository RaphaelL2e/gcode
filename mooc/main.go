package main

import (
	"fmt"
	. "gcode/mooc/DataStructure/millcode"
	"gcode/utils"
)

// data structure main

func main() {
	//TestPrintN()
	//TestF()
}

func TestPrintN() {
	timer := utils.NewTimer()
	timer.StaticsTime()
	PrintN(2 << 4)
	timer.StaticsTime()
	PrintN2(2 << 4)
	timer.StaticsTime()
	PrintN(2 << 8)
	timer.StaticsTime()
	PrintN2(2 << 8)
	timer.StaticsTime()
	PrintN(2 << 15)
	timer.StaticsTime()
	PrintN2(2 << 15)
	timer.StaticsTime()
}

func TestF() {
	timer := utils.NewTimer()

	MaxN := 10
	arr := make([]float64, MaxN)
	for i := 0; i < MaxN; i++ {
		arr[i] = float64(i)
	}
	timer.StaticsTime()
	f1 := F1(MaxN-1, arr, 1.1)
	fmt.Println(f1)
	timer.StaticsTime()
	f2 := F2(MaxN-1, arr, 1.1)
	fmt.Println(f2)
	timer.StaticsTime()
}
