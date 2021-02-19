package main

import (
	"fmt"
	"go_protobuf/filerw"
	"go_protobuf/jsonrw"
	simplepb "go_protobuf/source/simple"
	"log"
)

func main() {
	sm := filerw.CreateSimple()
	readWriteFile(sm)
	readWriteJSON(sm)

}

func readWriteJSON(sm *simplepb.SimpleMessage) {
	fmt.Println("Convert to and from a JSON string")

	jsonString := jsonrw.ToJSON(sm)
	fmt.Println("JSON String Converted From Proto Message :  ", string(jsonString))

	sm2 := &simplepb.SimpleMessage{}
	jsonrw.FromJSON(jsonString, sm2)
	fmt.Println("Proto Message Converted From JSON String :  ", sm2)

}

func readWriteFile(sm *simplepb.SimpleMessage) {
	fmt.Println("Read and Write to and from a Binary File")

	if err := filerw.WriteToFile("simple.bin", sm); err != nil {
		log.Fatalln("Writing to file failed", err)
	}

	sm2 := &simplepb.SimpleMessage{}
	if err2 := filerw.ReadFromFile("simple.bin", sm2); err2 != nil {
		log.Fatalln("Reading from file failed", err2)
	}

	fmt.Println(sm)
	fmt.Println(sm2)
}
