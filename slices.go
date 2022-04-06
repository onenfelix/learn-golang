package main

import "fmt"

func main() {
	//create a slice with a length of 3
	s := make([]string, 3)
    fmt.Println("emp:", s)
	//set and get values from a slice
	s[0] = "felix"
    s[1] = "onen"
    s[2] = "cleared"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
	//print length of the slice
	fmt.Println("len:", len(s))
	//append to a slice
	s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

	//we can also copy a slice
	c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
	//slices also support the slice operator
	l := s[2:5]
    fmt.Println("sl1:", l)
	l = s[:5]
    fmt.Println("sl2:", l)
	l = s[2:]
    fmt.Println("sl3:", l)
	//declare and initialize a slice variable
	t := []string{"g", "h", "i"}
	t = append(t,"onen")
    fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)


}