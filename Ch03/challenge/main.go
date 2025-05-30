// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import "fmt"

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	// Your code goes here.
	colors := make([]Color, 0)
	for i, name := range colorNames {
		color := Color{name, hexValues[i]}
		colors = append(colors, color)
	}
	// Create a slice of Color objects
	return colors
}

type Color struct {
	Name string
	Hex  int
}

func main() {
	// This is how your code will be called.
	// Your answer should be a slice of Color objects.
	// You can edit this code to try different testing cases.
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := slicesToObjects(colorNames, hexValues)
	fmt.Println(colors)
	return
}
