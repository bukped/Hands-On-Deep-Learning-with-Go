package main

import (
	"fmt"

	G "gorgonia.org/gorgonia"
)

func main() {
	g := G.NewGraph()

	// defenisikan konstanta a dan b
	a := G.NewScalar(g, G.Float64, G.WithName("a"))
	b := G.NewScalar(g, G.Float64, G.WithName("b"))

	// defenisikan variable x dan y
	x := G.NewScalar(g, G.Float64, G.WithName("x"))
	y := G.NewScalar(g, G.Float64, G.WithName("y"))

	// formula persamaan elips (x/a)^2 + (y/b)^2 - 1
	formula1 := G.Must(G.Pow(x, G.NewConstant(2.0)))
	formula2 := G.Must(G.Pow(y, G.NewConstant(2.0)))
	formula3 := G.Must(G.Pow(a, G.NewConstant(2.0)))
	formula4 := G.Must(G.Pow(b, G.NewConstant(2.0)))
	formula5 := G.Must(G.Div(formula1, formula3))
	formula6 := G.Must(G.Div(formula2, formula4))
	formula7 := G.Must(G.Add(formula5, formula6))
	formula8 := G.Must(G.Sub(G.NewConstant(1.0), formula7))

	machine := G.NewTapeMachine(g)
	defer machine.Close()

	// inisiasi nilai konstanta a dan b
	G.Let(a, 2.0)
	G.Let(b, 4.0)

	// inisiasi nilai variabel x dan y
	G.Let(x, 1.0)
	G.Let(y, 3.0)

	// menjalankan model fungsi
	if err := machine.RunAll(); err != nil {
		fmt.Print(err)
	}

	// cetak hasil
	fmt.Printf("Hasilnya : %v", formula8.Value().Data())

}
