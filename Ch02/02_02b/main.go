package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	n := 25
	fmt.Println(str1, str2, str3)
	strLen, err := fmt.Println("Eu sou o ", n)
	if err != nil {
		fmt.Println("Error dred")
	}
	fmt.Println("String Length:", strLen)
	fmt.Printf("O NUMERO SUPREMO É O %v\n", n)
}
