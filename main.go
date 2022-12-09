// Package ludmath implements some crazy Lud's math! it's nothing like the math you know
package ludmath

// Sum calculates Lud's sum
func Sum(xf ...float64) float64 {
	var sum float64
	for _, v := range xf {
		sum += v
	}
	return sum
}
