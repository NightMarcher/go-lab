package basic

import (
	"fmt"
	"math"
)

// string
func DemoString() {
	str := "ABCabc"
	fmt.Printf("\n%T => %#v\n", str, str)
	for index, char := range str {
		fmt.Println(index, ":", char)
	}
	fmt.Println(len(str))
	fmt.Println(str[:3])
	fmt.Println(str[len(str)-3:])
}

// array
func DemoArray() {
	array := [3]byte{'A', 'B', 'C'}
	fmt.Printf("\n%T => %c\n", array, array)
	for index, value := range array {
		fmt.Println(index, ":", value)
	}
}

// slice
func DemoSlice() {
	slice := []int{2, 3, 5, 7}
	fmt.Printf("\n%T => %+v\n", slice, slice)
	for index, value := range slice {
		fmt.Println(index, ":", value)
	}
	fmt.Println(append(slice, 11, 13))
}

// map
func DemoMap() {
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

	delete(dict, "A")
	fmt.Println(dict)

	toUpdate := map[string]int{
		"D": 13,
		"E": 14,
		"F": 15,
	}
	for key, value := range toUpdate {
		dict[key] = value
	}
	fmt.Println(dict)
}

type rect struct {
	name          string
	width, height float32
}

// struct
func DemoStruct() {
	fmt.Printf("\n%+v\n", rect{"rectangle", 2, 3})
	fmt.Println(rect{name: "square", width: 4, height: 4})

	mine := rect{name: "myRect"}
	pointerMine := &mine
	fmt.Println(pointerMine)
	fmt.Println(pointerMine.width, pointerMine.height)

	pointerMine.width, pointerMine.height = 5, 6
	fmt.Println(pointerMine.width, pointerMine.height)
}

func (r rect) area() float32 {
	return r.width * r.height
}

type circle struct {
	name   string
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type geometry interface {
	area() float32
}

func measure(g geometry) {
	fmt.Printf("\nmeasuring %+v => area = %f\n", g, g.area())
}

// interface
func DemoInterface() {
	myRect := rect{"myRect", 3, 4}
	myCircle := circle{"myCircle", 5}
	measure(myRect)
	measure(myCircle)
}

func handleValue(val int) {
	val = 0
}

func handlePointer(ptr *int) {
	*ptr = 0
}

// pointer
func DemoPointer() {
	value := 10
	fmt.Println("\nvalue:", value)
	handleValue(value)
	fmt.Println("after handleValue:", value)

	fmt.Println("pointer:", &value)
	handlePointer(&value)
	fmt.Println("after handlePointer:", value)
}
