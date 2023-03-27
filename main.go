package main

import (
	"io/ioutil"
	"j2p/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
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
	fname := "book.bin"
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
