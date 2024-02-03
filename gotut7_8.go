package main

/*
Part 7-8 of Go tutorial.
Methods in Go Value receivers / Pointer receivers
Methods that just access values are called value receivers and
Methods that can modify information are pointer receivers.
*/
import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

// Simple car struct
type car struct {
	gas_pedal      uint16 //min 0 - max 65535
	break_pedal    uint16
	steering_wheel int16 // -32k - +32k
	top_speed_kmh  float64
}

// Method to calculate kmh
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

// Pointer receiver method to change the top speed
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{
		gas_pedal:      22341,
		break_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0}

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
