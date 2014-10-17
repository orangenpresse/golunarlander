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
	simulation.lander.Update(interval, thrusterState)
}
