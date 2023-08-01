package main

import (
	"fmt"

	tf "github.com/galeone/tensorflow/tensorflow/go"
)

func main() {
	// Initialize the TensorFlow Graph
	graph := tf.NewGraph()

	// Create the first constant tensor with value 3.0
	tensorA, _ := tf.NewTensor(3.0)

	// Create the second constant tensor with value 5.0
	tensorB, _ := tf.NewTensor(5.0)

	// Define the addition operation
	addOp, _ := graph.Add(tensorA, tensorB)

	// create new session to run the computation
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	// Execute the operation and get the result
	result, err := session.Run(nil, []tf.Output{addOp}, nil)
	if err != nil {
		panic(err)
	}

	// The result is single-element tensor, so we extract the value
	additionResult := result[0].Value().(float32)

	fmt.Println("Addition result: ", additionResult)
}
