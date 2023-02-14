//bisection method for
//calculation of equations

package main

import (
	"fmt"
	"math"
)

func signs(x float64) float64 {
	return x - math.Pow(math.E, -x)
}

func bisection(a, b float64) float64 {

	if signs(a)*signs(b) >= 0 {
		// print("You have not assumed right a and b\n")
		return -1
	}

	c := a
	counter := 1

	for (b - a) >= 0.001 {

		fmt.Println("Number of iteration:", counter, "\t\t", a, b, (a+b)/2)

		c = (a + b) / 2

		if signs(c) == 0.0 {
			break
		}

		if signs(c)*signs(a) < 0 {
			b = c
		} else {
			a = c
		}

		counter += 1
	}
	return c
}

func main() {
	fmt.Println()
	var a float64 = 0
	var b float64 = 1
	c := bisection(a, b)
	fmt.Println()
	fmt.Printf("The value of root is: %v", c)
	// fmt.Println(signs(0.5))

}
