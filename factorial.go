package main

import "fmt"

func main() {
	fmt.Println("Enter a number to compute factorial : ")
	var k int
	fmt.Scan(&k)
	fmt.Println("factorial is : ", factorial(k))
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	var a int = 1
	for i := 1; i <= num; i++ {
		a *= i
	}
	return a
}
