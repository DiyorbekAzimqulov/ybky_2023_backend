package encode_decode_benchmark

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func MarshalJSON(people Person) []byte {
	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
	}
	return jsonData
}

func UnmarshalJSON(data []byte, people *Person) {
	err := json.Unmarshal(data, people)
	if err != nil {
		fmt.Println("Error unmarshalling! ", err)
	}
}

func MarshalBSON(people Person) []byte {
	bsonData, err := bson.Marshal(people)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
	}
	return bsonData
}

func UnmarshalBSON(data []byte, people *Person) {
	err := bson.Unmarshal(data, people)
	if err != nil {
		fmt.Println("Error unmarshalling! ", err)
	}
}

func MarshalProtoBuf(person *Personn) []byte {
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("There is marshalling error: ", err)
	}
	return data
}

func UnmarshalProtoBuf(data []byte, person *Personn) {
	err := proto.Unmarshal(data, person)
	if err != nil {
		fmt.Println("Error unmarshalling: ", err)
	}
}

func encodeGOB(person Person, buffer *bytes.Buffer) {
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding GOB: ", err)
	}
}

func decodeGOB(person *Person, gobData []byte) {
	buffer := bytes.NewBuffer(gobData)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(person)
	if err != nil {
		fmt.Println("Error decoding GOB: ", err)
	}
}

func BenchmarkEncodeJSON(b *testing.B) {
	people := Person{"Da", 21}
	for i := 0; i < b.N; i++ {
		MarshalJSON(people)
	}
}

func BenchmarkDecodeJSON(b *testing.B) {
	people := Person{"Da", 21}
	jsonData := MarshalJSON(people)
	var person1 Person
	for i := 0; i < b.N; i++ {
		UnmarshalJSON(jsonData, &person1)
	}
}

func BenchmarkEncodeBSON(b *testing.B) {
	people := Person{"Da", 21}
	for i := 0; i < b.N; i++ {
		MarshalBSON(people)
	}
}

func BenchmarkDecodeBSON(b *testing.B) {
	people := Person{"Da", 21}
	bsonData := MarshalBSON(people)
	var person1 Person
	for i := 0; i < b.N; i++ {
		UnmarshalBSON(bsonData, &person1)
	}
}

func BenchmarkEncodeProtoBuf(b *testing.B) {
	person := &Personn{
		Name: "Da",
		Age:  21,
	}
	for i := 0; i < b.N; i++ {
		MarshalProtoBuf(person)
	}
}

func BenchmarkDecodeProtoBuf(b *testing.B) {
	person := &Personn{
		Name: "Da",
		Age:  21,
	}
	protoBufData := MarshalProtoBuf(person)
	var person1 Personn
	for i := 0; i < b.N; i++ {
		UnmarshalProtoBuf(protoBufData, &person1)
	}
}

func BenchmarkEncodeGOB(b *testing.B) {
	person := Person{"Da", 21}
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		encodeGOB(person, &buffer)
	}
}

func BenchmarkDecodeGOB(b *testing.B) {
	var buffer bytes.Buffer
	person := Person{"Da", 21}
	encodeGOB(person, &buffer)
	for i := 0; i < b.N; i++ {
		decodeGOB(&person, buffer.Bytes())
	}
}
