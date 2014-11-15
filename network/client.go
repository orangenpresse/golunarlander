package network

import (
	"bufio"
	"fmt"
	//"io"
	"container/list"
	"net"
	"os"
	"strings"
)

type Session struct {
	Incoming   chan string
	Outgoing   chan string
	Conn       net.Conn
	ClientList *list.List
}

type Client struct {
	connected   bool
	Connection  net.Conn
	outgoing    chan string
	receiveFunc func(data string)
}

func (c *Client) sendStuff(conn net.Conn) {
	for {
		data := <-c.outgoing
		_, err := conn.Write([]byte(data))
		checkError(err)
	}
}

func (c *Client) SendData(data string) {
	if c.connected == false {
		return
	}
	select {
	case c.outgoing <- data:
	default: //Discard values if channel is full
	}

}

func (c *Client) receiveStuff(conn net.Conn) {
	buffer := bufio.NewReader(conn)
	for {
		str, err := buffer.ReadString(';')
		checkError(err)

		if len(str) > 0 {
			c.receiveFunc(str)
		}
		if err != nil {
			break
		}
	}
}

func (c *Client) Connect(address string) {
	reader := bufio.NewReader(os.Stdin)

	if conn, err := net.Dial("tcp", address); err != nil {
		c.connected = false
		checkError(err)
		return
	} else {
		bufferId := bufio.NewReader(conn)
		myId, err := bufferId.ReadString('\n')
		checkError(err)

		fmt.Println(myId)

		fmt.Println("Who's there?")
		username, err := reader.ReadString('\n')
		username = strings.TrimRight(string(username), "\n")

		checkError(err)

		_, err = conn.Write([]byte(username))
		checkError(err)

		go c.receiveStuff(conn)
		c.connected = true
		c.sendStuff(conn)
	}
}

func NewClient(receiveFunc func(string)) *Client {
	client := new(Client)
	client.receiveFunc = receiveFunc
	client.outgoing = make(chan string)
	return client
}

// gimme some errors if con does not work
func checkError(err error) {
	if err != nil {
		fmt.Println("#### Some realy fatal shit is going on: ", err.Error())
	}
}
