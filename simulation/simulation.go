package simulation

import (
	_ "fmt"
	_ "math"
	_ "time"
)

const (
	G              = 1.635
	interval       = 0.001
	endTime        = 200000
	slownessFactor = 0.1
)

type Simulation struct {
	lander *Lander
}

func (simulation *Simulation) Start() {
	simulation.lander = new(Lander)
	simulation.lander.velocity.X = 400
	simulation.lander.position.Y = 600

	simulation.lander.state.Fuel = 100
	simulation.lander.state.Exploded = false

	simulation.lander.velocity.X = 0
	simulation.lander.velocity.Y = 0

	simulation.lander.thrust = 5.0
	simulation.lander.crashTolerance = 2.0
}

func (simulation *Simulation) GetLander() *Lander {
	return simulation.lander
}

func (lander *Lander) GetPosition() Vector2D {
	return lander.position
}

func (lander *Lander) GetLanderState() LanderState {
	return lander.state
}

func (simulation *Simulation) Update(timeDelta int64, TrusterOn bool) {
	var acceleration float64 = 0

	if TrusterOn {
		acceleration += simulation.lander.thrust
	}

	if simulation.lander.position.Y > 0.0 {
		acceleration -= G
	} else {
		if simulation.lander.velocity.Y > simulation.lander.crashTolerance {
			simulation.lander.state.Exploded = true
			simulation.lander.state.Fuel = 0
		}

		simulation.lander.velocity.Y = 0.0
		simulation.lander.position.Y = 0.0
	}

	var interval float64 = float64(timeDelta) / (1000000000 * slownessFactor)

	simulation.lander.velocity.Y += acceleration * interval
	simulation.lander.position.Y += simulation.lander.velocity.Y * interval
}
