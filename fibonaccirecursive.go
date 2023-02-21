package main

import "fmt"

func main() {
	fmt.Println("Enter the size of series : ")
	var size uint
	fmt.Scan(&size)
	printfib(size)
	fmt.Println()

}

func fibbonacci(num uint) uint {
	if num == 1 {
		return 0
	} else if num == 2 {
		return 1
	} else {
		return fibbonacci(num-2) + fibbonacci(num-1)
	}

}

func printfib(n uint) {
	var i uint
	for i = 1; i <= n; i++ {
		fmt.Printf("%d, ", fibbonacci(i))
	}
}
