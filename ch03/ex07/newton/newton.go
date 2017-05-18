package newton

import (
	"math/cmplx"
)

//X4minus1 solve x^4 - 1
func X4minus1(x complex128) complex128 {
	result := formula(x)
	for {
		temp := formula(result)
		if cmplx.Abs(result-temp) > 0.1 {
			result = temp
		} else {
			break
		}
	}
	return result
}

func formula(x complex128) complex128 {
	result := x - (x*x*x*x-1.0)/(4*x*x*x)
	return result
}
