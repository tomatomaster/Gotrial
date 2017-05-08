package newton

import (
	"fmt"
	"math"
)

//X4minus1 solve x^4 - 1
func X4minus1(x float64) float64 {
	result := formula(x)
	for {
		temp := formula(result)
		if math.Abs(result-temp) > 0.000000001 {
			result = temp
		} else {
			break
		}
	}
	return result
}

func formula(x float64) float64 {
	result := x - float64(math.Pow(x, float64(4))-1.0)/(4*math.Pow(x, float64(3)))
	fmt.Printf("x1=%f result=%f\n", x, result)
	return result
}
