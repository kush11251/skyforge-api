package utils

import (
	"math"
)

func CalculateDistance(a float64, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
