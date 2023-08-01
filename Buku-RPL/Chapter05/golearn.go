package main

import (
	"fmt"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

func main() {
	// Generate some random data for demonstration purposes
	x := [][]float64{[]float64{1, 2, 3, 4, 5}}
	y := []float64{3, 4, 2, 5, 6}

	// Create a new regression model
	model := linear.NewLeastSquares(base.BatchGA, .000001, 0, 800, x, y)

	// learn
	model.Learn()

	// Make predictions
	predictions, _ := model.Predict(y)

	// Print the predictions
	fmt.Println("Predictions:", predictions)
}
