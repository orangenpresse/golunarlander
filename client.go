package main

import (
	"fmt"
	"github.com/orangenpresse/golunarlander/network"
)

func main() {
	client := network.NewClient(receive)
	client.Connect("127.0.0.1:4711")
}

func receive(data string) {
	fmt.Println(data)
}
