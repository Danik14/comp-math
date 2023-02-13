package main

import (
	"fmt"
)

// 3x-0.1y-0.2z=7.85; 0.1x+7y-0.3z=-19.3; 0.3x-0.2y+10z=71.4
func findx1(x2, x3 float64) float64 {
	// return (12 + x2 - 2*x3) / 5
	return (7.85 + 0.1*x2 + 0.2*x3) / 3
}

func findx2(x1, x3 float64) float64 {
	return (-19.3 - 0.1*x1 + 0.3*x3) / 7
}

func findx3(x1, x2 float64) float64 {
	// return (6 - x1 - x2) / 4
	return (71.4 - 0.3*x1 + 0.2*x2) / 10
}

func gauss_seidal() {
	x1, x2, x3 := 0.0, 0.0, 0.0
	for i := 0; i < 10; i++ {
		x1 = findx1(x2, x3)
		x2 = findx2(x1, x3)
		x3 = findx3(x1, x2)
		fmt.Printf("Iteration #%d x1: %f,\tx2: %f,\tx3: %f\n", i+1, x1, x2, x3)
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	gauss_seidal()
}
