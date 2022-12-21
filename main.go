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

func main() {
	d := Dimension{5, 6, 7} // initialized to 5, 6, 7
	fmt.Println(d.Area())   // here it changes the value of width to 75 using pointer
	fmt.Println(d)          // this prints 5, 75, 7 as the dimension struct values
	fmt.Println(d.Volume()) // tyring to update width to 100
	fmt.Println(d)          // it will still pring 5, 75, 7 .. because here we are not trying to update with pointers again
}

// ReadMe
// Why pointers
// If we use pointers, the values will be permanently change
// to the pointer address. We can only modify it using the pointers later
// changing the value again will not be enough

func (d *Dimension) Area() int {
	// changing the value of the width to new value using pointers
	// so in future if any other method tried to update it with direct
	// value modification will not change it
	d.width = 75
	// if you need to change it, you need to change it later using the
	// pointers itself. direct assignment wont work again

	return d.length * d.width
}

// adding another method to the struct
func (d Dimension) Volume() int {
	d.width = 100
	return d.length * d.width * d.height
}
