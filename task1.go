package main

import (
	"fmt"
)

func claculating(x int16, y uint8, z float32) float32 {

	return 2*float32(x) + float32(y)*float32(y) - 3/z
}
func main() {
	var (
		x int16   = 10
		y uint8   = 5
		z float32 = 1.0
	)
	fmt.Println(claculating(x, y, z))
}
