package main

import "fmt"

func main() {
	//sum numbers in a slice using range
	//here we ignore the index with _
	nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
	fmt.Println("sum:", sum)
	//without ignoring index

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index: ", i)
		}
	}

	//range iterates over key value pairs on maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

	//iterating over only keys of a map
	for k := range kvs {
		fmt.Println("key", k)
	}

	//range on strings iterates over unicode code points
	for i, c := range "go" {
        fmt.Println(i, c)
    }

}