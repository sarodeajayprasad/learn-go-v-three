package main

import "fmt"

func main() {
	x := [5]int{}

	for i := 0; i < 5; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}

}
