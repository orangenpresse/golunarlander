package multiplayer

import (
	data "github.com/orangenpresse/golunarlander/dataObjects"
	"github.com/orangenpresse/golunarlander/lander"
)

type PlayerLander struct {
	Id   string
	Name string
	PosX float32
	PosY float32
}

func (pl *PlayerLander) GetPosition() data.Vector2D {
	return data.Vector2D{float64(pl.PosX), float64(pl.PosY)}
}

func (pl *PlayerLander) IsExploded() bool {
	return false
}

func (pl *PlayerLander) IsLanded() bool {
	return false
}

func (pl *PlayerLander) GetFuelLevel() int64 {
	return 0
}

func (pl *PlayerLander) GetLanderState() lander.LanderState {
	return lander.LanderState{}
}

func (pl *PlayerLander) IsThrusting() data.ThrusterState {
	return data.ThrusterState{true, true, true}
}

func (pl *PlayerLander) Update(timeDelta float64, thrusterState data.ThrusterState) {

}

func (pl *PlayerLander) SetPosition(position data.Vector2D) {
	pl.PosX = float32(position.X)
	pl.PosY = float32(position.Y)
}
