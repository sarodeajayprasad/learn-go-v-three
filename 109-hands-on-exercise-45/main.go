package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x, 52)

	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", x)
	fmt.Printf("cap - %v, len - %v\n", cap(x), len(x))
	fmt.Printf("cap - %v, len - %v\n", cap(y), len(y))

	x = y

	fmt.Printf("%v\n", x)

	z := []int{56, 57, 58, 59, 60}
	x = append(x, z...)
	fmt.Println(x)
}
