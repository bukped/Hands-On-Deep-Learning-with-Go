package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()

	// definisi konstanta a dan b
	a := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("a"))
	b := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("b"))

	// definisi variabel x dan y
	x := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))

	// persamaan elips ((x/a)^2) + ((y/b)^2) - 1
	form1 := gorgonia.Must(gorgonia.Pow(x, gorgonia.NewConstant(2.0)))
	form2 := gorgonia.Must(gorgonia.Pow(y, gorgonia.NewConstant(2.0)))
	form3 := gorgonia.Must(gorgonia.Pow(a, gorgonia.NewConstant(2.0)))
	form4 := gorgonia.Must(gorgonia.Pow(b, gorgonia.NewConstant(2.0)))
	form5 := gorgonia.Must(gorgonia.Div(form1, form3))
	form6 := gorgonia.Must(gorgonia.Div(form2, form4))
	form7 := gorgonia.Must(gorgonia.Add(form5, form6))
	form8 := gorgonia.Must(gorgonia.Sub(gorgonia.NewConstant(1.0), form7))

	// kompilasi dan eksekusi graph
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()

	// inisiasi nilai konstanta a dan b
	gorgonia.Let(a, 2.0)
	gorgonia.Let(b, 4.0)

	// inisiasi nilai variabel x dan y
	gorgonia.Let(x, 1.0)
	gorgonia.Let(y, 3.0)

	// jalankan perhitungan
	if err := machine.RunAll(); err != nil {
		fmt.Println(err)
	}

	// tampilkan hasil
	fmt.Printf("Hasilnya : %v", form8.Value().Data())
}
