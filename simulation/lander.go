package simulation

type Vector2D struct {
	X float64
	Y float64
}

type LanderState struct {
	Exploded bool
	Fuel     int // 100% - 0%
}

type Lander struct {
	position       Vector2D
	velocity       Vector2D
	thrust         float64
	crashTolerance float64
	state          LanderState
}

func (lander *Lander) GetPosition() Vector2D {
	return lander.position
}

func (lander *Lander) GetLanderState() LanderState {
	return lander.state
}
