package main

import (
	"go-lab/pkg/algo"
	"go-lab/pkg/basic"
)

func main() {
	basic.DemoString()

	basic.DemoArray()

	basic.DemoSlice()

	basic.DemoMap()

	basic.DemoStruct()

	basic.DemoInterface()

	basic.DemoPointer()

	basic.DemoError()

	basic.DemoTime()

	basic.DemoTimer()

	basic.DemoTicker()

	basic.DemoPool()

	basic.DemoWaitGroup()

	algo.Sum(1, 2, 3)
	array := [5]float32{4, 5, 6}
	algo.Sum(array[:]...)

	algo.Fib(10)
}
