package lander

import (
	data "github.com/orangenpresse/golunarlander/dataObjects"
)

type LanderInterface interface {
	GetPosition() data.Vector2D
	IsExploded() bool
	IsLanded() bool
	GetFuelLevel() int64
	IsThrusting() data.ThrusterState
	Update(timeDelta float64, thrusterState data.ThrusterState)

	SetPosition(position data.Vector2D)
}
