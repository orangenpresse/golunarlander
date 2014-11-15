package multiplayer

import (
	"fmt"
	"github.com/orangenpresse/golunarlander/lander"
	"github.com/orangenpresse/golunarlander/network"
)

type Multiplayer struct {
	client  *network.Client
	landers []lander.LanderInterface
	lander  lander.LanderInterface
}

func NewMultiplayer(lander lander.LanderInterface) *Multiplayer {
	multi := new(Multiplayer)
	multi.lander = lander

	client := network.NewClient(multi.receive)
	go client.Connect("127.0.0.1:4711")

	multi.client = client
	return multi
}

func (m *Multiplayer) receive(data string) {
	fmt.Println(data)
}

func (m *Multiplayer) SendUpdate() {
	m.client.SendData("test")
}

func (m *Multiplayer) GetLanders() {

}
