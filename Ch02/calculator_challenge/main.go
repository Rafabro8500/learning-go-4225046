package main

import (
	"fmt"
	"strconv"
)

func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	val1, err := strconv.ParseFloat(value1, 64)
	if err != nil {
		fmt.Printf("couldn't convert string %s to float64 %w", value2, err)
		return 0
	}

	// Convert the second string to a float64
	val2, err := strconv.ParseFloat(value2, 64)
	if err != nil {
		fmt.Printf("couldn't convert string %s to float64 %w", value2, err)
		return 0
	}
	// Calculate and return the result
	return val1 + val2
}

func main() {
	// This is how your code will be called.
	// Your answer should be the sum of the two parameter after
	//   converting them to float64 values.
	// You can edit this code to try different testing cases.
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println(result)
}
