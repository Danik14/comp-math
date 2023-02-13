package main

import (
	"fmt"
	"math"
)

func getEq(x float64) float64 {
	// we are taking equation as x^3+x-1
	var f float64 = math.Pow(x, 3) - 2*x - 5
	return f
}

func secant(x1, x2, E float64) {
	var n float64 = 0
	var xm float64 = 0
	var x0 float64 = 0
	var c float64 = 0
	if getEq(x1)*getEq(x2) < 0 {
		// calculate the intermediate value
		x0 = (x1*getEq(x2) - x2*getEq(x1)) / (getEq(x2) - getEq(x1))

		// check if x0 is root of equation or not
		c = getEq(x1) * getEq(x0)

		// update the value of interval
		x1 = x2
		x2 = x0

		// update number of iteration
		n++

		xm = (x1*getEq(x2) - x2*getEq(x1)) / (getEq(x2) - getEq(x1))

		counter := 2
		for math.Abs(xm-x0) >= E {
			fmt.Printf("x%d=%v\n", counter, x2)
			counter++

			// calculate the intermediate value
			x0 = (x1*getEq(x2) - x2*getEq(x1)) / (getEq(x2) - getEq(x1))

			// check if x0 is root of equation or not
			c = getEq(x1) * getEq(x0)

			// update the value of interval
			x1 = x2
			x2 = x0

			// update number of iteration
			n++

			// if x0 is the root of equation then break the loop
			if c == 0 {
				break
			}
			xm = (x1*getEq(x2) - x2*getEq(x1)) / (getEq(x2) - getEq(x1))
		}

		fmt.Println("Root of the given equation=", x0)
		fmt.Println("No. of iterations = ", n)
	} else {
		fmt.Println("Can not find a root in the given interval")
	}
}

// Driver code
func main() {
	// initializing the values
	var x0 float64 = 2
	var x1 float64 = 3
	var E float64 = 0.00001
	secant(x0, x1, E)
	fmt.Println(getEq(3))
}
