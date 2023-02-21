package main

import "fmt"

func main() {
	fmt.Println("Enter size of series : ")
	var a int
	fmt.Scan(&a)
	fmt.Println("Number of even values in series is : ", fibbonacci(a))
}

func fibbonacci(num int) int {
	var (
		a1    int = 0
		a2    int = 1
		a3    int = a1 + a2
		a4    int
		count int
	)
	if num == 1 {
		fmt.Println(a1)
	} else if num == 2 {
		fmt.Println(a1, a2)
	} else if num > 2 {

		for i := 0; i < num; i++ {
			if i < 2 {
				// for 0, 1
			} else {
				a4 = a3
				if a3%2 == 0 {
					count++
				}
				a1 = a2
				a2 = a3
				a3 = a1 + a4
			}
		}
	}
	return count
}
