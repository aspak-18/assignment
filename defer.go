package main

import (
	"fmt"
)

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("Cannot divide by zero!")
	}

	return a / b
}

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))
}
