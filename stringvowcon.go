package main

import (
	"fmt"
)

func main() {
	var u string
	fmt.Println("Enter a string : ")
	fmt.Scanf("%s", &u)
	countvowcons(u)

}

func countvowcons(str string) {
	var vowels = [10]rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	var consonants = [42]rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z', 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z'}
	var (
		count    int
		countcon int
	)
	for _, value := range str {
		for i := 0; i < len(vowels); i++ {
			if value == vowels[i] {
				count++
			}
		}

	}
	for _, val := range str {
		for i := 0; i < len(consonants); i++ {
			if val == consonants[i] {
				countcon++
			}
		}
	}
	fmt.Println("Number of vowels is ", count, " Number of consonant is ", countcon)
}
