package main

import "fmt"

func digui(n int) int {

	fmt.Println("a:", n)
	if n < 1 {
		fmt.Println("b:", n)
		return 0
	}

	fmt.Println("c:", n)
	return digui(n - 1)
}

func main() {
	digui(100)
}
