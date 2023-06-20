package main
import (
	 "fmt"
	 "log"

	 G "gorgonia.org/gorgonia"
)

func main() {
	// create a new graph
	g := G.NewGraph()
	// define variable
	a := G.NewScalar(g, G.Float64, G.WithName("a"))
	b := G.NewScalar(g, G.Float64, G.WithName("b"))
	c := G.NewScalar(g, G.Float64, G.WithName("c"))

	// define variable x
	x := G.NewScalar(g, G.Float64, G.WithName("x"))

	// parabola equation
	y := G.Must(G.Add(G.Must(G.Mul(a, G.Must(G.Pow(x, G.NewConstant(2.0))))), G.Must(G.Add(G.Must(G.Mul(b, x)), c))))

	// compile and run graph
	machine := G.NewTapeMachine(g)
	defer machine.Close()

	// assign value to variable
	G.Let(a, 2.0)
	G.Let(b, 3.0)
	G.Let(c, 1.0)

	// assign value to variable x
	G.Let(x, 0.0)

	// run calculation

	if err := machine.RunAll(); err != nil {
		log.Fatal(err)
	}

	// show result
	fmt.Printf("y = %v", y.Value())

}