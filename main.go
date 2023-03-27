package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"j2p/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

type Person struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	data     = flag.String("d", `{ "id": 1234, "name": "John Doe", "email": "jdoe@example.com" }`, "json string")
	fileName = flag.String("f", "book.bin", "file name")
)

func main() {
	flag.Parse()

	var person Person
	if err := json.Unmarshal([]byte(*data), &person); err != nil {
		fmt.Printf("Failed to unmarshal json string: %v", err)
	}

	p := pb.Person{
		Id:    person.Id,
		Name:  person.Name,
		Email: person.Email,
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	book := &pb.AddressBook{
		People: []*pb.Person{&p},
	}

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(*fileName, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
