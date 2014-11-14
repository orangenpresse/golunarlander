package simulation

import (
	"fmt"
	data "github.com/orangenpresse/golunarlander/dataObjects"
	"github.com/orangenpresse/golunarlander/lander"
	_ "math"
	_ "time"
)

const (
	interval       = 0.001
	endTime        = 200000
	slownessFactor = 0.1

	SimulationAreaWidth = 1280
)

type Simulation struct {
	lander lander.LanderInterface
}

func (simulation *Simulation) Start(options *data.Options) {
	simulation.lander = lander.NewDefaultLander(options)
	fmt.Println("Simulation started")
}

func (simulation *Simulation) GetLander() lander.LanderInterface {
	return simulation.lander
}

func (simulation *Simulation) Update(timeDelta int64, thrusterState data.ThrusterState) {
	var interval float64 = float64(timeDelta) / (1000000000 * slownessFactor)
	simulation.lander.Update(interval, thrusterState)

	simulation.enforceLanderInSimulationArea()

	//fmt.Printf("%f\t%f\n", simulation.lander.velocity.X, simulation.lander.position.X)
	//fmt.Println(simulation.lander.position)
	//fmt.Printf("Bottom: %b, Left: %b, Right %b\n", thrusterState.Bottom, thrusterState.Left, thrusterState.Right)
}

func (simulation *Simulation) enforceLanderInSimulationArea() {
	landerPos := simulation.lander.GetPosition()

	if landerPos.X > (SimulationAreaWidth / 2) {
		landerPos.X = -(SimulationAreaWidth / 2)
		simulation.lander.SetPosition(landerPos)
	} else if landerPos.X < -(SimulationAreaWidth / 2) {
		landerPos.X = (SimulationAreaWidth / 2)
		simulation.lander.SetPosition(landerPos)
	}
}
