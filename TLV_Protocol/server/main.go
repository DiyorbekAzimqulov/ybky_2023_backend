package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"ybky/TLV_Protocol/person"
)

const (
	HOST = "localhost"
	PORT = "8000"
	TYPE = "tcp"
)

const (
	TypePerson = 1
)

func SerializeInt(value int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(value))
	return buf
}

// DeserializeInt deserializes bytes to an integer value.
func DeserializeInt(data []byte) int {
	return int(binary.BigEndian.Uint32(data))
}

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error while trying to listen on: ", HOST, PORT)
	}
	// close the Listener after server terminates
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection!", err)
		}

		go HandleRequest(conn)
	}
}

func HandleRequest(conn net.Conn) {
	// type
	dataType := make([]byte, 4)

	_, err := conn.Read(dataType)
	if err != nil {
		fmt.Println("Error while reading to the buffer from tcp connection: ", err)
	}
	dataLenght := make([]byte, 4)

	_, err = conn.Read(dataLenght)
	if err != nil {
		fmt.Println("Error while reading to the buffer from tcp connection: ", err)
	}

	data := make([]byte, DeserializeInt(dataLenght))

	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("Error while reading to the buffer from tcp connection: ", err)
	}
	switch DeserializeInt(dataType) {
	case TypePerson:
		var p person.Person
		p.Decode(data)
		fmt.Println(p)
	}
	// close conn
	conn.Close()
}
