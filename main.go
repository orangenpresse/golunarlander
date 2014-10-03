package main

import (
	"fmt"
	_ "math"
	_ "time"
)

const (
	G        = 1.635
	interval = 0.001
	endTime  = 200000
)

type Vector2D struct {
	x float64
	y float64
}

type Lander struct {
	position Vector2D
	velocity Vector2D
	thrust   float64
}

type BetterLAnder struct {
	Lander
}

func (lander *Lander) GetPosition() Vector2D {
	return lander.position
}

func physic(lander *Lander, channel chan int64) {
	for ; currentTime <= endTime; currentTime += int64(interval * 1000) {

		acceleration := -G + lander.thrust
		lander.velocity.y += acceleration * interval

		lander.position.y += lander.velocity.y * interval

		if lander.position.y <= 0.0 {
			fmt.Printf("Aufprall bei t=%d mit v=%f\n", currentTime, lander.velocity.y)
			break
		}

		//fmt.Printf("t=%d, velY=%f, posY=%f\n", currentTime, lander.velocity.y, lander.position.y)
	}

	channel <- currentTime
}

func control(lander *Lander, channel chan int64) {
	for lander.position.y >= 0.5 {
		if lander.velocity.y < -20 {
			lander.thrust = 3
		} else {
			lander.thrust = 0
		}
	}

	channel <- currentTime
}

var currentTime int64 = 0

func main() {
	lander := new(Lander)

	lander.position.x = 0
	lander.position.y = 10000

	lander.velocity.x = 0
	lander.velocity.y = 0

	channel := make(chan int64)

	go physic(lander, channel)
	go control(lander, channel)

	result := <-channel
	result = <-channel
	fmt.Printf("Ende bei t=%d\n", result)
}
