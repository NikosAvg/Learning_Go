package main

import (
	"fmt"
)

/*
Part 3 of Go tutorial.
Types in Go
*/

/*
Basic types in Go:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

*/

/*
Function definition in Go
func Name_of_function (var1 type1, ... ,varn typen) (return type1, ... ,return type n){
	.......
	return ...
}


*/

//	func add(x float64, y float64) float64 {
//		return x + y
//	}
func add(x, y float32) float32 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}
func main() {
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5
	// var num1, num2 float64 = 5.6, 9.5
	var num1, num2 float32 = 5.6, 9.5

	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1, w2))
}
