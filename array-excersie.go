package main

import "fmt"

func main() {

	b := [5]int{ 1, 2, 3, 4, 5 }

	var reversedArray [5]int


	last_index := len(b) - 1

	for i := 0; i < len(b); i ++ {
		reversedArray[last_index - i] = b[i]
	}


	fmt.Println("This is the reversed Array", reversedArray)
}