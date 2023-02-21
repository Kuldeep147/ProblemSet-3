package main

import "fmt"

func main() {
	fmt.Println("Enter a number : ")
	var h int
	fmt.Scan(&h)
	fmt.Println(countnonprimeodd(h))

}
func isprime(n int) bool {
	if n == 2 {
		return true
	} else if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countnonprimeodd(num int) int {
	var count int
	for i := 1; i < num; i++ {
		if i%2 == 1 && !isprime(i) {
			count++
		}
	}
	return count
}
