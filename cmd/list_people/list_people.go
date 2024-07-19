package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/madatsci/go-protobuf-example/protobuf/examples/go/tutorialpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	fname := "out_file"

	// Read the existing address book.
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Printf("%+v\n", book)
}
