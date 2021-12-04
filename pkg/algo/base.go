package algo

import "fmt"

func Sum(nums ...float32) float32 {
	var total float32 = 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("\nSum(%v) = %f\n", nums, total)
	return total
}

func Fib(num int) []int {
	list := []int{}
	if num <= 0 {
		fmt.Printf("\nFib(%d) = %v\n", num, list)
		return list
	}
	last, this := 0, 1
	for i := 0; i < num; i++ {
		list = append(list, this)
		last, this = this, last+this
	}
	fmt.Printf("\nFib(%d) = %v\n", num, list)
	return list
}
