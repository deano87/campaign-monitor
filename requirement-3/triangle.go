package requirements

import "fmt"
import "math"

/**
Calculate a triangle area based on Heron's formula:
Area = √p(p−a)(p−b)(p−c)
where p = (a + b + c) / 2
*/
func getTriangleArea(len1 int, len2 int, len3 int) float64 {
	if !isValidTriangle(len1, len2, len3) {
		panic(fmt.Sprintf("InvalidTriangleException"))
	}

	var p = float64(len1+len2+len3) / 2 // perimeter
	var areaPower = p * (p - float64(len1)) * (p - float64(len2)) * (p - float64(len3))

	return math.Sqrt(areaPower)
}

/**
Returns true if the lengths provided can make a triangle, false otherwise
*/
func isValidTriangle(len1 int, len2 int, len3 int) bool {
	if len1 < 0 || len2 < 0 || len3 < 0 {
		return false
	}

	if len1+len2 <= len3 || len1+len3 <= len2 || len3+len2 <= len1 {
		return false
	}

	return true
}
