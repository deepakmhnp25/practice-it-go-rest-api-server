package main

import (
	"fmt"
)

// defining a struct, very similar to a pojo class in java
type Dimension struct {
	length int
	width  int
	height int
}

// adding a method to the struct (class)
// methods are written outside the struct definition
func (d Dimension) Area() int {
	return d.length * d.width
}

// adding another method to the struct
func (d Dimension) Volume() int {
	return d.length * d.width * d.height
}

func main() {
	d := Dimension{5, 5, 5}
	fmt.Println("The Area is ", d.Area)
	fmt.Println("The Volume is ", d.Volume)
}
