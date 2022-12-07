package math

// Нахождения среднего в срезе чисел
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Нахождения минимального значения в срезе чисел
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}

	return min
}

// Нахождение максимального значения в срезе чисел
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max

}
