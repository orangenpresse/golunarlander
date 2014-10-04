package simulation

type Vector2D struct {
	X float64
	Y float64
}

type Thruster struct {
	Acceleration    float64
	FuelConsumption float64
}

type Tank struct {
	Size  float64
	Level float64
}

type Lander struct {
	position Vector2D
	velocity Vector2D
	thruster Thruster
	tank     Tank

	crashTolerance float64
	exploded       bool
}

func (lander *Lander) GetPosition() Vector2D {
	return lander.position
}

func (lander *Lander) IsExploded() bool {
	return lander.exploded
}

func (lander *Lander) GetFuelLevel() int64 {
	return int64((100.0 / lander.tank.Size) * lander.tank.Level)
}

func New() *Lander {
	lander := new(Lander)

	lander.position.X = 400
	lander.position.Y = 600

	lander.velocity.X = 0
	lander.velocity.Y = 0

	lander.thruster.Acceleration = 5.0
	lander.thruster.FuelConsumption = 1.0

	lander.tank.Size = 100
	lander.tank.Level = lander.tank.Size

	lander.exploded = false
	lander.crashTolerance = 2.0

	return lander
}
