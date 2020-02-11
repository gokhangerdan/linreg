package main

import "fmt"

func calcSum(x []float64) float64 {
	var sum float64 = 0.

	for i := range x {
		sum += x[i]
	}
	return sum
}

func calcMean(x []float64) float64 {
	var mean float64

	mean = calcSum(x) / float64(len(x))

	return mean
}

// Least Square Method
func linregTrainLSM(x []float64, y []float64) map[string]float64 {
	var xMean float64 = calcMean(x)
	var yMean float64 = calcMean(y)

	var m float64
	var b float64

	var l1 float64 = 0.
	var l2 float64 = 0.

	for i := range x {
		l1 += ((x[i] - xMean) * (y[i] - yMean))
		l2 += ((x[i] - xMean) * (x[i] - xMean))
	}

	m = l1 / l2

	b = yMean - m*xMean

	model := make(map[string]float64)
	model["b"] = b
	model["m"] = m

	return model
}

func linregPred(x float64, model map[string]float64) float64 {
	var m float64 = model["m"]
	var b float64 = model["b"]
	var y float64
	y = m*x + b
	return y
}

func main() {

	df := make(map[string]interface{})

	x := []float64{1., 2., 3., 4., 5., 6.}
	y := []float64{2., 4., 6., 8., 10., 12.}

	df["x"] = x
	df["y"] = y

	var trainX []float64 = df["x"].([]float64)
	var trainY []float64 = df["y"].([]float64)

	model := linregTrainLSM(trainX, trainY)

	pred := linregPred(7., model)

	fmt.Println(pred)
}
