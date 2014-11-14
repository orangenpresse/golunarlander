package lander

type Thruster struct {
	Acceleration    float64
	FuelConsumption float64
	Thrusting       bool
}

func (thruster *Thruster) CalculateAccelerationDelta(interval float64, tank *Tank) float64 {
	if thruster.Thrusting {
		tank.Level -= thruster.FuelConsumption * interval
		return thruster.Acceleration
	} else {
		return 0.0
	}
}
