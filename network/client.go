package main

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

func sendStuff(conn net.Conn) {
	message := bufio.NewReader(os.Stdin)
	fmt.Println("Tell me something")
	for {
		line, err := message.ReadString('\n')
		if err != nil {
			fmt.Println("###Some shitty shit happend")
			conn.Close()
			break
		}
		_, err = conn.Write([]byte(line))
		checkError(err)
	}
}

func receiveStuff(conn net.Conn) {
	buffer := bufio.NewReader(conn)
	for {
		str, err := buffer.ReadString('\n')
		checkError(err)

		if len(str) > 0 {
			fmt.Println("SERVER RETURNS:", str)
		}
		if err != nil {
			break
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	conn, err := net.Dial("tcp", "127.0.0.1:4711")
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

	go receiveStuff(conn)
	sendStuff(conn)
}

// gimme some errors if con does not work
func checkError(err error) {
	if err != nil {
		fmt.Println("#### Some realy fatal shit is going on: ", err.Error())
	}
}
