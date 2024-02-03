package main

/*
Part 6 of Go tutorial.
Structs in Go
*/
import (
	"fmt"
)

// Simple car struct
type car struct {
	gas_pedal      uint16 //min 0 - max 65535
	break_pedal    uint16
	steering_wheel int16 // -32k - +32k
	top_speed_kmh  float64
}

func main() {
	a_car := car{
		gas_pedal:      22341,
		break_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0}

	fmt.Println(a_car.gas_pedal)

}
