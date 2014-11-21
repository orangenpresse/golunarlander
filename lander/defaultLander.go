package lander

import (
	"fmt"
	data "github.com/orangenpresse/golunarlander/dataObjects"
	"math"
)

type DefaultLander struct {
	position data.Vector2D
	velocity data.Vector2D

	thrusterBottom Thruster
	thrusterLeft   Thruster
	thrusterRight  Thruster
	tank           Tank

	crashTolerance float64
	exploded       bool
	options        *data.Options
}

func NewDefaultLander(options *data.Options) *DefaultLander {
	lander := new(DefaultLander)

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

func (lander *DefaultLander) GetPosition() data.Vector2D {
	return lander.position
}

func (lander *DefaultLander) IsExploded() bool {
	return lander.exploded
}

func (lander *DefaultLander) IsLanded() bool {
	return math.Abs(lander.position.Y) < 0.01 && math.Abs(lander.velocity.Y) < 0.01
}

func (lander *DefaultLander) GetFuelLevel() int64 {
	return int64((100.0 / lander.tank.Size) * lander.tank.Level)
}

func (lander *DefaultLander) IsThrusting() data.ThrusterState {
	return data.ThrusterState{lander.thrusterBottom.Thrusting, lander.thrusterLeft.Thrusting, lander.thrusterRight.Thrusting}
}

func (this *DefaultLander) Update(timeInterval float64, thrusterState data.ThrusterState) {
	acceleration := data.Vector2D{}

	this.setThrust(thrusterState)

	acceleration.Y += this.thrusterBottom.CalculateAccelerationDelta(timeInterval, &this.tank)
	acceleration.X += this.thrusterLeft.CalculateAccelerationDelta(timeInterval, &this.tank)
	acceleration.X -= this.thrusterRight.CalculateAccelerationDelta(timeInterval, &this.tank)

	acceleration.Y += this.calculateFallingAcceleration()

	this.updatePosition(acceleration, timeInterval)
}

func (this *DefaultLander) SetPosition(position data.Vector2D) {
	this.position.X = position.X
	this.position.Y = position.Y
}

func (this *DefaultLander) GetLanderState() LanderState {
	state := LanderState{}

	state.Position = this.position
	state.Velocity = this.velocity
	state.Rotation = 0
	state.ThrusterState = data.ThrusterState{this.thrusterBottom.Thrusting, this.thrusterLeft.Thrusting, this.thrusterRight.Thrusting}
	state.Exploded = this.exploded
	state.Landed = math.Abs(this.position.Y) < 0.01 && math.Abs(this.velocity.Y) < 0.01

	return state
}

// Private
func (lander *DefaultLander) setThrust(state data.ThrusterState) {

	if lander.options.DebugMode {
		if lander.tank.Level <= 0.0 {
			lander.tank.Level = lander.tank.Size
		}
		lander.exploded = false
	}

	lander.thrusterBottom.Thrusting = state.Bottom && lander.tank.Level > 0.0
	lander.thrusterLeft.Thrusting = state.Left && lander.tank.Level > 0.0
	lander.thrusterRight.Thrusting = state.Right && lander.tank.Level > 0.0
}

func (this *DefaultLander) updatePosition(acceleration data.Vector2D, timeInterval float64) {
	this.velocity.Y += acceleration.Y * timeInterval
	this.position.Y += this.velocity.Y * timeInterval

	this.velocity.X += acceleration.X * timeInterval
	this.position.X += this.velocity.X * timeInterval
}

func (this *DefaultLander) calculateFallingAcceleration() float64 {
	if this.position.Y > 0.0 {
		return -data.G
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
