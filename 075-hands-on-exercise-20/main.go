package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(400)
	fmt.Printf("The name of the variable is x and its values is %v\n", x)
	/*
		if x <= 100 {
			fmt.Println("less than 100")
		} else if x >= 101 && x <= 200 {
			fmt.Println("between 101 and 200")
		} else if x >= 201 && x <= 250 {
			fmt.Println("between 201 and 250")
		} else {
			fmt.Println("this was more than 250")
		}
	*/
	switch {
	case x <= 100:
		fmt.Println("less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("betwee 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("this was more than 250")
	}
}
