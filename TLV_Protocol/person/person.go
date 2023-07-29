package person

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
)

type Person struct {
	Name string
	Age  int
}

func (person *Person) Encode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(person)
	if err != nil {
		return nil, err
	}
	data := buf.Bytes()
	return data, nil
}

func (person *Person) Decode(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(person)
	if err != nil {
		return err
	}
	return nil
}

func (person *Person) Read(conn io.Reader) error {
	buffer := make([]byte, 2)
	if _, err := conn.Read(buffer); err != nil {
		return err
	}
	length := binary.LittleEndian.Uint16(buffer)
	buffer = make([]byte, length)
	if _, err := conn.Read(buffer); err != nil {
		return err
	}
	if err := person.Decode(buffer); err != nil {
		return err
	}
	return nil
}

func (person *Person) Write(conn io.Writer) error {
	data, err := person.Encode()
	if err != nil {
		return err
	}
	length := make([]byte, 2)
	binary.LittleEndian.PutUint16(length, uint16(len(data)))
	if _, err := conn.Write(length); err != nil {
		return err
	}
	if _, err := conn.Write(data); err != nil {
		return err
	}
	return nil
}
