package simulation

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

type Tank struct {
	Size  float64
	Level float64
}

type Lander struct {
	position       Vector2D
	velocity       Vector2D
	thrusterBottom Thruster
	thrusterLeft   Thruster
	thrusterRight  Thruster
	tank           Tank

	crashTolerance float64
	exploded       bool
	options        *Options
}

func New(options *Options) *Lander {
	lander := new(Lander)

	lander.options = options

	lander.position.X = 0
	lander.position.Y = 800

	lander.velocity.X = 0
	lander.velocity.Y = 0

	lander.thrusterBottom.Acceleration = 5.0
	lander.thrusterBottom.FuelConsumption = 5.0

	lander.thrusterLeft.Acceleration = 5.0
	lander.thrusterLeft.FuelConsumption = 1.0

	lander.thrusterRight.Acceleration = 5.0
	lander.thrusterRight.FuelConsumption = 1.0

	lander.tank.Size = 100
	lander.tank.Level = lander.tank.Size

	lander.exploded = false
	lander.crashTolerance = 5.0

	return lander
}

func (lander *Lander) GetPosition() Vector2D {
	return lander.position
}

func (lander *Lander) IsExploded() bool {
	return lander.exploded
}

func (lander *Lander) IsLanded() bool {
	return math.Abs(lander.position.Y) < 0.01 && math.Abs(lander.velocity.Y) < 0.01
}

func (lander *Lander) GetFuelLevel() int64 {
	return int64((100.0 / lander.tank.Size) * lander.tank.Level)
}

func (lander *Lander) setThrust(state ThrusterState) {

	if lander.options.DebugMode {
		lander.tank.Level = lander.tank.Size
		lander.exploded = false
	}

	lander.thrusterBottom.Thrusting = state.Bottom && lander.tank.Level > 0.0
	lander.thrusterLeft.Thrusting = state.Left && lander.tank.Level > 0.0
	lander.thrusterRight.Thrusting = state.Right && lander.tank.Level > 0.0
}

func (lander *Lander) IsThrusting() ThrusterState {
	return ThrusterState{lander.thrusterBottom.Thrusting, lander.thrusterLeft.Thrusting, lander.thrusterRight.Thrusting}
}

func (this *Lander) Update(timeInterval float64, thrusterState ThrusterState) {
	acceleration := Vector2D{}

	this.setThrust(thrusterState)

	acceleration.Y += this.thrusterBottom.CalculateAccelerationDelta(timeInterval, &this.tank)
	acceleration.X += this.thrusterLeft.CalculateAccelerationDelta(timeInterval, &this.tank)
	acceleration.X -= this.thrusterRight.CalculateAccelerationDelta(timeInterval, &this.tank)

	acceleration.Y += this.calculateFallingAcceleration()

	this.updatePosition(acceleration, timeInterval)
}

func (this *Lander) updatePosition(acceleration Vector2D, timeInterval float64) {
	this.velocity.Y += acceleration.Y * timeInterval
	this.position.Y += this.velocity.Y * timeInterval

	this.velocity.X += acceleration.X * timeInterval
	this.position.X += this.velocity.X * timeInterval
}

func (this *Lander) calculateFallingAcceleration() float64 {
	if this.position.Y > 0.0 {
		return -G
	} else {
		if this.velocity.Y < -this.crashTolerance && !this.options.DebugMode {
			fmt.Printf("Crashed: v=%f\n", this.velocity.Y)
			this.exploded = true
			this.tank.Level = 0
		}

		this.velocity.Y = 0.0
		this.position.Y = 0.0
		this.velocity.X = 0.0

		return 0.0
	}
}
