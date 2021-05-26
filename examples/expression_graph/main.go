package main

import (
	"fmt"
	mat "github.com/nlpodyssey/spago/pkg/mat32"
	"github.com/nlpodyssey/spago/pkg/ml/ag"
	"github.com/nlpodyssey/spago/pkg/ml/ag/encoding/dot"
	"log"
	"os"
)

func main() {
	g := ag.NewGraph()

	w := g.NewVariableWithName(mat.NewScalar(3.0), true, "w")
	b := g.NewVariableWithName(mat.NewScalar(1), true, "b")
	x := g.NewVariableWithName(mat.NewScalar(2), false, "x")

	// z = σ(wx+b)
	y := g.Add(g.Mul(w, x), b)
	z := g.Sigmoid(y)

	fmt.Fprintln(os.Stderr, z.Value()) // ignore errors

	out, err := dot.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
