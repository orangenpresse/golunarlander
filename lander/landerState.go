package lander

import (
	data "github.com/orangenpresse/golunarlander/dataObjects"
)

type LanderState struct {
	ThrusterState data.ThrusterState
	Position      data.Vector2D
	Velocity      data.Vector2D
	Rotation      float64
	Exploded      bool
	Landed        bool
}
