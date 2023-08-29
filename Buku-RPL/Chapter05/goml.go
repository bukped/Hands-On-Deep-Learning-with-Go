package main

import (
	"fmt"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

func main() {
	// Generate some random data for demo
	x := [][]float64{[]float64{1, 2, 3, 4, 5}}
	y := []float64{3, 4, 5}

	// create new regression model
	model := linear.NewLeastSquares(base.BatchGA, .0000001, 0, 800, x, y)

	// learn
	model.Learn()

	// make predictions
	predictions, _ := model.Predict(y)

	// show result
	fmt.Println("Predictions: ", predictions)
}
