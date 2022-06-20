package main

import (
	"fmt"
	"math"
)

func main() {
	squareRootWithIf(8)
	squareRootWithIf(-2)
	squareRootWithIfElse(21)
	squareRootWithIfElse(-9)
	squareRootWithIfElseIf(12)
	squareRootWithIfElseIf(-4)
	squareRootWithIfElseIf(0)
}

func squareRootWithIf(arg int) {
	if arg < 0 {
		fmt.Printf("Argument must be positive: %d\n", arg)
		return
	}
	squareRoot := math.Sqrt(float64(arg))
	fmt.Printf("Square root of %d is %f\n", arg, squareRoot)
}

func squareRootWithIfElse(arg int) {
	if arg < 0 {
		fmt.Printf("Argument must be positive: %d\n", arg)
	} else {
		squareRoot := math.Sqrt(float64(arg))
		fmt.Printf("Square root of %d is %f\n", arg, squareRoot)
	}
}

func squareRootWithIfElseIf(arg int) {
	if arg < 0 {
		fmt.Printf("Argument must be positive: %d\n", arg)
	} else if arg == 0 {
		fmt.Printf("Square root of %d is %d\n", arg, 0)
	} else {
		squareRoot := math.Sqrt(float64(arg))
		fmt.Printf("Square root of %d is %f\n", arg, squareRoot)
	}
}
