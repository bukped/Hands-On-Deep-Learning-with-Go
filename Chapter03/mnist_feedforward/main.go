package main
import (
     "fmt"
     "log"
     . "gorgonia.org/gorgonia"
 )

 func main() {
	g := NewGraph()
  }

  a = NewScalar(g, Float64, WithName("a"))
  b = NewScalar(g, Float64, WithName("b"))

  c, err = Add(a,b)

  machine := NewTapeMachine(g)

  Let(a, 1.0)
  Let(b, 2.0)
  machine.RunAll()