package main

import "fmt"

func main() {
	var a [5]int
    fmt.Println("emp:", a)

	a[4] = 100

	fmt.Println("set value  index 4", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
}