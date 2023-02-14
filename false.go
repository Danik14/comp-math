//false point method for
//calculation of equations

package main

import (
	"fmt"
	"math"
)

const (
	MAX_ITER = 15
)

// // An example function whose solution is determined using
// // Regular Falsi Method. The function is x^3 - x^2  + 2
func getEq(x float64) float64 {
	//x^3-2x-5=0
	// var f float64 = math.Cos(x) - x*math.Pow(math.E, x)
	// var f float64 = x*math.Log10(x) - 1.2
	var f float64 = math.Pow(x, 3) - 2*x - 5
	return f
}

// // Prints root of func(x) in interval [a, b]
func regulaFalsi(a, b float64) {
	if getEq(a)*getEq(b) >= 0 {
		fmt.Println("You have not assumed right a and b")
		return
	}

	var c float64 = a // Initialize result

	for i := 0; i < MAX_ITER; i++ {

		// Find the point that touches x axis
		c = (a*getEq(b) - b*getEq(a)) / (getEq(b) - getEq(a))

		// Check if the above found point is root
		if getEq(c) == 0 {
			break
		} else if getEq(c)*getEq(a) < 0 { // Decide the side to repeat the steps
			b = c
		} else {
			a = c
		}
		fmt.Printf("x%d = %v\n", i+2, c)

	}
	fmt.Println("The value of root is : ", c)
}

// Driver program to test above function
func main() {
	// Initial values assumed
	// x^3-2x-5=0
	var a float64 = 2
	var b float64 = 3
	regulaFalsi(a, b)
}
