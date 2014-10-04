package simulation

import (
	_ "fmt"
	_ "math"
	_ "time"
)

type Vector2D struct {
	X float64
	Y float64
}

const (
	G              = 1.635
	interval       = 0.001
	endTime        = 200000
	slownessFactor = 0.1
)

type Lander struct {
	position       Vector2D
	velocity       Vector2D
	thrust         float64
	crashTolerance float64
	state          LanderState
}

type Simulation struct {
	lander *Lander
}

type LanderState struct {
	exploded bool
	fuel     int // 100% - 0%
}

func (simulation *Simulation) Start() {
	simulation.lander = new(Lander)
	simulation.lander.position.X = 400
	simulation.lander.position.Y = 600

	simulation.lander.state.fuel = 100
	simulation.lander.state.exploded = false

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

func (simulation *Simulation) Update(timeDelta int64, thrusterOn bool) {
	var acceleration float64 = 0

	if thrusterOn {
		acceleration += simulation.lander.thrust
	}

	if simulation.lander.position.Y > 0.0 {
		acceleration -= G
	} else {
		if simulation.lander.velocity.Y > simulation.lander.crashTolerance {
			simulation.lander.state.exploded = true
			simulation.lander.state.fuel = 0
		}

		simulation.lander.velocity.Y = 0.0
		simulation.lander.position.Y = 0.0
	}

	var interval float64 = float64(timeDelta) / (1000000000 * slownessFactor)

	simulation.lander.velocity.Y += acceleration * interval
	simulation.lander.position.Y += simulation.lander.velocity.Y * interval
}
