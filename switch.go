package main

import ( 
	"fmt"
	"time" 
)

func main() {

	i := 3
    fmt.Print("Write ", i, " as ")
    switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

	t := time.Now()

	switch {

	case t.Hour() < 12:
		fmt.Println("It's Before noon")
	default:
		fmt.Println("Its After noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a Bool")
		case int:
			fmt.Println("I'm a Int")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}

	whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}