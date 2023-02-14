package main

import (
	"fmt"
	"math"
)

func maxOffDiagonal(mat [][]float64) (int, int, float64) {
	var maxValue float64
	var maxI, maxJ int
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if math.Abs(mat[i][j]) > maxValue {
				maxValue = math.Abs(mat[i][j])
				maxI = i
				maxJ = j
			}
		}
	}
	return maxI, maxJ, maxValue
}

func rotate(mat [][]float64, p, q int) {
	n := len(mat)
	var theta float64
	if mat[p][p] == mat[q][q] {
		theta = math.Pi / 4
	} else {
		theta = 0.5 * math.Atan2(2*mat[p][q], mat[q][q]-mat[p][p])
	}
	c := math.Cos(theta)
	s := math.Sin(theta)
	for i := 0; i < n; i++ {
		if i != p && i != q {
			t := mat[i][p]
			mat[i][p] = c*t - s*mat[i][q]
			mat[i][q] = s*t + c*mat[i][q]
		}
	}
	t := mat[p][p]
	mat[p][p] = c*c*t - 2*s*c*mat[p][q] + s*s*mat[q][q]
	mat[q][q] = s*s*t + 2*s*c*mat[p][q] + c*c*mat[q][q]
	mat[p][q] = 0
	mat[q][p] = 0
}

func eigen(mat [][]float64) ([]float64, [][]float64) {
	var eigenValues []float64
	var eigenVectors [][]float64
	n := len(mat)
	for i := 0; i < n; i++ {
		eigenValues = append(eigenValues, mat[i][i])
		eigenVectors = append(eigenVectors, make([]float64, n))
		eigenVectors[i][i] = 1
	}
	for i := 0; i < 100; i++ {
		p, q, maxValue := maxOffDiagonal(mat)
		if maxValue < 1e-10 {
			break
		}
		rotate(mat, p, q)
		rotate(eigenVectors, p, q)
	}
	for i := 0; i < n; i++ {
		eigenValues[i] = mat[i][i]
	}
	return eigenValues, eigenVectors
}

func main() {
	mat := [][]float64{
		{8, -4, 2},
		{4, 0, 2},
		{0, -2, -4},
	}
	eigenValues, eigenVectors := eigen(mat)
	fmt.Println("Eigenvalues:", eigenValues)
	fmt.Println("Eigenvectors:")
	for i := 0; i < len(eigenVectors); i++ {
		fmt.Println(eigenVectors[i])
	}
}
