package main

import (
	"fmt"
)

func claculating(x int16, y uint8, z float32, F func(float32, float32, float32) float32) float32 {
	return F(float32(x), float32(y), z)
}
func main() {
	var (
		x int16   = 10
		y uint8   = 5
		z float32 = 1.0
	)
	fmt.Println(claculating(x, y, z, func(a, b, c float32) float32 { return 2*a + b*b - 3/c }))
}
