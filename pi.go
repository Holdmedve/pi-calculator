package main

func CalculatePi() float64 {
	n := 2.0
	result := 3.0
	terms := 100

	for i := 1; i <= terms; i++ {
		term := 4 / (n * (n+1) * (n+2))
		if i % 2 == 0 {
			result -= term
		} else {
			result += term
		}
		n += 2.0
	}

	return result
}