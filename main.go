package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("Soma 5 + 5 = %v\n", sum(5, 5))
}
