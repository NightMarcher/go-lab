package main

import (
    "fmt"
    "time"
    "go-lab/pkg"
)


// import
func demoImport() {
    hello_string := pkg.Hello("Bryan")
    fmt.Println(hello_string)
}

// array
func demoArray() {
    list := []int{2, 3, 5, 7}
    fmt.Printf("\n%T => %+v\n", list, list)
    for index, value := range list {
        fmt.Println(index, ":", value)
    }
    fmt.Println(append(list, 11, 13))
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

// time
func demoTime() {
    now := time.Now()
    fmt.Printf("\n%T => %+v\n", now, now)
}

func sum(args ...int) int {
    total := 0
    for _, num := range args {
        total += num
    }
    fmt.Printf("\nsum(%v) = %d\n", args, total)
    return total
}

func fib(num int) []int {
    list := []int{}
    if num <= 0 {
        return list
    }
    last, this := 0, 1
    for i := 0; i < num; i++ {
        last, this = this, last + this
        list = append(list, this)
    }
    fmt.Printf("\nfib(%d) = %+v\n", num, list)
    return list
}

func main() {
    demoImport()

    demoArray()

    demoString()

    demoMap()

    demoTime()

    sum(1, 2, 3)

    fib(10)
}
