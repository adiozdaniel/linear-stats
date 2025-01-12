package pkg

func LinearRegression(x, y []float64) (float64, float64) {
	var n = float64(len(x))

	var sumX, sumY, sumXY, sumX2 float64

	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	b0 := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b1 := (sumY - b0*sumX) / n

	return b0, b1
}
