package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
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
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	p1 := person.Person{Name: "Da", Age: 21}

	data, err := p1.Encode()
	if err != nil {
		fmt.Println("Error while encoding Person struct: ", err.Error())
		os.Exit(1)
	}
	dataType := TypePerson

	_, err = conn.Write(SerializeInt(dataType))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	dataLength := len(data)

	_, err = conn.Write(SerializeInt(dataLength))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	_, err = conn.Write(data)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
}
