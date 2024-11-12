package main

import (
	"fmt"
	"math"
)

// sigmoid activation function
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
	s := sigmoid(x)
	return s * (1 - s)
}

func main() {
	input := 0.5
	fmt.Printf("sigmoid(%v) = %v\n", input, sigmoid(input))
	fmt.Printf("sigmoid derivative(%v) = %v\n", input, sigmoidDerivative(input))
}
