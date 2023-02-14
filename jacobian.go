//jacobian method for
//calculation of equations

package main

import (
	"fmt"
)

// Solve the equations 5x-y+z=10; 2x+4y=12; x+y+5z=-1 by Jacobi's method and choose the correct
// 13x+5y+-3z+u=18; 2x+12y+z-4u=13; x-4y+10z+u=29; 2x+y-3z+9u=31
func findx1(x2, x3 float64) float64 {
	// return (12 + x2 - 2*x3) / 5
	return (10 + x2 - x3) / 5
	// return
}

func findx2(x1, x3 float64) float64 {
	// return (-25 - 3*x1 + 2*x3) / 8
	return (12 - 2*x1) / 4
}

func findx3(x1, x2 float64) float64 {
	// return (6 - x1 - x2) / 4
	return (-1 - x1 - x2) / 5
}

func findx4(x1, x2, x3 float64)

func jacobi() {
	// var x1, x2, x3 float64
	// x1, x2, x3 = findx1(x2, x3), findx2(x1, x3), findx3(x1, x2)
	x1, x2, x3 := 0.0, 0.0, 0.0
	tmpx1, tmpx2, tmpx3 := 0.0, 0.0, 0.0
	for i := 0; i < 15; i++ {
		tmpx1 = findx1(x2, x3)
		tmpx2 = findx2(x1, x3)
		tmpx3 = findx3(x1, x2)
		x1 = tmpx1
		x2 = tmpx2
		x3 = tmpx3
		fmt.Printf("Iteration #%d x1: %f,\tx2: %f,\tx3: %f\n", i+1, x1, x2, x3)
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	jacobi()
}
