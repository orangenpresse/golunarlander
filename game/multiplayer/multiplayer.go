package multiplayer

import (
	"github.com/orangenpresse/golunarlander/lander"
	"github.com/orangenpresse/golunarlander/network"
)

type Multiplayer struct {
	client  network.Client
	landers []lander.LanderInterface
}

func NewMultiplayer() *Multiplayer {
	client := network.NewClient()
	client.Connect("127.0.0.1:4711")

	multi := new(Multiplayer)
	multi.client = client

	return multi
}

func (m *Multiplayer) GetLanders() {

}
