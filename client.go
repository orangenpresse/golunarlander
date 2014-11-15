package main

import (
	"github.com/orangenpresse/golunarlander/network"
)

func main() {
	client := network.NewClient()
	client.Connect("127.0.0.1:4711")
}
