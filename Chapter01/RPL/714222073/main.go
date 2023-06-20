package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
	G "gorgonia.org/gorgonia"
)

func main() {
	g := G.NewGraph()

	a := gorgonia.NewScalar(g, G.Float64, G.WithName("a"))
	b := gorgonia.NewScalar(g, G.Float64, G.WithName("b"))

	x := G.NewScalar(g, G.Float64, G.WithName("x"))
	y := G.NewScalar(g, G.Float64, G.WithName("y"))

	form1 := G.Must(G.Pow(x, G.NewConstant(2.0)))
	form2 := G.Must(G.Pow(y, G.NewConstant(2.0)))
	form3 := G.Must(G.Pow(a, G.NewConstant(2.0)))
	form4 := G.Must(G.Pow(b, G.NewConstant(2.0)))
	form5 := G.Must(G.Div(form1, form3))
	form6 := G.Must(G.Div(form2, form4))
	form7 := G.Must(G.Add(form5, form6))
	form8 := G.Must(G.Sub(G.NewConstant(1.0), form7))

	machine := G.NewTapeMachine(g)
	defer machine.Close()

	G.Let(a, 2.0)
	G.Let(b, 4.0)
	G.Let(x, 1.0)
	G.Let(y, 3.0)

	if err := machine.RunAll(); err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Hasilnya : %v", form8.Value().Data())

}
