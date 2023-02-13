package main

import (
	"fmt"
	"math"
)

// // An example function whose solution is determined using
// // Bisection Method. The function is x^3 - x^2  + 2
// x3-2x-5 = 0
func getFunc(x float64) float64 {
	// return x*math.Log10(x) - 1.2
	return math.Pow(x, 3) - 2*x - 5
}

// // Derivative of the above function which is 3*x^x - 2*x
func derivFunc(x float64) float64 {
	// return (math.Log(x) + 1) / math.Log(10)
	return 3*math.Pow(x, 2) - 2
}

// // Function to find the root
func newtonRaphson(x float64) {
	h := getFunc(x) / derivFunc(x)
	counter := 1
	for math.Abs(h) >= 0.00001 {
		fmt.Printf("Iteration #%d, value is %v\n", counter, x)
		counter++
		h = getFunc(x) / derivFunc(x)

		// x(i+1) = x(i) - f(x) / f'(x)
		x = x - h
	}

	fmt.Println("The value of the root is :", x)
}

func main() {
	var x0 float64 = 1 // Initial values assumed
	newtonRaphson(x0)
	fmt.Println(math.Pow(32, 0.25))
}
