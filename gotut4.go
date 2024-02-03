package main

import "fmt"

/*
Part 4 of Go tutorial.
Pointers Introduction
*/

func main() {
	x := 15
	a := &x         // memory address
	fmt.Println(a)  // Prints the memory address of the variable x
	fmt.Println(*a) // Prints the value of the variable a points to
	*a = 5          // Change the value of the variable this points at
	fmt.Println(x)
}
