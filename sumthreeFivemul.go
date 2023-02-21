package main

import "fmt"

func main() {
	fmt.Println("Enter a number : ")
	var k int
	fmt.Scan(&k)
	fmt.Println(sum(k))
}

func sum(n int) int {
	var sum int
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
