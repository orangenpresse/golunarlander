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

type ThrusterState struct {
	Bottom bool
	Left   bool
	Right  bool
}

func (simulation *Simulation) Start() {
	simulation.lander = New()
}

func (simulation *Simulation) GetLander() *Lander {
	return simulation.lander
}

func (simulation *Simulation) Update(timeDelta int64, thrusterState ThrusterState) {
	var interval float64 = float64(timeDelta) / (1000000000 * slownessFactor)
	//	var acceleration float64 = 0

	simulation.lander.Update(interval, thrusterState)

	/*	if simulation.lander.thruster.Thrusting {
			acceleration += simulation.lander.thruster.Acceleration
			simulation.lander.tank.Level -= simulation.lander.thruster.FuelConsumption * interval
		}

		if simulation.lander.position.Y > 0.0 {
			acceleration -= G
		} else {
			if simulation.lander.velocity.Y < -simulation.lander.crashTolerance {
				fmt.Printf("Crashed: v=%f\n", simulation.lander.velocity.Y)
				simulation.lander.exploded = true
				simulation.lander.tank.Level = 0
			}

			simulation.lander.velocity.Y = 0.0
			simulation.lander.position.Y = 0.0
		}

		simulation.lander.velocity.Y += acceleration * interval
		simulation.lander.position.Y += simulation.lander.velocity.Y * interval
	*/
}
