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
	Connection  net.Conn
	receiveFunc func(data string)
}

func (c *Client) sendStuff(conn net.Conn) {
	message := bufio.NewReader(os.Stdin)
	fmt.Println("Tell me something")
	for {
		if line, err := message.ReadString('\n'); err != nil {
			fmt.Println("###Some shitty shit happend")
			conn.Close()
		} else {
			_, err = conn.Write([]byte(line))
			checkError(err)
		}
	}
}

func (c *Client) receiveStuff(conn net.Conn) {
	buffer := bufio.NewReader(conn)
	for {
		str, err := buffer.ReadString('\n')
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

	conn, err := net.Dial("tcp", address)
	checkError(err)

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
	c.sendStuff(conn)
}

func NewClient(receiveFunc func(string)) *Client {
	client := new(Client)
	client.receiveFunc = receiveFunc
	return client
}

// gimme some errors if con does not work
func checkError(err error) {
	if err != nil {
		fmt.Println("#### Some realy fatal shit is going on: ", err.Error())
	}
}
