package main

import (
	"fmt"
	"go-lab/pkg"
	"time"
)

// import
func demoImport() {
	hello_world := pkg.Hello("Bryan")
	fmt.Println(hello_world)
}

// array
func demoArray() {
	array := [3]byte{'A', 'B', 'C'}
	fmt.Printf("\n%T => %c\n", array, array)
	for index, value := range array {
		fmt.Println(index, ":", value)
	}
}

// slice
func demoSlice() {
	slice := []int{2, 3, 5, 7}
	fmt.Printf("\n%T => %+v\n", slice, slice)
	for index, value := range slice {
		fmt.Println(index, ":", value)
	}
	fmt.Println(append(slice, 11, 13))
}

// string
func demoString() {
	str := "ABCabc"
	fmt.Printf("\n%T => %#v\n", str, str)
	for index, char := range str {
		fmt.Println(index, ":", char)
	}
}

// map
func demoMap() {
	dict := map[string]int{
		"A": 10,
		"B": 11,
		"C": 12,
	}
	fmt.Printf("\n%T => %+v\n", dict, dict)
	for k, v := range dict {
		fmt.Println(k, ":", v)
	}
	value, exists := dict["D"]
	fmt.Println(value, exists)
}

func handleValue(val int) {
	val = 0
}

func handlePointer(ptr *int) {
	*ptr = 0
}

// pointer
func demoPointer() {
	value := 10
	fmt.Println("\nvalue:", value)
	handleValue(value)
	fmt.Println("after handleValue:", value)

	fmt.Println("pointer:", &value)
	handlePointer(&value)
	fmt.Println("after handlePointer:", value)
}

// time
func demoTime() {
	now := time.Now()
	fmt.Printf("\n%T => %+v\n", now, now)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("\nsum(%v) = %d\n", nums, total)
	return total
}

func fib(num int) []int {
	list := []int{}
	if num <= 0 {
		return list
	}
	last, this := 0, 1
	for i := 0; i < num; i++ {
		last, this = this, last+this
		list = append(list, this)
	}
	fmt.Printf("\nfib(%d) = %+v\n", num, list)
	return list
}

func main() {
	demoImport()

	demoArray()

	demoSlice()

	demoString()

	demoMap()

	demoPointer()

	demoTime()

	sum(1, 2, 3)

	array := [5]int{4, 5, 6}
	sum(array[:]...)

	fib(10)
}
