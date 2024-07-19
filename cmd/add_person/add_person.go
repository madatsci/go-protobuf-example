package main

import (
	"log"
	"os"

	pb "github.com/madatsci/go-protobuf-example/protobuf/examples/go/tutorialpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	book := &pb.AddressBook{}

	person := &pb.Person{
		Name:  "Andrei",
		Id:    1234,
		Email: "sedemaqu@gmail.com",
	}

	book.People = []*pb.Person{person}

	fname := "out_file"

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
