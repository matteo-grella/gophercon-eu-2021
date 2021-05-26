package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
	mat "github.com/nlpodyssey/spago/pkg/mat32"
	"github.com/nlpodyssey/spago/pkg/mat32/rand"
	"github.com/nlpodyssey/spago/pkg/ml/ag"
	"github.com/nlpodyssey/spago/pkg/ml/initializers"
	"github.com/nlpodyssey/spago/pkg/ml/losses"
	"github.com/nlpodyssey/spago/pkg/ml/nn"
	"github.com/nlpodyssey/spago/pkg/ml/optimizers/gd"
	"github.com/nlpodyssey/spago/pkg/ml/optimizers/gd/sgd"
	"time"
)

var _ nn.Model = &Model{}

type Model struct {
	nn.BaseModel
	W nn.Param `spago:"type:weights"`
}

func (m *Model) Forward(xs []ag.Node) []ag.Node {
	linear := func(x ag.Node) ag.Node {
		return m.Graph().Mul(m.W, x) // Wx
	}
	return ag.Map(linear, xs)
}

func main() {
	// Instantiate a new model.
	model := &Model{
		W: nn.NewParam(mat.NewEmptyDense(1, 1)),
	}

	// Initialize the model with random weights.
	initializers.XavierUniform(model.W.Value(), 1.0, rand.NewLockedRand(42))

	// Use a plain stochastic gradient descent method with momentum.
	updater := sgd.New(sgd.NewConfig(0.01, 0.9, false))

	// Creates a new optimizer for the model.
	optimizer := gd.NewOptimizer(updater, nn.NewDefaultParamsIterator(model))

	// Define the cost function, in this case Mean Squared Error.
	criterion := losses.MSESeq

	// Create dummy data of the type `y = 3x`.
	points := 1000
	xValues, yValues := getDummyDataset(points)

	// collect the losses over the epochs
	lossTrend := make([]float64, 0)

	for epoch := 0; epoch < 100; epoch++ {
		// you can beat the occurrence of a new epoch e.g. for learning rate annealing
		optimizer.IncEpoch()

		// get a new computational graph
		g := ag.NewGraph()

		// get a new processor for the current computation
		proc := nn.ReifyForTraining(model, g).(*Model)

		// Converting x and y values to graph nodes
		inputs := make([]ag.Node, points)
		labels := make([]ag.Node, points)
		for i := 0; i < points; i++ {
			inputs[i] = g.NewScalarWithName(xValues[i], fmt.Sprintf("x%d", i))
			labels[i] = g.NewScalarWithName(yValues[i], fmt.Sprintf("y%d", i))
		}

		// get output (i.e. prediction) from the model
		outputs := proc.Forward(inputs)

		// get loss for the predicted output
		loss := criterion(g, outputs, labels, true) // true = reduce mean

		// get the gradients
		g.Backward(loss)

		// update parameters
		optimizer.Optimize()

		// collect the loss for statistics
		lossTrend = append(lossTrend, float64(loss.ScalarValue()))

		asciigraph.Clear()

		fmt.Printf("Epoch %d | Loss %.4f | Learning coefficient: %.2f\n\n",
			epoch, loss.ScalarValue(), model.W.ScalarValue())

		fmt.Println(asciigraph.Plot(lossTrend, asciigraph.Precision(4), asciigraph.Height(10), asciigraph.Width(epoch)))
		time.Sleep(100 * time.Millisecond) // slow down the process just for demo ;)
	}

	fmt.Printf("\nLearned coefficient: %.2f\n", model.W.ScalarValue())
}

// getDummyDataset returns a dummy dataset for training using a very basic linear equation:
//    `y = 3x`
// Here, `x` is the independent variable and `y` is the dependent variable.
func getDummyDataset(n int) (xValues, yValues []float32) {
	xValues = make([]float32, n)
	yValues = make([]float32, n)
	for i := 0; i < n; i++ {
		x := float64(i) / float64(n) // data normalization
		xValues[i] = float32(x)
		yValues[i] = float32(3 * x)
	}
	return
}
