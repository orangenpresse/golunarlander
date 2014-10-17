package simulation

import (
	"fmt"
	_ "math"
	_ "time"
)

const (
	G              = 1.635
	interval       = 0.001
	endTime        = 200000
	slownessFactor = 0.1

	SimulationAreaWidth = 1280
)

type Options struct {
	DebugMode bool
}

type Simulation struct {
	lander *Lander
}

type ThrusterState struct {
	Bottom bool
	Left   bool
	Right  bool
}

func (simulation *Simulation) Start(options *Options) {
	simulation.lander = New(options)
	fmt.Println("Simulation started")
}

func (simulation *Simulation) GetLander() *Lander {
	return simulation.lander
}

func (simulation *Simulation) Update(timeDelta int64, thrusterState ThrusterState) {
	var interval float64 = float64(timeDelta) / (1000000000 * slownessFactor)
	simulation.lander.Update(interval, thrusterState)

	simulation.enforceLanderInSimulationArea()

	//fmt.Printf("%f\t%f\n", simulation.lander.velocity.X, simulation.lander.position.X)
	//fmt.Println(simulation.lander.position)
	//fmt.Printf("Bottom: %b, Left: %b, Right %b\n", thrusterState.Bottom, thrusterState.Left, thrusterState.Right)
}

func (simulation *Simulation) enforceLanderInSimulationArea() {
	if simulation.lander.position.X > (SimulationAreaWidth / 2) {
		simulation.lander.position.X = -(SimulationAreaWidth / 2)
	} else if simulation.lander.position.X < -(SimulationAreaWidth / 2) {
		simulation.lander.position.X = (SimulationAreaWidth / 2)
	}
}
