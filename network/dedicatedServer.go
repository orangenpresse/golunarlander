package network

import (
	"container/list"
	"fmt"
	"net"
	"strconv"
)

type ClientConnection struct {
	Parent     *Server
	Id         int
	Name       string
	LastState  string
	Connection net.Conn
	FromClient chan string
	ToClient   chan string
}

func (this *ClientConnection) ReaderThread() {
	for {
		buffer := make([]byte, 2048)
		bytesRead, error := this.Connection.Read(buffer)

		if error != nil {
			fmt.Println("Readen gefailed")
			this.Connection.Close()
			return
		}

		bufferString := string(buffer[0:bytesRead])
		this.LastState = bufferString
		this.Parent.SendUpdateToAllClients()
		fmt.Println("Gereaded: " + bufferString)
	}
}

func (this *ClientConnection) send(data string) {
	fmt.Printf("Sending to Client(%d, %s): %s\n", this.Id, this.Name, data)
	this.Connection.Write([]byte(data))
}

type Server struct {
	Clients  *list.List
	NextID   int
	Listener net.Listener
	Dirty    bool
}

func CreateServer(port int16) Server {
	var err error
	server := Server{}
	server.Clients = list.New()
	server.NextID = 1
	server.Listener, err = net.Listen("tcp", ":"+strconv.Itoa(int(port)))

	if err != nil {
		fmt.Println(err.Error())
	}

	return server
}

func (this *Server) SendUpdateToAllClients() {
	fmt.Println("I'm so Dirty, sending update")
	updateData := this.makeUpdate()

	for e := this.Clients.Front(); e != nil; e = e.Next() {
		fmt.Println("Sending to:", e.Value)
		e.Value.(*ClientConnection).send(updateData)
	}
}

func (this *Server) makeUpdate() string {
	data := ""
	for e := this.Clients.Front(); e != nil; e = e.Next() {
		data += strconv.Itoa(e.Value.(*ClientConnection).Id) + ","
		data += e.Value.(*ClientConnection).Name + ","
		data += e.Value.(*ClientConnection).LastState + ";"
	}

	defer fmt.Println(data)
	return data + "\n"
}

func (this *Server) OnNewConnection(connection net.Conn) {

	connection.Write([]byte(strconv.Itoa(this.NextID) + "\n"))
	buffer := make([]byte, 1024)
	bytesRead, err := connection.Read(buffer)

	if err != nil {
		fmt.Println("Error reading Name from Client: " + err.Error())
		return
	}
	client := ClientConnection{}
	client.Parent = this
	client.Id = this.NextID
	client.Name = string(buffer[0:bytesRead])
	client.LastState = ""
	client.Connection = connection
	client.FromClient = make(chan string)
	client.ToClient = make(chan string)
	this.Clients.PushBack(&client)
	go client.ReaderThread()

	this.NextID++
	fmt.Println("Connected Client: " + client.Name + " with ID:" + strconv.Itoa(client.Id))
}

func NewServer(port int16) {
	server := CreateServer(port)
	defer server.Listener.Close()

	for {
		fmt.Println("Waiting...")
		connection, err := server.Listener.Accept()

		if err != nil {
			fmt.Println("Error while a client tried to connect: " + err.Error())
			continue
		}
		server.OnNewConnection(connection)
	}

}
