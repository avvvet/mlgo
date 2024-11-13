package main

import (
	"fmt"
	"math"
)

// define initial weights
var (
	w1, w2 float64 = 0.5, -0.5
	w3, w4 float64 = 0.3, 0.8
	w5, w6 float64 = -0.7, 0.6
)

// sigmoid activation function
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
	s := sigmoid(x)
	return s * (1 - s)
}

// forward pass

func forwardPass(x1, x2 float64) float64 {
	// l hidden layer activation
	h1 := sigmoid(w1*x1 + w2*x2)
	h2 := sigmoid(w3*x1 + w4*x2)

	// output layer activation
	output := sigmoid(w5*h1 + w6*h2)

	return output
}

func main() {
	input := 0.5
	fmt.Printf("sigmoid(%v) = %v\n", input, sigmoid(input))
	fmt.Printf("sigmoid derivative(%v) = %v\n", input, sigmoidDerivative(input))

	// forward pass
	x1, x2 := 1.0, 0.5
	output := forwardPass(x1, x2)
	fmt.Printf("output of the network: %v\n", output)
}
