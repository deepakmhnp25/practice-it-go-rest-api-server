package main

import (
	"fmt"
)

// Method definitions with single return types
func area(length, width, height int) int {
	return length * width
}

// Method definiton with multiple return types
func dimentions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height
	return
}

func main() {
	area, volume := dimentions(5, 5, 5)
	fmt.Println("area is ", area, ", volume is ", volume)
}
