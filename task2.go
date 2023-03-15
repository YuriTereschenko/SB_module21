package main

import "fmt"

func outerFunc(x, y int, F func(int, int) int) {
	defer fmt.Printf("Result of inner func is: %v \n", F(x, y))
	fmt.Println("Func has been finished")
}
func main() {
	vA := 10
	vB := 5
	outerFunc(vA, vB, func(a, b int) int { return a + b })
	outerFunc(vA, vB, func(a, b int) int { return a - b })
	outerFunc(vA, vB, func(a, b int) int { return a * b })
}
