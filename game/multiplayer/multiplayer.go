package multiplayer

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/orangenpresse/golunarlander/lander"
	"github.com/orangenpresse/golunarlander/network"
	"strconv"
)

type Multiplayer struct {
	client  *network.Client
	landers chan []lander.LanderInterface
	Lander  lander.LanderInterface
}

func NewMultiplayer(localLander lander.LanderInterface) *Multiplayer {
	multi := new(Multiplayer)
	multi.Lander = localLander
	multi.landers = make(chan []lander.LanderInterface, 1)

	client := network.NewClient(multi.receive)
	go client.Connect("127.0.0.1:4711")

	multi.client = client
	return multi
}

func (m *Multiplayer) receive(data string) {
	landers := make([]lander.LanderInterface, 0)

	buffer := bytes.NewBufferString(data)
	csvReader := csv.NewReader(buffer)

	for {
		if line, err := csvReader.Read(); err != nil || len(line) < 4 {
			break
		} else {
			landers = append(landers, generateLander(line))
		}
	}

	select {
	case m.landers <- landers:
	default: //Discard landers if not read yet
	}

}

func generateLander(csv []string) lander.LanderInterface {
	// [Id Name X Y]
	return &PlayerLander{csv[0], csv[1], parseFloat(csv[2]), parseFloat(csv[3])}
}

func (m *Multiplayer) SendUpdate() {
	pos := m.Lander.GetPosition()
	m.client.SendData(fmt.Sprintf("%f,%f", pos.X, pos.Y))
}

func (m *Multiplayer) GetLanders() []lander.LanderInterface {
	select {
	case landers := <-m.landers:
		return landers
	default:
		return make([]lander.LanderInterface, 0)
	}
}

func parseFloat(input string) float32 {
	if res, err := strconv.ParseFloat(input, 32); err == nil {
		return float32(res)
	}
	return 0
}
