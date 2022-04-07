package main

import "fmt"

func main() {
	//create a Map and assign it key value pairs
	m := make(map[string]int)

	m["k1"] = 7
    m["k2"] = 13
	fmt.Println("map:", m)
	//Get a value for a key with a name
	v1 := m["k1"]
    fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))
	//remove a key value pair from a map
	delete(m, "k2")
    fmt.Println("map:", m)
	//check for a key whether its missing 
	//or has zero value(0 or "") 
	//using optional second return
	_, prs := m["k1"] 
    fmt.Println("prs:", prs)

	//declare and initialize a new map in the sameline
	n := map[string]int{"foor":1, "bar":2}
    fmt.Println("map:", n)

}