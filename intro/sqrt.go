package main

import (
	"fmt"
	"math"
)

func mySqrt(data float64) float64 {
	var z, current, delta float64 = 1.0, 0, 1

	for delta > 0.00001 {
		current = z - (math.Pow(z, 2)-data)/(2*z)
		delta = math.Abs(current - z)
		fmt.Println(delta)
		z = current

	}
	return current
}
