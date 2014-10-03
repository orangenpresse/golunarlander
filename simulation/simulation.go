package simulation

import (
	"fmt"
	_ "math"
	_ "time"
)

type Vector2D struct {
	X float64
	Y float64
}

const (
	G        = 1.635
	interval = 0.001
	endTime  = 200000
)

type Lander struct {
	position Vector2D
	velocity Vector2D
	thrust   float64
}

func (lander *Lander) GetPosition() Vector2D {
	return lander.position
}

func physic(lander *Lander, channel chan int64) {
	for ; currentTime <= endTime; currentTime += int64(interval * 1000) {

		acceleration := -G + lander.thrust
		lander.velocity.Y += acceleration * interval

		lander.position.Y += lander.velocity.Y * interval

		if lander.position.Y <= 0.0 {
			fmt.Printf("Aufprall bei t=%d mit v=%f\n", currentTime, lander.velocity.Y)
			break
		}

		//fmt.Printf("t=%d, velY=%f, posY=%f\n", currentTime, lander.velocity.Y, lander.position.Y)
	}

	channel <- currentTime
}

func control(lander *Lander, channel chan int64) {
	for lander.position.Y >= 0.5 {
		if lander.velocity.Y < -20 {
			lander.thrust = 3
		} else {
			lander.thrust = 0
		}
	}

	channel <- currentTime
}

var currentTime int64 = 0

// func main() {
// 	lander := new(Lander)

// 	lander.position.X = 0
// 	lander.position.Y = 10000

// 	lander.velocity.X = 0
// 	lander.velocity.Y = 0

// 	channel := make(chan int64)

// 	go physic(lander, channel)
// 	go control(lander, channel)

// 	result := <-channel
// 	result = <-channel
// 	fmt.Printf("Ende bei t=%d\n", result)
// }
