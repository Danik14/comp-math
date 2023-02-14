package main

import (
	"fmt"
)

func getSlope(n int, sumX float64, sumY float64, sumXY float64, sumXX float64) float64 {
	return (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXX - (sumX * sumX))
}

func getIntercept(n int, m float64, sumY float64, sumX float64) float64 {
	if sumY-m*sumX < 1e-7 {
		return 0.0
	}
	return (sumY - m*sumX) / float64(n)
}

func main() {
	fmt.Println("Enter the number of points:")
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println("Enter the points:")
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f %f", &x[i], &y[i])
	}

	fmt.Println(x, y)

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumXX := 0.0
	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXX += x[i] * x[i]
	}

	slope := getSlope(n, sumX, sumY, sumXY, sumXX)
	intercept := getIntercept(n, slope, sumY, sumX)

	fmt.Printf("The slope is: %f\n", slope)
	fmt.Printf("The intercept is: %f\n", intercept)
}

//7
//-3 -2 -1 0 1 2 3
// 4.63 2.11 0.67 0.09 0.63 2.15 4.58
