package millcode

import "math"

func F1(n int, a []float64, x float64) float64 {
	p := a[0]
	for i := 1; i <= n; i++ {
		p += a[i] * math.Pow(x, float64(i))
	}
	return p
}

func F2(n int, a []float64, x float64) float64 {
	p := a[n]
	for i := n; i > 0; i-- {
		p = a[i-1] + x*p
	}
	return p
}
