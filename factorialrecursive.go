package main

import "fmt"

func main() {
	fmt.Println("Enter a number to compute its factorial : ")
	var n int
	fmt.Scan(&n)
	a := factorial(n)
	fmt.Println(a)
}
func factorial(a int) int {
	if a == 0 || a == 1 {
		return a
	}
	return a * factorial(a-1)
}
