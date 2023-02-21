package main

import "fmt"

func main() {
	fmt.Println("Enter number a : ")
	var (
		first  int
		second int
	)
	fmt.Scan(&first)
	fmt.Println("Enter nuber b : ")
	fmt.Scan(&second)
	fmt.Println(sum(first, second))
	fmt.Println(diff(first, second))
	fmt.Println(getsum_n_diff(first, second))
}
func sum(a int, b int) int {
	return a + b
}
func diff(a int, b int) int {
	return a - b
}
func getsum_n_diff(a int, b int) (int, int) {
	return (a + b), (a - b)
}
